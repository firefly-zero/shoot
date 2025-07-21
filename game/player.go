package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

const (
	playerD     = 16
	bulletD     = 4
	bulletSpeed = 2.
)

type Player struct {
	peer   firefly.Peer
	pad    *firefly.Pad
	btns   firefly.Buttons
	pos    firefly.Point
	health int
}

func loadPlayers() *Set[Player] {
	peers := firefly.GetPeers().Slice()
	players := newSet[Player]()
	for i, peer := range peers {
		players.add(&Player{
			peer: peer,
			pos: firefly.Point{
				X: 60 + 30*i,
				Y: 60 + 30*i,
			},
			health: 3,
		})
	}
	return players
}

func (p *Player) update() {
	btns := firefly.ReadButtons(p.peer)
	pad, touched := firefly.ReadPad(p.peer)

	justPressed := btns.JustPressed(p.btns)
	if justPressed.AnyPressed() {
		origin := firefly.Point{
			X: p.pos.X + playerD/2 - bulletD/2,
			Y: p.pos.Y + playerD/2 - bulletD/2,
		}
		bullet := &Projectile{
			d:   bulletD,
			dmg: 1,
		}
		if justPressed.S {
			bullet.dy = bulletSpeed
			origin.Y += playerD/2 + bulletD/2
		} else if justPressed.N {
			bullet.dy = -bulletSpeed
			origin.Y -= playerD/2 + bulletD/2
		}
		if justPressed.W {
			bullet.dx = -bulletSpeed
			origin.X -= playerD/2 + bulletD/2
		} else if justPressed.E {
			bullet.dx = bulletSpeed
			origin.X += playerD/2 + bulletD/2
		}
		bullet.origin = origin
		bullet.pos = origin
		projectiles.items.add(bullet)
	}
	p.btns = btns

	if touched {
		if p.pad != nil {
			dx := (pad.X - p.pad.X) / 20
			dx = clamp(dx, -10, 10)
			dy := (pad.Y - p.pad.Y) / 20
			dy = clamp(dy, -10, 10)

			newX := clamp(p.pos.X+dx, 0, firefly.Width-playerD)
			newY := clamp(p.pos.Y-dy, 0, firefly.Height-playerD)

			newPos := firefly.Point{X: newX, Y: newY}
			p.pos = collideBricksPlayer(p.pos, newPos)
		}
		p.pad = &pad
	} else {
		p.pad = nil
	}
}

func (p *Player) render() {
	s := firefly.Style{
		StrokeColor: firefly.ColorBlue,
		StrokeWidth: 1,
	}
	firefly.DrawCircle(p.pos, playerD, s)
}

// Make sure that the new player position doesn't place the player inside a brick.
func collideBricksPlayer(oldPos, newPos firefly.Point) firefly.Point {
	bricks := level.bricks.iter()
	for {
		brick := bricks.next()
		if brick == nil {
			break
		}
		newPos = collideBrickPlayer(oldPos, newPos, brick)
	}
	return newPos
}

func collideBrickPlayer(oldPos, newPos firefly.Point, brick *Brick) firefly.Point {
	if isCollidingBrickPlayer(newPos, brick) {
		return oldPos
	} else {
		return newPos
	}
}

func isCollidingBrickPlayer(pos firefly.Point, brick *Brick) bool {
	if pos.X+playerD < brick.pos.X {
		return false
	}
	if pos.X > brick.pos.X+brickSize.W {
		return false
	}
	if pos.Y+playerD < brick.pos.Y {
		return false
	}
	if pos.Y > brick.pos.Y+brickSize.H {
		return false
	}
	return true
}
