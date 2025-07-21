package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	brickSize = firefly.Size{W: 16, H: 16}
)

type Brick struct {
	pos    firefly.Point
	health int
}

func newBrick(x, y int) *Brick {
	return &Brick{
		pos:    firefly.Point{X: x, Y: y},
		health: 4,
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
	{
		style := firefly.Style{FillColor: firefly.ColorOrange}
		firefly.DrawRect(b.pos, brickSize, style)
	}
	{
		style := firefly.Style{StrokeColor: firefly.ColorWhite, StrokeWidth: 1}
		firefly.DrawRect(
			b.pos.Add(firefly.Point{X: 2, Y: 2}),
			brickSize.Sub(firefly.Size{W: 4, H: 4}),
			style,
		)
	}
	b.renderCracks()
}

func (b Brick) renderCracks() {
	style := firefly.Style{FillColor: firefly.ColorWhite}
	if b.health <= 3 {
		firefly.DrawTriangle(
			firefly.Point{X: b.pos.X, Y: b.pos.Y + 2},
			firefly.Point{X: b.pos.X, Y: b.pos.Y + 4},
			firefly.Point{X: b.pos.X + 2, Y: b.pos.Y + 3},
			style,
		)
	}
	if b.health <= 2 {
		firefly.DrawTriangle(
			firefly.Point{X: b.pos.X + brickSize.W, Y: b.pos.Y + 4},
			firefly.Point{X: b.pos.X + brickSize.W, Y: b.pos.Y + 6},
			firefly.Point{X: b.pos.X + brickSize.W - 2, Y: b.pos.Y + 5},
			style,
		)
	}
	if b.health <= 1 {
		firefly.DrawTriangle(
			firefly.Point{X: b.pos.X + 2, Y: b.pos.Y + brickSize.H},
			firefly.Point{X: b.pos.X + 6, Y: b.pos.Y + brickSize.H},
			firefly.Point{X: b.pos.X + 4, Y: b.pos.Y + brickSize.H - 4},
			style,
		)
	}
}
