package game

type Level struct {
	bricks List[Brick]
}

func loadLevel() Level {
	bricks := List[Brick]{head: newBrick(20, 20)}
	bricks = bricks.prepend(newBrick(40, 40))
	bricks = bricks.prepend(newBrick(200, 20))
	return Level{bricks: bricks}
}

func (l Level) render() {
	bricks := &l.bricks
	for bricks != nil {
		brick := bricks.head
		bricks = bricks.tail
		brick.render()
	}
}
