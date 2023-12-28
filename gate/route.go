package gate

import (
	"github.com/0990/goserver/network"
	"github.com/golang/protobuf/proto"
)

func Route2ServerID(msg proto.Message, serverID int32) {
	Gate.RegisterRawSessionMsgHandler(msg, func(session network.Session, message proto.Message) {
		s, exist := SMgr.sesID2Session[session.ID()]
		if !exist {
			return
		}
		if !s.logined {
			return
		}
		Gate.GetServerById(serverID).RouteSession2Server(session.ID(), message)
	})
}
