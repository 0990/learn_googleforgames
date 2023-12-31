package gate

import (
	"github.com/0990/avatar-fight-server/conf"
	"github.com/0990/avatar-fight-server/msg/cmsg"
	"github.com/0990/goserver"
	"github.com/0990/goserver/server"
)

var Gate *server.Gate

var SMgr *SessionMgr

func Init(serverID int32, addr string, config goserver.Config) error {
	s, err := server.NewGate(serverID, addr, config)
	if err != nil {
		return err
	}
	Gate = s

	Route2ServerID(&cmsg.ReqMatch{}, conf.CenterServerID)
	registerHandler()
	SMgr = newSessionMgr()
	return nil
}

func Run() {
	Gate.Run()
}
