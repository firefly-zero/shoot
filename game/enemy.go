package game

import "github.com/firefly-zero/firefly-go/firefly"

type Enemy struct {
	id     int
	pos    firefly.Point
	d      int
	health int
}

func (e Enemy) bbox() BBox {
	return BBox{
		Point: e.pos,
		Size:  firefly.Size{W: e.d, H: e.d},
	}
}

func (e *Enemy) update() bool {
	player := e.pickPlayer()
	dx := (player.pos.X + playerR - e.d/2 - e.pos.X)
	dy := (player.pos.Y + playerR - e.d/2 - e.pos.Y)
	dx = clamp(dx, -1, 1)
	dy = clamp(dy, -1, 1)
	if dx == 0 && dy == 0 {
		return false
	}

	// Collide the new coordinates with bricks.
	bbox := BBox{
		Point: firefly.Point{X: e.pos.X + dx, Y: e.pos.Y + dy},
		Size:  firefly.Size{W: e.d, H: e.d},
	}
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
