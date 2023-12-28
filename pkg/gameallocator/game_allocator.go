package gameallocator

import (
	"errors"
	"github.com/0990/avatar-fight-server/pkg/matchmaker"
	"net"
)

type GameAllocator interface {
	Allocate(room *matchmaker.Match) (*net.TCPAddr, error)
}

type Type int

const (
	Goroutine Type = iota
	Agones    Type = 1
)

type Config struct {
	Type   Type
	Agones AgonesArgs
}

type AgonesArgs struct {
	EndPoint   string
	Namespace  string
	CertFile   string
	KeyFile    string
	CaCertFile string
}

func New(c Config) (GameAllocator, error) {
	switch c.Type {
	case Agones:
		return NewGameAllocatorAgones("172.21.248.145:30192", "", "certs/client.crt", "certs/client.key", "certs/ca.crt")
	case Goroutine:
		return NewGameAllocatorGoroutine(), nil
	default:
		return nil, errors.New("invalid game allocator mode")
	}
}
