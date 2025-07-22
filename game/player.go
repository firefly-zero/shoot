package game

import (
	"math"

	"github.com/firefly-zero/firefly-go/firefly"
)

const (
	playerD     = 16
	playerR     = playerD / 2
	bulletD     = 4
	bulletSpeed = 2.
	maxHealth   = 4
)

type Player struct {
	peer   firefly.Peer
	pad    *firefly.Pad
	btns   firefly.Buttons
	pos    firefly.Point
	color  firefly.Color
	health int
}

func loadPlayers() *Set[Player] {
	peers := firefly.GetPeers().Slice()
	players := newSet[Player]()
	for i, peer := range peers {
		players.add(&Player{
			peer:   peer,
			pos:    placePlayer(i),
			health: maxHealth,
			color:  pickPlayerColor(i),
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
			X: p.pos.X + playerR - bulletD/2,
			Y: p.pos.Y + playerR - bulletD/2,
		}
		bullet := &Projectile{
			d:   bulletD,
			dmg: 1,
		}
		if justPressed.S {
			bullet.dy = bulletSpeed
			origin.Y += playerR + bulletD/2
		} else if justPressed.N {
			bullet.dy = -bulletSpeed
			origin.Y -= playerR + bulletD/2
		}
		if justPressed.W {
			bullet.dx = -bulletSpeed
			origin.X -= playerR + bulletD/2
		} else if justPressed.E {
			bullet.dx = bulletSpeed
			origin.X += playerR + bulletD/2
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
	{
		angle := firefly.Radians(2 * math.Pi * float32(p.health) / maxHealth)
		s := firefly.Style{
			FillColor: p.color,
		}
		firefly.DrawSector(p.pos, playerD, firefly.Radians(0), angle, s)
	}

	s := firefly.Style{
		StrokeColor: p.color,
		StrokeWidth: 1,
	}
	firefly.DrawCircle(p.pos, playerD, s)
}

// Pick a random starting position for a player.
func placePlayer(quadrant int) firefly.Point {
	for {
		x := int(firefly.GetRandom() % (firefly.Width/2 - playerD))
		y := int(firefly.GetRandom() % (firefly.Height/2 - playerD))
		if quadrant == 1 || quadrant == 2 {
			x += firefly.Width / 2
		}
		if quadrant == 1 || quadrant == 3 {
			y += firefly.Height / 2
		}
		p := firefly.Point{X: x, Y: y}
		if !isCollidingBricksPlayer(p) {
			return p
		}
	}
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

// Check if the player movement collides with a brick and adjust the new coordinates.
func collideBrickPlayer(oldPos, newPos firefly.Point, brick *Brick) firefly.Point {
	b := BBox{
		Point: newPos,
		Size:  firefly.Size{W: playerD, H: playerD},
	}
	return b.Collide(oldPos, brick.bbox())
}

// Check if the player at the given position collides with any brick.
func isCollidingBricksPlayer(pos firefly.Point) bool {
	bricks := level.bricks.iter()
	for {
		brick := bricks.next()
		if brick == nil {
			break
		}
		if isCollidingBrickPlayer(pos, brick) {
			return true
		}
	}
	return false
}

// Check if the given brick collides with the player at the given position
func isCollidingBrickPlayer(pos firefly.Point, brick *Brick) bool {
	b := BBox{
		Point: pos,
		Size:  firefly.Size{W: playerD, H: playerD},
	}
	return b.Collides(brick.bbox())
}

func pickPlayerColor(i int) firefly.Color {
	switch i {
	case 0:
		return firefly.ColorDarkBlue
	case 1:
		return firefly.ColorCyan
	case 2:
		return firefly.ColorBlue
	default:
		return firefly.ColorLightBlue
	}
}
