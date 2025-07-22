package game

import "github.com/firefly-zero/firefly-go/firefly"

var (
	projectiles *Projectiles
	enemies     *Enemies
	players     *Set[Player]
	level       *Level
)

func Boot() {
	projectiles = &Projectiles{items: newSet[Projectile]()}
	enemies = &Enemies{
		items:    newSet[Enemy](),
		nextID:   1,
		nextWave: 120,
	}
	level = loadLevel()
	players = loadPlayers()
}

func Update() {
	projectiles.update()
	enemies.update()
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
	enemies.render()
	iter := players.iter()
	for {
		p := iter.next()
		if p == nil {
			break
		}
		p.render()
	}
}
