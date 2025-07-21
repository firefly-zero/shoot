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

			p.pos = firefly.Point{X: newX, Y: newY}
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
