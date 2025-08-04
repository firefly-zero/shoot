package game

import "github.com/firefly-zero/firefly-go/firefly"

var (
	font        firefly.Font
	hub         bool
	projectiles *Projectiles
	enemies     *Enemies
	players     *Set[Player]
	level       *Level
	score       *Score
	title       *Title
)

func Boot() {
	font = firefly.LoadFile("font", nil).Font()
	openHub()
}

func Update() {
	if title != nil {
		title.update()
		return
	}
	projectiles.update()
	enemies.update()
	for _, player := range players.iter() {
		if player == nil {
			continue
		}
		player.update()
	}
}

func Render() {
	firefly.ClearScreen(firefly.ColorWhite)
	if title != nil {
		title.render()
		return
	}
	level.render()
	score.render()
	projectiles.render()
	enemies.render()
	for _, player := range players.iter() {
		if player == nil {
			continue
		}
		player.render()
	}
}

func openHub() {
	hub = true
	resetGame()
}

func resetGame() {
	title = nil
	score = newScore()
	projectiles = &Projectiles{items: newSet[Projectile]()}
	enemies = newEnemies()
	level = loadLevel()
	players = loadPlayers()
}
