package client

import (
	"fmt"
	"github.com/0990/avatar-fight-server/msg/cmsg"
	"github.com/0990/goserver/network"
	"time"
)

type Client struct {
	token      string
	userId     uint64
	nickname   string
	engine     *Engine
	gameEngine *Engine
}

func NewClient(targetAddr string, token string) *Client {
	e := NewEngine(targetAddr)
	err := e.Run()
	if err != nil {
		panic(err)
	}

	c := &Client{}
	c.engine = e
	c.token = token
	c.nickname = token

	e.RegisterSessionMsgHandler(c.onRespLogin)
	e.RegisterSessionMsgHandler(c.onSNotifyGameReady)
	return c
}

func (c *Client) StartLogin() {
	c.engine.SendMsg(&cmsg.ReqLogin{
		Token: c.token,
	})
}

func (c *Client) onRespLogin(session network.Session, msg *cmsg.RespLogin) {
	fmt.Println("respLogin", msg)
	c.userId = msg.UserID
	c.engine.SendMsg(&cmsg.ReqMatch{})
}

func (c *Client) onSNotifyGameReady(session network.Session, msg *cmsg.SNotifyGameReady) {
	fmt.Println("onSNotifyGameReady", msg)

	e := NewEngine(msg.GameAddr)
	err := e.Run()
	if err != nil {
		panic(err)
	}
	c.gameEngine = e

	c.gameEngine.SendMsg(&cmsg.ReqGameLogin{
		UserId:   c.userId,
		Nickname: c.nickname,
	})

	c.gameEngine.RegisterSessionMsgHandler(c.onRespGameLogin)
	c.gameEngine.RegisterSessionMsgHandler(c.onSNotifyGameStart)
	c.gameEngine.RegisterSessionMsgHandler(c.onSyncFrame)
	c.gameEngine.RegisterSessionMsgHandler(c.onNotifyGameOver)
}

func (c *Client) onRespGameLogin(session network.Session, msg *cmsg.RespGameLogin) {
	fmt.Println("RespGameLogin", msg)
	c.gameEngine.SendMsg(&cmsg.ReqClientReady{})
}

func (c *Client) onSNotifyGameStart(session network.Session, msg *cmsg.SNotifyGameStart) {
	fmt.Println("SNotifyGameStart", msg)
	time.AfterFunc(time.Second*5, func() {
		c.gameEngine.SendMsg(&cmsg.ReqGameOver{})
	})
}

func (c *Client) onSyncFrame(session network.Session, msg *cmsg.SSyncFrame) {
	fmt.Println("SSyncFrame", msg)
}

func (c *Client) onNotifyGameOver(session network.Session, msg *cmsg.SNotifyGameOver) {
	fmt.Println("SNoticeGameOver", msg)
}
