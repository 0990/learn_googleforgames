package game

import (
	"github.com/0990/avatar-fight-server/msg/cmsg"
	"github.com/0990/goserver/network"
)

func (a *App) registerHandler() {
	a.engine.RegisterNetWorkEvent(a.onConnect, a.onDisconnect)
	a.engine.RegisterSessionMsgHandler(a.onReqLogin)
	a.engine.RegisterSessionMsgHandler(a.onReqClientReady)
	a.engine.RegisterSessionMsgHandler(a.onReqGameInput)
	a.engine.RegisterSessionMsgHandler(a.onReqGameOver)
}

func (a *App) onConnect(conn network.Session) {
	a.sessionMgr.sesID2Session[conn.ID()] = &Session{
		session: conn,
		sesID:   conn.ID(),
	}
}

func (a *App) onDisconnect(conn network.Session) {
	if session, ok := a.sessionMgr.sesID2Session[conn.ID()]; ok && session.logined {
		a.game.eventProcessor.AddEvent(Event{
			Type:   PlayerLogout,
			UserId: session.userID,
		})
	}

	delete(a.sessionMgr.sesID2Session, conn.ID())
}

func (a *App) onReqLogin(session network.Session, msg *cmsg.ReqGameLogin) {
	a.sessionMgr.SetSessionLogined(session.ID(), msg.UserId)

	a.game.eventProcessor.AddEvent(Event{
		Type:   PlayerLogin,
		UserId: msg.UserId,
		Sender: session,
		Data:   &EventPlayerLoginData{session: session, userId: msg.UserId, nickname: msg.Nickname},
	})
}

func (a *App) onReqClientReady(session network.Session, msg *cmsg.ReqClientReady) {
	resp := &cmsg.RespClientReady{}
	s, ok := a.sessionMgr.sesID2Session[session.ID()]

	if !ok {
		resp.Err = 1
		session.SendMsg(resp)
		return
	}

	if !s.logined {
		resp.Err = 2
		session.SendMsg(resp)
		return
	}

	a.game.eventProcessor.AddEvent(Event{
		Type:   ClientMsg_ReqClientReady,
		UserId: s.userID,
		Data:   msg,
		Sender: session,
	})
}

func (a *App) onReqGameInput(session network.Session, msg *cmsg.ReqGameInput) {
	s, ok := a.sessionMgr.sesID2Session[session.ID()]

	if !ok {
		return
	}

	if !s.logined {
		return
	}

	a.game.eventProcessor.AddEvent(Event{
		Type:   ClientMsg_ReqInput,
		UserId: s.userID,
		Data:   msg,
		Sender: session,
	})
}

func (a *App) onReqGameOver(session network.Session, msg *cmsg.ReqGameOver) {
	s, ok := a.sessionMgr.sesID2Session[session.ID()]

	if !ok {
		return
	}

	if !s.logined {
		return
	}

	a.game.eventProcessor.AddEvent(Event{
		Type:   ClientMsg_ReqGameOver,
		UserId: s.userID,
		Data:   msg,
		Sender: session,
	})
}
