package game

import "github.com/firefly-zero/firefly-go/firefly"

var (
	projectiles *Projectiles
	players     []*Player
	level       *Level
)

func Boot() {
	projectiles = &Projectiles{}
	level = loadLevel()
	players = loadPlayers()
}

func Update() {
	projectiles.update()
	for _, p := range players {
		p.update()
	}
}

func Render() {
	firefly.ClearScreen(firefly.ColorWhite)
	level.render()
	projectiles.render()
	for _, p := range players {
		p.render()
	}
}
