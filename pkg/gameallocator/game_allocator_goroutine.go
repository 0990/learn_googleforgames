package gameallocator

import (
	"fmt"
	"github.com/0990/avatar-fight-server/game"
	"github.com/0990/avatar-fight-server/pkg/matchmaker"
	"net"
	"sync"
)

type AllocatorGoroutine struct {
}

func NewGameAllocatorGoroutine() *AllocatorGoroutine {
	return &AllocatorGoroutine{}
}

func (a *AllocatorGoroutine) Allocate(room *matchmaker.Match) (*net.TCPAddr, error) {
	g, err := game.New(":0", false)
	if err != nil {
		return nil, err
	}

	var result *net.TCPAddr
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}

			fmt.Println("game end")
		}()

		g.Run(func(address *net.TCPAddr) {
			defer wg.Done()

			result = address
		})
	}()
	wg.Wait()

	return result, nil
}
