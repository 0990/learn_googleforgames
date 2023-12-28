package game

import "fmt"

type PanicInfo struct {
	Tip string
}

func (p *PanicInfo) ToString() string {
	return fmt.Sprintf("游戏强制结束，原因:%s", p.Tip)
}

func PanicStopGame(tip string) {
	panic(&PanicInfo{
		Tip: tip,
	})
}
