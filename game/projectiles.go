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
		p.update()
		if !p.inBounds() {
			items.remove()
		} else {
			bricks := level.bricks.iter()
			for {
				brick := bricks.next()
				if brick == nil {
					break
				}
				if p.isCollidingBrick(brick) {
					items.remove()
					brick.health -= p.dmg
					if brick.health <= 0 {
						bricks.remove()
					}
					break
				}
			}

			players := players.iter()
			for {
				player := players.next()
				if player == nil {
					break
				}
				if p.isCollidingPlayer(player) {
					items.remove()
					player.health -= p.dmg
					if player.health <= 0 {
						players.remove()
					}
					break
				}
			}
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
