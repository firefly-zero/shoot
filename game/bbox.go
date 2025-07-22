package game

import "github.com/firefly-zero/firefly-go/firefly"

// Bounding box. Used for collision detection.
type BBox struct {
	firefly.Point
	firefly.Size
}

func (b BBox) Collides(c BBox) bool {
	if b.X+b.W <= c.X {
		return false
	}
	if b.X >= c.X+c.W {
		return false
	}
	if b.Y+b.H <= c.Y {
		return false
	}
	if b.Y >= c.Y+c.H {
		return false
	}
	return true

}
