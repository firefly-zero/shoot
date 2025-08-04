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
		pos: firefly.P(x, y),
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
	p := firefly.P(b.pos.X+4, b.pos.Y+11)
	font.Draw(b.c, p, textColor)
}
