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
	p.pos = firefly.Point{
		X: p.origin.X + int(p.age*p.dx),
		Y: p.origin.Y + int(p.age*p.dy),
	}

	if !p.inBounds() {
		return false
	}
	bbox := p.bbox()

	bricks := level.bricks.iter()
	for {
		brick := bricks.next()
		if brick == nil {
			break
		}
		if bbox.collides(brick.bbox()) {
			brick.health -= p.dmg
			if brick.health <= 0 {
				bricks.remove()
			}
			return false
		}
	}

	players := players.iter()
	for {
		player := players.next()
		if player == nil {
			break
		}
		if bbox.collides(player.bbox()) {
			player.health -= p.dmg
			if player.health <= 0 {
				dropDeadPlayers()
			}
			return false
		}
	}

	enemies := enemies.items.iter()
	for {
		enemy := enemies.next()
		if enemy == nil {
			break
		}
		if bbox.collides(enemy.bbox()) {
			enemy.health -= p.dmg
			if enemy.health <= 0 {
				enemies.remove()
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
		Size:  firefly.Size{W: p.d, H: p.d},
	}
}

func (p Projectile) render() {
	s := firefly.Style{FillColor: firefly.ColorYellow}
	firefly.DrawCircle(p.pos, p.d, s)
}
