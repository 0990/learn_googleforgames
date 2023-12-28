package game

import (
	"github.com/0990/avatar-fight-server/msg/cmsg"
)

func (g *Game) handEvent(e Event) {
	p, ok := g.GetPlayerById(e.UserId)
	if !ok {
		return
	}

	switch e.Type {
	case PlayerDisconnect:
		//g.onPlayerDisconnect(p)
	case PlayerReconnect:
		//g.onPlayerReconnect(p, e.Data.(appframe.Session))
	case ClientMsg_ReqInput:
		g.onPlayerGameInput(p, e.Data.(*cmsg.ReqGameInput))
	case ClientMsg_ReqGameOver:
		g.onPlayerGameOver(p, e.Data.(*cmsg.ReqGameOver))
	//case ClientMsg_ReportFrameKeyData:
	//	d.base.onPlayerReportFrameKeyData(p, e.Data.(*cmsg.CReqReportFrameKeyData))
	//case ClientMsg_ReportSeatData:
	//	d.base.onPlayerReportAllSeatData(p, e.Data.(*cmsg.CReportSeatData))
	//case ClientMsg_ReqBroadcast:
	//	d.base.onPlayerBroadcast(p, e.Data.(*cmsg.CReqBroadcast))
	//case ClientMsg_Ping:
	//	g.onPlayerGamePing(p, e.Data.(*cmsg.CReqGamePing))
	default:

	}
}

func (g *Game) onPlayerGameInput(p *Player, req *cmsg.ReqGameInput) {
	g.inputQueue.Push(&cmsg.SeatInput{
		Seat:  p.seat,
		Input: req.Input,
	})
}

func (g *Game) onPlayerGameOver(p *Player, req *cmsg.ReqGameOver) {
	if g.closed.CompareAndSwap(false, true) {
		close(g.stop)
	}
}
