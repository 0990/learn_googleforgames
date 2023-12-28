package game

import (
	coresdk "agones.dev/agones/pkg/sdk"
	sdk "agones.dev/agones/sdks/go"
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"
)

type AgoneSDK interface {
	Run()
	Ready()
	Shutdown()
	OnAllocate(func(users []string))
}

func NewAgoneSDK(dummy bool) (AgoneSDK, error) {
	if dummy {
		return &DummyAgones{}, nil
	}
	return newAgones()
}

type DummyAgones struct {
}

func (d *DummyAgones) Run() {

}

func (d *DummyAgones) Ready() {

}

func (d *DummyAgones) Shutdown() {

}

func (a *DummyAgones) OnAllocate(f func(users []string)) {
}

type Agones struct {
	sdk        *sdk.SDK
	onAllocate func(users []string)
}

func newAgones() (*Agones, error) {
	s, err := sdk.NewSDK()
	if err != nil {
		return nil, err
	}

	return &Agones{
		sdk: s,
	}, nil
}

func (a *Agones) OnAllocate(f func(users []string)) {
	a.onAllocate = f
}

func (a *Agones) Run() {
	a.watchGameServerEvents()
	go a.doHealth(context.Background())
}

func (a *Agones) doHealth(ctx context.Context) {
	tick := time.Tick(2 * time.Second)
	for {
		log.Printf("Health Ping")
		err := a.sdk.Health()
		if err != nil {
			log.Fatalf("Could not send health ping, %v", err)
		}
		select {
		case <-ctx.Done():
			log.Print("Stopped health pings")
			return
		case <-tick:
		}
	}
}

func (a *Agones) Ready() {
	err := a.sdk.Ready()
	if err != nil {
		log.Fatalf("Could not send Ready, %v", err)
	}
}

func (a *Agones) Shutdown() {
	err := a.sdk.Shutdown()
	if err != nil {
		log.Fatalf("Could not send Shutdown, %v", err)
	}
}

func (a *Agones) watchGameServerEvents() {
	gs, err := a.sdk.GameServer()
	if err != nil {
		log.Fatalf("Could not get Game Server, %v", err)
	}

	lastState := gs.Status.State

	err = a.sdk.WatchGameServer(func(gs *coresdk.GameServer) {
		j, err := json.Marshal(gs)
		if err != nil {
			log.Fatalf("error mashalling GameServer to JSON: %v", err)
		}
		log.Printf("GameServer Event: %s \n", string(j))

		if lastState != gs.Status.State {
			log.Printf("GameServer State change: %s -> %s \n", lastState, gs.Status.State)
			lastState = gs.Status.State
		}

		annoUsers, ok := gs.ObjectMeta.Annotations["users"]
		if ok {
			log.Printf("watchGameServerEvents users: %v \n", annoUsers)
			users := strings.Split(annoUsers, ";")
			if a.onAllocate != nil {
				a.onAllocate(users)
			}
		}
	})

	if err != nil {
		log.Fatalf("Could not watch Game Server events, %v", err)
	}
}
