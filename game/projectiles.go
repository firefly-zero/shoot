package game

type Projectiles struct {
	items *List[*Projectile]
}

func (ps *Projectiles) update() {
	items := ps.items
	for items != nil {
		p := items.item
		p.update()
		if !p.inBounds() {
			items.remove()
		}
		items = items.next
	}
}

func (ps Projectiles) render() {
	items := ps.items
	for items != nil {
		items.item.render()
		items = items.next
	}
}
