package game

import "github.com/firefly-zero/firefly-go/firefly"

var (
	projectiles *Projectiles
	players     *Set[Player]
	level       *Level
)

func Boot() {
	projectiles = &Projectiles{items: newSet[Projectile]()}
	level = loadLevel()
	players = loadPlayers()
}

func Update() {
	projectiles.update()
	iter := players.iter()
	for {
		p := iter.next()
		if p == nil {
			break
		}
		p.update()
	}
}

func Render() {
	firefly.ClearScreen(firefly.ColorWhite)
	level.render()
	projectiles.render()
	iter := players.iter()
	for {
		p := iter.next()
		if p == nil {
			break
		}
		p.render()
	}
}
