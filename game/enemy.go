package game

import "github.com/firefly-zero/firefly-go/firefly"

type Enemy struct {
	pos    firefly.Point
	d      int
	health int
}

func (e *Enemy) update() bool {
	player := e.pickPlayer()
	dx := (player.pos.X + playerR - e.pos.X) / 10
	dy := (player.pos.Y + playerR - e.pos.Y) / 10
	dx = clamp(dx, -2, 2)
	dy = clamp(dy, -2, 2)
	if dx == 0 && dy == 0 {
		return false
	}
	dp := firefly.Point{X: dx, Y: dy}
	e.pos = e.pos.Add(dp)
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
