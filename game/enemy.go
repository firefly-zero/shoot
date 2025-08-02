package game

import "github.com/firefly-zero/firefly-go/firefly"

type Enemy struct {
	id     int
	pos    firefly.Point
	d      int
	health int
	stuck  int
}

func (e Enemy) bbox() BBox {
	return BBox{
		Point: e.pos,
		Size:  firefly.Size{W: e.d, H: e.d},
	}
}

func (e *Enemy) update() bool {
	player := e.pickPlayer()
	if player == nil {
		return false
	}

	dx := (player.pos.X + playerR - e.d/2 - e.pos.X)
	dy := (player.pos.Y + playerR - e.d/2 - e.pos.Y)
	dx = clamp(dx, -1, 1)
	dy = clamp(dy, -1, 1)
	if dx == 0 && dy == 0 {
		return false
	}

	// Collide the new coordinates.
	bbox := BBox{
		Point: firefly.Point{X: e.pos.X + dx, Y: e.pos.Y + dy},
		Size:  firefly.Size{W: e.d, H: e.d},
	}

	// If the enemy is stuck on a brick for too long,
	// explode the enemy, damaging the brick.
	if e.stuck > 30 {
		bricks := level.bricks.iter()
		for {
			brick := bricks.next()
			if brick == nil {
				break
			}
			if bbox.collides(brick.bbox()) {
				brick.health -= 1
				if brick.health <= 0 {
					bricks.remove()
				}
				return false
			}
		}
	}

	// Collide the enemy with other enemies.
	bbox.Point = level.collide(e.pos, bbox)
	enemies := enemies.items.iter()
	for {
		enemy := enemies.next()
		if enemy == nil {
			break
		}
		if enemy.id == e.id {
			continue
		}
		bbox.Point = bbox.collide(e.pos, enemy.bbox())
	}

	// If enemy collides with a player, explode.
	if bbox.collides(player.bbox()) {
		player.health -= 1
		if player.health <= 0 {
			dropDeadPlayers()
		}
		return false
	}

	// ITrack for how long the enemy is not moving.
	if e.pos == bbox.Point {
		e.stuck++
	} else {
		e.stuck = 0
	}
	e.pos = bbox.Point
	return true
}

func (e Enemy) pickPlayer() *Player {
	return players.iter().next()
}

func (e Enemy) render() {
	firefly.DrawCircle(e.pos, e.d, firefly.Style{
		FillColor: firefly.ColorRed,
	})
}
