package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

type Letter struct {
	c   string
	pos firefly.Point
}

func newLetter(c byte, x, y int) *Letter {
	return &Letter{
		c:   string(rune(c)),
		pos: firefly.Point{X: x, Y: y},
	}
}

func (b Letter) render() {
	firefly.DrawRect(
		b.pos,
		brickSize,
		firefly.Style{
			FillColor:   firefly.ColorLightGray,
			StrokeColor: firefly.ColorGray,
			StrokeWidth: 1,
		},
	)
	firefly.DrawText(
		b.c, font,
		b.pos.Add(firefly.Point{X: 4, Y: 11}),
		firefly.ColorDarkGray,
	)
}
