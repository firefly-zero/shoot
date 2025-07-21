package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

const playerD = 16

type Player struct {
	peer firefly.Peer
	pad  *firefly.Pad
	pos  firefly.Point
}

func loadPlayers() []*Player {
	peers := firefly.GetPeers().Slice()
	players := make([]*Player, len(peers))
	for i, peer := range peers {
		players[i] = &Player{
			peer: peer,
			pos: firefly.Point{
				X: 60 + 10*i,
				Y: 60 + 10*i,
			},
		}
	}
	return players
}

func (p *Player) update() {
	pad, touched := firefly.ReadPad(p.peer)
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
	bricks := &level.bricks
	for bricks != nil {
		brick := bricks.head
		bricks = bricks.tail
		newPos = collideBrickPlayer(oldPos, newPos, brick)
	}
	return newPos
}

func collideBrickPlayer(oldPos, newPos firefly.Point, brick Brick) firefly.Point {
	if isCollidingBrickPlayer(newPos, brick) {
		return oldPos
	} else {
		return newPos
	}
}

func isCollidingBrickPlayer(pos firefly.Point, brick Brick) bool {
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
