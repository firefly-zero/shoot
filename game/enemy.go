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
		Point: firefly.P(e.pos.X+dx, e.pos.Y+dy),
		Size:  firefly.S(e.d, e.d),
	}

	// If the enemy is stuck on a brick for too long,
	// explode the enemy, damaging the brick.
	if e.stuck > 30 {
		for i, brick := range level.bricks.iter() {
			if brick == nil {
				continue
			}
			if bbox.collides(brick.bbox()) {
				brick.health -= 1
				if brick.health <= 0 {
					level.bricks.remove(i)
				}
				return false
			}
		}
	}

	// Collide the enemy with other enemies.
	bbox.Point = level.collide(e.pos, bbox)
	for _, enemy := range enemies.items.iter() {
		if enemy == nil {
			continue
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
	for _, p := range players.iter() {
		return p
	}
	return nil
}

func (e Enemy) render() {
	firefly.DrawCircle(e.pos, e.d, firefly.Style{
		FillColor: firefly.ColorRed,
	})
}
