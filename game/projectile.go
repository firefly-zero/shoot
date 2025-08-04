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

func (p *Projectile) update() bool {
	p.age += 1.
	p.pos = firefly.P(
		p.origin.X+int(p.age*p.dx),
		p.origin.Y+int(p.age*p.dy),
	)

	if !p.inBounds() {
		return false
	}
	bbox := p.bbox()

	for i, brick := range level.bricks.iter() {
		if brick == nil {
			continue
		}
		if bbox.collides(brick.bbox()) {
			brick.health -= p.dmg
			if brick.health <= 0 {
				level.bricks.remove(i)
			}
			return false
		}
	}

	for _, player := range players.iter() {
		if player == nil {
			continue
		}
		if bbox.collides(player.bbox()) {
			player.health -= p.dmg
			if player.health <= 0 {
				dropDeadPlayers()
			}
			return false
		}
	}

	for i, enemy := range enemies.items.iter() {
		if enemy == nil {
			continue
		}
		if bbox.collides(enemy.bbox()) {
			enemy.health -= p.dmg
			if enemy.health <= 0 {
				enemies.items.remove(i)
				score.decrement()
			}
			return false
		}
	}

	return true
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

func (p Projectile) bbox() BBox {
	return BBox{
		Point: p.pos,
		Size:  firefly.S(p.d, p.d),
	}
}

func (p Projectile) render() {
	s := firefly.Solid(firefly.ColorYellow)
	firefly.DrawCircle(p.pos, p.d, s)
}
