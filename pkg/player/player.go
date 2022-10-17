package player

import (
	"fmt"
)

var (
	MoveSpeedX float64 = 5
	MoveSpeedY float64 = 5
)

type Player struct {
	PosX float64
	PosY float64
	R    float64
}

func (p *Player) String() string {
	return fmt.Sprintf("X: %v; Y: %v", p.PosX, p.PosY)
}

func NewPlayer() *Player {
	return &Player{
		PosX: 1,
		PosY: 1,
		R:    10,
	}
}
