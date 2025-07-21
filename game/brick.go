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

func (b Brick) left() int {
	return b.pos.X
}

func (b Brick) right() int {
	return b.pos.X + brickSize.W
}

func (b Brick) top() int {
	return b.pos.Y
}

func (b Brick) bottom() int {
	return b.pos.Y + brickSize.H
}

func (b Brick) render() {
	firefly.DrawRect(b.pos, brickSize, brickStyle)
}
