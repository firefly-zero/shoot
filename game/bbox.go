package game

import "github.com/firefly-zero/firefly-go/firefly"

// Bounding box. Used for collision detection.
type BBox struct {
	firefly.Point
	firefly.Size
}

func (b BBox) collides(c BBox) bool {
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

// Check if the boxes colllide and displace one if needed.
//
// Assume self to be the new position of a movable actor
// and the argument to be an immovable object.
func (b BBox) collide(oldPos firefly.Point, c BBox) firefly.Point {
	if !b.collides(c) {
		return b.Point
	}

	// A simple solution for a collision is to return oldPos.
	// However, it makes the bricks sticky. To make it easier to slide
	// alongs the brick edges, we need to project the new position
	// on the brick's surface.
	//
	// Maybe there is a more unified solution rather than handling
	// every brick surface explicitly but I'm not smart enough for this.

	// left surface
	right := firefly.Point{X: b.X + b.W, Y: oldPos.Y + b.H/2}
	if right.X < c.X+c.W && right.Y >= c.Y && right.Y <= c.Y+c.H {
		b.X = c.X - b.W
		return b.Point
	}

	// right surface
	left := firefly.Point{X: oldPos.X, Y: oldPos.Y + playerR}
	if left.X > c.X && left.Y >= c.Y && left.Y <= c.Y+c.H {
		b.X = c.X + c.W
		return b.Point
	}

	// top surface
	bottom := firefly.Point{X: oldPos.X + playerR, Y: oldPos.Y + b.H}
	if bottom.Y > c.Y+c.H && bottom.X <= c.X && bottom.X >= c.X+c.W {
		b.Y = c.Y - b.H
		return b.Point
	}

	// bottom surface
	top := firefly.Point{X: oldPos.X + playerR, Y: oldPos.Y + b.H}
	if top.Y > c.Y && top.X <= c.X && top.X >= c.X+c.W {
		b.Y = c.Y + c.H
		return b.Point
	}

	return oldPos

}
