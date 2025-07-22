package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

type Letter struct {
	c      string
	pos    firefly.Point
	active bool
}

func newLetter(c byte, x, y int) *Letter {
	return &Letter{
		c:   string(rune(c)),
		pos: firefly.Point{X: x, Y: y},
	}
}

func (b Letter) bbox() BBox {
	return BBox{
		Point: b.pos,
		Size:  brickSize,
	}
}
func (b Letter) render() {
	style := firefly.Style{
		FillColor:   firefly.ColorLightGray,
		StrokeColor: firefly.ColorGray,
		StrokeWidth: 1,
	}
	textColor := firefly.ColorDarkGray
	if b.active {
		style.FillColor = firefly.ColorDarkGreen
		style.StrokeColor = firefly.ColorGreen
		textColor = firefly.ColorWhite
	}
	firefly.DrawRect(
		b.pos,
		brickSize,
		style,
	)
	firefly.DrawText(
		b.c, font,
		b.pos.Add(firefly.Point{X: 4, Y: 11}),
		textColor,
	)
}
