package game

import "github.com/firefly-zero/firefly-go/firefly"

type Enemies struct {
	nextID   int
	items    *Set[Enemy]
	nextWave int
	waves    int
}

func newEnemies() *Enemies {
	return &Enemies{
		items:    newSet[Enemy](),
		nextID:   1,
		nextWave: 120,
		waves:    0,
	}
}

func (es *Enemies) update() {
	if hub {
		return
	}
	if es.nextWave == 0 {
		es.nextWave = max(120-es.waves*4, 30)
		es.waves++
		es.spawnEnemy()
	} else {
		es.nextWave -= 1
	}
	for i, enemy := range es.items.iter() {
		if enemy == nil {
			continue
		}
		keep := enemy.update()
		if !keep {
			es.items.remove(i)
		}
	}
}

func (es *Enemies) spawnEnemy() {
	var pos firefly.Point
	switch firefly.GetRandom() % 4 {
	case 0:
		pos = firefly.P(-10, -10)
	case 1:
		pos = firefly.P(firefly.Width+10, -10)
	case 2:
		pos = firefly.P(10, firefly.Height+10)
	default:
		pos = firefly.P(firefly.Width+10, firefly.Height+10)
	}
	if es.items.len() < 5 {
		es.items.add(&Enemy{
			id:     es.nextID,
			pos:    pos,
			d:      8,
			health: 1,
		})
		es.nextID += 1
	}
}

func (es Enemies) render() {
	for _, enemy := range es.items.iter() {
		if enemy == nil {
			continue
		}
		enemy.render()
	}
}
