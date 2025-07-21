package game

import "github.com/firefly-zero/firefly-go/firefly"

type Enemies struct {
	items    *Set[Enemy]
	nextWave int
}

func (es *Enemies) update() {
	if es.nextWave == 0 {
		es.nextWave = 60
		es.items.add(&Enemy{
			pos:    firefly.Point{X: 10, Y: 10},
			d:      8,
			health: 1,
		})
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
