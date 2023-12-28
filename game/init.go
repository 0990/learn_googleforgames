package game

import (
	"fmt"
	"net"
)

type App struct {
	engine     *Engine
	sessionMgr *SessionMgr
	game       *Game
	agonesSDK  AgoneSDK
}

func New(addr string, agones bool) (*App, error) {
	engine := NewEngine(addr)
	sessionMgr := newSessionMgr()

	game := newGame("1", func(g *Game) {})
	agonesSDK, err := NewAgoneSDK(!agones)
	if err != nil {
		return nil, err
	}
	agonesSDK.Run()
	app := &App{
		engine:     engine,
		sessionMgr: sessionMgr,
		game:       game,
		agonesSDK:  agonesSDK,
	}
	app.registerHandler()
	return app, nil
}

func (app *App) Run(readyCallback func(address *net.TCPAddr)) {
	app.engine.Run()
	app.agonesSDK.OnAllocate(func(users []string) {
		fmt.Printf("agones allocate:%v \n", users)
		app.game.eventProcessor.AddEvent(Event{
			Type: AgonesAllocate,
			Data: &EventAgonesAllocateData{users: users},
		})
	})

	app.agonesSDK.Ready()
	if readyCallback != nil {
		readyCallback(app.engine.ListenAddr())
	}
	app.game.Run()
	app.engine.Close()
	app.agonesSDK.Shutdown()
}
