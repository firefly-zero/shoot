package game

type Projectiles struct {
	items *Set[Projectile]
}

func (ps *Projectiles) update() {
	for i, p := range ps.items.iter() {
		if p == nil {
			continue
		}
		keep := p.update()
		if !keep {
			ps.items.remove(i)
		}
	}
}

func (ps Projectiles) render() {
	for _, p := range ps.items.iter() {
		if p == nil {
			continue
		}
		p.render()
	}
}
