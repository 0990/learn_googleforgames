package game

import (
	"github.com/0990/goserver/network"
	"github.com/0990/goserver/service"
	"net"
	"time"
)

type Engine struct {
	worker      service.Worker
	networkMgr  *network.Mgr
	onCloseFuns []func()
}

func NewEngine(addr string) *Engine {
	p := &Engine{}
	p.worker = service.NewWorker()
	p.networkMgr = network.NewMgr(addr, p.worker)
	return p
}

func (p *Engine) Run() {
	p.worker.Run()
	p.networkMgr.Run()
}

func (p *Engine) ListenAddr() *net.TCPAddr {
	return p.networkMgr.ListenAddr()
}

func (p *Engine) Close() {
	for _, f := range p.onCloseFuns {
		f()
	}
	p.networkMgr.Close()
	p.worker.Close()
}

func (p *Engine) Post(f func()) {
	p.worker.Post(f)
}

func (p *Engine) AfterPost(duration time.Duration, f func()) {
	p.worker.AfterPost(duration, f)
}

func (p *Engine) RegisterNetWorkEvent(onNew, onClose func(conn network.Session)) {
	p.networkMgr.RegisterEvent(onNew, onClose)
}

func (p *Engine) RegisterSessionMsgHandler(cb interface{}) {
	p.networkMgr.RegisterSessionMsgHandler(cb)
}

func (p *Engine) Worker() service.Worker {
	return p.worker
}
