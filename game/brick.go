package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

var brickSize = firefly.S(16, 16)

type Brick struct {
	pos    firefly.Point
	health int
}

func newBrick(x, y int) *Brick {
	return &Brick{
		pos:    firefly.P(x, y),
		health: 4,
	}
}

func (b Brick) bbox() BBox {
	return BBox{
		Point: b.pos,
		Size:  brickSize,
	}
}

func (b Brick) render() {
	firefly.DrawRect(
		b.pos,
		brickSize,
		firefly.Solid(firefly.ColorOrange),
	)
	firefly.DrawRect(
		b.pos.Add(firefly.P(2, 2)),
		brickSize.Sub(firefly.S(4, 4)),
		firefly.Outlined(firefly.ColorWhite, 1),
	)
	b.renderCracks()
}

func (b Brick) renderCracks() {
	style := firefly.Solid(firefly.ColorWhite)
	if b.health <= 3 {
		firefly.DrawTriangle(
			firefly.P(b.pos.X, b.pos.Y+2),
			firefly.P(b.pos.X, b.pos.Y+4),
			firefly.P(b.pos.X+2, b.pos.Y+3),
			style,
		)
	}
	if b.health <= 2 {
		firefly.DrawTriangle(
			firefly.P(b.pos.X+brickSize.W, b.pos.Y+4),
			firefly.P(b.pos.X+brickSize.W, b.pos.Y+6),
			firefly.P(b.pos.X+brickSize.W-2, b.pos.Y+5),
			style,
		)
	}
	if b.health <= 1 {
		firefly.DrawTriangle(
			firefly.P(b.pos.X+2, b.pos.Y+brickSize.H),
			firefly.P(b.pos.X+6, b.pos.Y+brickSize.H),
			firefly.P(b.pos.X+4, b.pos.Y+brickSize.H-4),
			style,
		)
	}
}
