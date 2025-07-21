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
			if items.prev != nil {
				items.prev.next = items.next
			} else {
				ps.items = items.next
			}
		} else {
			bricks := level.bricks
			for bricks != nil {
				brick := bricks.item
				if p.isCollidingBrick(brick) {
					if items.prev != nil {
						items.prev.next = items.next
					} else {
						ps.items = items.next
					}
					break
				}
				bricks = bricks.next
			}
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
