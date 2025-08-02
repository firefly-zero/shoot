package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

type BCircle struct {
	BBox
}

func (c BCircle) collides(b BBox) bool {
	if !c.BBox.collides(b) {
		return false
	}
	center := c.center()
	r2 := c.W * c.H / 4          // squared radius
	wl := b.X - center.X         // width left
	wr := center.X - (b.X + b.W) // width right
	ht := b.Y - center.Y         // height top
	hb := center.Y - (b.Y + b.H) // height bottom
	if wl > 0 && ht > 0 {        // top-left
		actD2 := wl*wl + ht*ht
		return actD2 < r2
	}
	if wr > 0 && ht > 0 { // top-right
		actD2 := wr*wr + ht*ht
		return actD2 < r2
	}
	if wl > 0 && hb > 0 { // bottom-left
		actD2 := wl*wl + hb*hb
		return actD2 < r2
	}
	if wr > 0 && hb > 0 { // bottom-right
		actD2 := wr*wr + hb*hb
		return actD2 < r2
	}

	return true
}

func (c BCircle) center() firefly.Point {
	return firefly.P(c.X+c.W/2, c.Y+c.H/2)
}

func (c BCircle) collide(oldPos firefly.Point, b BBox) firefly.Point {
	if !c.collides(b) {
		return c.Point
	}
	return c.collideSides(oldPos, b)
}
