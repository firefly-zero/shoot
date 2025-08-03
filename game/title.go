package game

import "github.com/firefly-zero/firefly-go/firefly"

type Title struct {
	msg string
	ttl int
}

func setTitle(msg string) {
	title = &Title{
		msg: msg,
		ttl: 180,
	}
}

func (t *Title) update() {
	t.ttl--
	if t.ttl <= 0 {
		openHub()
	}
}

func (t Title) render() {
	x := (firefly.Width - font.LineWidth(t.msg)) / 2
	firefly.DrawText(t.msg, font, firefly.Point{X: x, Y: 80}, firefly.ColorBlack)
}
