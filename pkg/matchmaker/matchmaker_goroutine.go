package matchmaker

import (
	"time"
)

type MakerGoroutine struct {
	users map[uint64]time.Time

	joinChan  chan uint64
	leaveChan chan uint64

	watchAssign func(*Match)
}

func NewMatchMakerGoroutine() *MakerGoroutine {
	p := &MakerGoroutine{}
	p.users = make(map[uint64]time.Time)
	p.joinChan = make(chan uint64, 10)
	p.leaveChan = make(chan uint64, 10)
	return p
}

func (p *MakerGoroutine) Run() {
	go p.run()
}

func (p *MakerGoroutine) run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			p.match()
		case userId := <-p.joinChan:
			p.users[userId] = time.Now()
		case userId := <-p.leaveChan:
			delete(p.users, userId)
		}
	}
}

func (p *MakerGoroutine) Join(userId uint64) error {
	p.joinChan <- userId
	return nil
}

func (p *MakerGoroutine) Leave(userId uint64) error {
	p.leaveChan <- userId
	return nil
}

func (p *MakerGoroutine) WatchMatch(f func(*Match)) {
	p.watchAssign = f
}

func (p *MakerGoroutine) match() {
	matchRooms := p.matchSimple()
	if p.watchAssign != nil {
		for _, v := range matchRooms {
			p.watchAssign(v)
		}
	}
}

func (p *MakerGoroutine) matchSimple() []*Match {
	result := make([]*Match, 0)

	maxCount := 5
	minCount := 2
	waitSec := 3

	for currUser, currJoinTime := range p.users {
		var matchSuccess bool

		mt := &Match{}
		mt.init()
		mt.add(currUser, currJoinTime)

		for otherUser, otherJoinTime := range p.users {
			if currUser == otherUser {
				continue
			}

			if mt.getCount() < maxCount {
				mt.add(otherUser, otherJoinTime)
			}

			if mt.getCount() >= maxCount {
				matchSuccess = true
				goto MatchSuccess
			}
		}

		if time.Now().Unix() >= currJoinTime.Unix()+int64(waitSec) {
			if mt.getCount() >= minCount {
				matchSuccess = true
				goto MatchSuccess
			}
		}

	MatchSuccess:
		if matchSuccess {
			result = append(result, mt)
			for userId, _ := range mt.Users {
				delete(p.users, userId)
			}
		}
	}

	return result
}
