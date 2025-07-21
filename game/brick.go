package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	brickSize  = firefly.Size{W: 16, H: 16}
	brickStyle = firefly.Style{FillColor: firefly.ColorOrange}
)

type Brick struct {
	pos    firefly.Point
	health int
}

func newBrick(x, y int) *Brick {
	return &Brick{
		pos:    firefly.Point{X: x, Y: y},
		health: 3,
	}
}

func (b Brick) render() {
	firefly.DrawRect(b.pos, brickSize, brickStyle)
}
