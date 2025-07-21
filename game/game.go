package game

import "github.com/firefly-zero/firefly-go/firefly"

var (
	players []Player
	level   Level
)

func Boot() {
	level = loadLevel()
	players = loadPlayers()
}

func Update() {
	for _, p := range players {
		p.update()
	}
}

func Render() {
	firefly.ClearScreen(firefly.ColorWhite)
	level.render()
	for _, p := range players {
		p.render()
	}
}
