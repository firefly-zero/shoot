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
	items := es.items.iter()
	for {
		p := items.next()
		if p == nil {
			break
		}
		keep := p.update()
		if !keep {
			items.remove()
		}
	}
}

func (es *Enemies) spawnEnemy() {
	var pos firefly.Point
	switch firefly.GetRandom() % 4 {
	case 0:
		pos = firefly.Point{X: -10, Y: -10}
	case 1:
		pos = firefly.Point{X: firefly.Width + 10, Y: -10}
	case 2:
		pos = firefly.Point{X: 10, Y: firefly.Height + 10}
	default:
		pos = firefly.Point{X: firefly.Width + 10, Y: firefly.Height + 10}
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
	items := es.items.iter()
	for {
		p := items.next()
		if p == nil {
			break
		}
		p.render()
	}
}
