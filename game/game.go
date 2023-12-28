package game

import (
	"fmt"
	"github.com/0990/avatar-fight-server/msg/cmsg"
	"github.com/0990/avatar-fight-server/util"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

type OverReason int8

const (
	Invalid OverReason = iota
	Killed
	Normal
)

type Game struct {
	gameId                         string
	createTime, startTime, endTime time.Time

	onGameEnd func(*Game)

	players []*Player

	stop chan struct{}

	closed atomic.Bool

	inputQueue *InputQueue

	eventProcessor *EventProcessor
	seatIndex      int32

	allocateUsers []string
}

func newGame(gameID string, onGameEnd func(*Game)) *Game {
	g := new(Game)
	g.gameId = gameID
	g.onGameEnd = onGameEnd
	g.inputQueue = NewInputQueue()
	g.eventProcessor = NewEventProcessor()
	g.stop = make(chan struct{}, 0)
	return g
}

func (g *Game) Run() {
	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(*PanicInfo); ok {
				fmt.Println(v.ToString())
			} else {
				fmt.Println(err)
				//gameutil.PrintStack()
			}
		}

		g.SetOver()
	}()

	ok := g.waitPlayerConnect(10 * time.Minute)
	if !ok {
		return
	}

	var users []uint64
	for _, v := range g.players {
		users = append(users, v.user.userID)
	}
	log.Printf("all player connected,allocate:%v,connect:%v \n", g.allocateUsers, users)
	ok = g.waitPlayerReady(10 * time.Minute)
	if !ok {
		return
	}
	log.Println("all player ready,game start")
	g.Start()
}

func (g *Game) waitPlayerConnect(duration time.Duration) bool {
	c := time.After(duration)
	for {
		select {
		case e, ok := <-g.eventProcessor.Event():
			if !ok {
				return false
			}

			g.handleCommonEvent(e)

			if e.Type != PlayerLogin {
				continue
			}

			data := e.Data.(*EventPlayerLoginData)

			_, ok = g.GetPlayerById(e.UserId)
			if ok {
				continue
			}

			g.JoinPlayer(data)

			e.Sender.SendMsg(&cmsg.RespGameLogin{})
			if len(g.players) >= len(g.allocateUsers) {
				return true
			}

			continue
		case <-c:
			return false
		}
	}
}

func (g *Game) waitPlayerReady(duration time.Duration) bool {
	c := time.After(duration)
	for {
		select {
		case e, ok := <-g.eventProcessor.Event():
			if !ok {
				return false
			}
			g.handleCommonEvent(e)

			if e.Type != ClientMsg_ReqClientReady {
				continue
			}

			p, ok := g.GetPlayerById(e.UserId)
			if !ok {
				continue
			}

			p.clientReady = true

			resp := &cmsg.RespClientReady{}
			e.Sender.SendMsg(resp)

			someoneNotReady := util.IsExist(g.players, func(p *Player) bool {
				return p.clientReady == false && !p.user.offline
			})

			if !someoneNotReady {
				return true
			}
			continue
		case <-c:
			return false
		}
	}
}

func (g *Game) handleCommonEvent(e Event) {
	switch e.Type {
	case PlayerLogout:
		p, ok := g.GetPlayerById(e.UserId)
		if !ok {
			return
		}
		p.user.offline = true
	case AgonesAllocate:
		data := e.Data.(*EventAgonesAllocateData)
		g.allocateUsers = data.users
	default:

	}
}

func (g *Game) Start() {
	var seats []*cmsg.Seat
	for _, v := range g.players {
		seats = append(seats, &cmsg.Seat{
			SeatId:   v.seat,
			UserId:   v.user.userID,
			Nickname: v.user.nickname,
		})
	}

	seed := int64(rand.Int31())

	msg := &cmsg.SNotifyGameStart{
		RandSeed: seed,
		Seats:    seats,
	}
	g.SendMsg2All(msg)

	logrus.WithFields(logrus.Fields{
		"gameId": g.gameId,
		"msg":    msg,
	}).Debug("SNotifyGameStart")

	g.writeLoop()
}

func (g *Game) writeLoop() bool {
	ticker := time.NewTicker(FrameIntervalMS * time.Millisecond)
	timeoutTimer := time.NewTimer(GameTimeout)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			g.tickFrame()
		case <-timeoutTimer.C:
			PanicStopGame("game run too long time")
		case <-g.stop:
			return true
		case e, ok := <-g.eventProcessor.events:
			if !ok {
				return false
			}
			g.handEvent(e)
		}
	}
	return false
}

func (g *Game) tickFrame() {
	id := g.inputQueue.FrameId()
	inputs := g.inputQueue.Inputs()

	msg := &cmsg.SSyncFrame{
		Frame:  id,
		Inputs: inputs,
	}
	g.SendMsg2All(msg)

	g.inputQueue.Tick()
}

func (g *Game) SetOver() {
	log.Printf("game over:%s", g.gameId)
	g.endTime = time.Now()
	g.SendMsg2All(&cmsg.SNotifyGameOver{})
}

func (p *Game) SendMsg2All(msg proto.Message) {
	for _, v := range p.players {
		v.user.Send2Client(msg)
	}
}
