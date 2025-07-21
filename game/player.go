package game

import "github.com/firefly-zero/firefly-go/firefly"

type Player struct {
	peer firefly.Peer
	pad  *firefly.Pad
	pos  firefly.Point
}

func loadPlayers() []Player {
	peers := firefly.GetPeers().Slice()
	players := make([]Player, len(peers))
	for i, peer := range peers {
		players[i] = Player{
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
	firefly.DrawCircle(p.pos, 16, s)
}
