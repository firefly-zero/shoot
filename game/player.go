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

// Remove players with zero health.
func dropDeadPlayers() {
	changed := false
	for i, player := range players.iter() {
		if player == nil {
			continue
		}
		if player.health <= 0 {
			changed = true
			players.remove(i)
		}
	}
	if players.empty() {
		setTitle("everyone is dead")
	} else if changed {
		score.decreaseTo(5 * players.len())
	}
}

func iAmAlive() bool {
	me := firefly.GetMe()
	for _, player := range players.iter() {
		if player == nil {
			continue
		}
		if player.peer == me {
			return true
		}
	}
	return false
}

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

func (p Player) bbox() BBox {
	return BBox{
		Point: p.pos,
		Size:  firefly.S(playerD, playerD),
	}
}

func (p *Player) update() {
	btns := firefly.ReadButtons(p.peer)
	pad, touched := firefly.ReadPad(p.peer)

	p.handleButtons(btns)
	p.btns = btns

	if touched {
		if p.pad != nil {
			dx := (pad.X - p.pad.X) / 20
			dx = clamp(dx, -10, 10)
			dy := (pad.Y - p.pad.Y) / 20
			dy = clamp(dy, -10, 10)

			newX := clamp(p.pos.X+dx, 0, firefly.Width-playerD)
			newY := clamp(p.pos.Y-dy, 0, firefly.Height-playerD)

			b := BBox{
				Point: firefly.P(newX, newY),
				Size:  firefly.S(playerD, playerD),
			}
			b.Point = level.collide(p.pos, b)
			for _, letter := range level.letters.iter() {
				if letter == nil {
					continue
				}
				if b.collides(letter.bbox()) {
					letter.active = true
					maybeStartGame()
				}
			}
			p.pos = b.Point
		}
		p.pad = &pad
	} else {
		p.pad = nil
	}
}

func (p *Player) handleButtons(btns firefly.Buttons) {
	justPressed := btns.JustPressed(p.btns)
	if justPressed.AnyPressed() {
		origin := firefly.P(
			p.pos.X+playerR-bulletD/2,
			p.pos.Y+playerR-bulletD/2,
		)
		bullet := &Projectile{
			d:   bulletD,
			dmg: 1,
		}
		if hub {
			bullet.dmg = 0
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
		p := firefly.P(x, y)
		if !isCollidingBricksPlayer(p) {
			return p
		}
	}
}

// Check if the player at the given position collides with any brick.
func isCollidingBricksPlayer(pos firefly.Point) bool {
	b := BBox{
		Point: pos,
		Size:  firefly.S(playerD, playerD),
	}
	return level.collides(b)
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
