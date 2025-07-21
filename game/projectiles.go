package game

type Projectiles struct {
	items *Set[Projectile]
}

func (ps *Projectiles) update() {
	items := ps.items.iter()
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

func (ps Projectiles) render() {
	items := ps.items.iter()
	for {
		p := items.next()
		if p == nil {
			break
		}
		p.render()
	}
}
