package matchmaker

import "fmt"

type MatchMaker interface {
	Run()
	Join(userId uint64) error
	Leave(userId uint64) error
	WatchMatch(f func(*Match))
}

type Type int

const (
	Goroutine Type = iota
	OpenMatch Type = 1
)

type Config struct {
	Type      Type
	OpenMatch OpenMatchArgs
}

type OpenMatchArgs struct {
	FrontendEndPoint  string
	BackendEndPoint   string
	MatchFunctionHost string
	MatchFunctionPort int32
}

func New(c Config) (MatchMaker, error) {
	switch c.Type {
	case Goroutine:
		return NewMatchMakerGoroutine(), nil
	case OpenMatch:
		return NewMatchMakerOpenMatch(c.OpenMatch)
	default:
		return nil, fmt.Errorf("unknown matchmaker type %d", c.Type)
	}
}
