package game

import "github.com/firefly-zero/firefly-go/firefly"

type Projectile struct {
	origin firefly.Point
	pos    firefly.Point
	age    float32
	dx     float32
	dy     float32
	d      int
	dmg    int
}

func (p *Projectile) update() {
	p.age += 1.
	p.pos = firefly.Point{
		X: p.origin.X + int(p.age*p.dx),
		Y: p.origin.Y + int(p.age*p.dy),
	}
}

func (p Projectile) inBounds() bool {
	if p.pos.X < 0 {
		return false
	}
	if p.pos.X+p.d > firefly.Width {
		return false
	}
	if p.pos.Y < 0 {
		return false
	}
	if p.pos.Y+p.d > firefly.Height {
		return false
	}
	return true
}

func (p Projectile) render() {
	s := firefly.Style{FillColor: firefly.ColorRed}
	firefly.DrawCircle(p.pos, p.d, s)
}

func (p Projectile) isCollidingBrick(brick *Brick) bool {
	if p.pos.X+p.d <= brick.pos.X {
		return false
	}
	if p.pos.X >= brick.pos.X+brickSize.W {
		return false
	}
	if p.pos.Y+p.d <= brick.pos.Y {
		return false
	}
	if p.pos.Y >= brick.pos.Y+brickSize.H {
		return false
	}
	return true
}

func (p Projectile) isCollidingPlayer(player *Player) bool {
	if p.pos.X+p.d <= player.pos.X {
		return false
	}
	if p.pos.X >= player.pos.X+playerD {
		return false
	}
	if p.pos.Y+p.d <= player.pos.Y {
		return false
	}
	if p.pos.Y >= player.pos.Y+playerD {
		return false
	}
	return true
}
