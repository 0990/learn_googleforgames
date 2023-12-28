package center

import (
	"fmt"
	"github.com/0990/avatar-fight-server/msg/cmsg"
	"github.com/0990/avatar-fight-server/pkg/gameallocator"
	"github.com/0990/avatar-fight-server/pkg/matchmaker"
	"github.com/0990/goserver"
	"github.com/0990/goserver/server"
	"time"
)

var Server *server.Server

var UMgr *UserMgr
var GMgr *GameMgr
var GMatchMaker matchmaker.MatchMaker

func Init(serverID int32, config goserver.Config) error {
	s, err := server.NewServer(serverID, config)
	if err != nil {
		return err
	}
	Server = s
	registerHandler()
	GMgr = NewGameMgr()
	UMgr = NewUserMgr()
	GMatchMaker, err = matchmaker.New(matchmaker.Config{
		Type: matchmaker.OpenMatch,
		OpenMatch: matchmaker.OpenMatchArgs{
			FrontendEndPoint:  "172.21.248.145:30504",
			BackendEndPoint:   "172.21.248.145:30505",
			MatchFunctionHost: "10.242.20.91",
			MatchFunctionPort: 50502,
		},
	})
	if err != nil {
		return err
	}

	gameAllocator, err := gameallocator.New(gameallocator.Config{
		Type: gameallocator.Agones,
	})

	if err != nil {
		return err
	}

	GMatchMaker.WatchMatch(func(match *matchmaker.Match) {
		go func() {
			now := time.Now()
			addr, err := gameAllocator.Allocate(match)
			fmt.Printf("elapse:%v,allocate addr:%v,%v \n", time.Since(now), addr, err)

			if err != nil {
				fmt.Println(err)
				return
			}

			Server.Post(func() {
				for userId, _ := range match.Users {
					u, ok := UMgr.FindUserByUserId(userId)
					if !ok {
						continue
					}
					msg := &cmsg.SNotifyGameReady{GameAddr: fmt.Sprintf("ws://%s", addr.String())}
					u.SendClientMsg(msg)
				}
			})
		}()
	})
	return nil
}

func Run() {
	Server.Run()
	GMatchMaker.Run()
}
