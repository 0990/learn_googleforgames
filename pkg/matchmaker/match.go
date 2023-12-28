package matchmaker

import (
	"time"
)

type Match struct {
	Users      map[uint64]time.Time
	Count      int
	RobotCount int
}

func NewMatch() *Match {
	return &Match{
		Users:      make(map[uint64]time.Time),
		Count:      0,
		RobotCount: 0,
	}
}

func (p *Match) init() {
	p.Users = make(map[uint64]time.Time)
	p.Count = 0
}

func (p *Match) getCount() int {
	return p.Count + p.RobotCount
}

func (p *Match) add(userId uint64, time time.Time) {
	p.Users[userId] = time
	p.Count += 1
}

func (p *Match) addRobot() {
	p.RobotCount++
}
