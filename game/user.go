package game

import (
	"github.com/0990/goserver/network"
	"github.com/golang/protobuf/proto"
)

type AccountType int8

const (
	_ AccountType = iota
	VISITOR
	WX
	ROBOT
)

type User struct {
	session     network.Session
	userID      uint64
	nickname    string
	accountType AccountType
	headImgUrl  string
	offline     bool
}

func (p *User) Send2Client(msg proto.Message) {
	if p.accountType == ROBOT {
		return
	}
	if p.offline {
		return
	}
	p.session.SendMsg(msg)
}
