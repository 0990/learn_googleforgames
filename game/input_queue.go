package game

import (
	"fmt"
	"github.com/0990/avatar-fight-server/msg/cmsg"
	"time"
)

type InputQueue struct {
	inputs  []*cmsg.SeatInput
	frameId uint32

	tickers []int64
}

func NewInputQueue() *InputQueue {
	return &InputQueue{inputs: make([]*cmsg.SeatInput, 0, 100), frameId: 1, tickers: make([]int64, 0, 10000)}
}

func (q *InputQueue) Reset() {
	q.inputs = make([]*cmsg.SeatInput, 0, 100)
	q.frameId = 1
	q.tickers = make([]int64, 0, 10000)
}

func (q *InputQueue) Push(input *cmsg.SeatInput) {
	q.inputs = append(q.inputs, input)

}

func (q *InputQueue) Tick() {
	q.tickers = append(q.tickers, time.Now().UnixMilli())
	q.frameId++
	q.inputs = make([]*cmsg.SeatInput, 0, 100)
}

func (q *InputQueue) FrameId() uint32 {
	return q.frameId
}

func (q *InputQueue) Inputs() []*cmsg.SeatInput {
	return q.inputs
}

func (q *InputQueue) PrintTicker() {
	m := make(map[int64]int)
	for i, v := range q.tickers {
		if i+1 >= len(q.tickers) {
			break
		}
		m[q.tickers[i+1]-v]++
	}
	fmt.Println(m)
}
