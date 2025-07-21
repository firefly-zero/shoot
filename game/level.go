package game

import "github.com/firefly-zero/firefly-go/firefly"

type Level struct {
	bricks *List[*Brick]
}

func loadLevel() *Level {
	var bricks *List[*Brick]
	file := firefly.LoadFile("lvl1", nil)
	x := 0
	y := 0
	for _, c := range file.Raw {
		switch c {
		case '\n':
			x = 0
			y += brickSize.H
		case '.', ' ':
			x += brickSize.W
		case '#':
			bricks = bricks.prepend(newBrick(x, y))
			x += brickSize.W
		}
	}
	return &Level{bricks: bricks}
}

func (l Level) render() {
	bricks := l.bricks
	for bricks != nil {
		brick := bricks.item
		bricks = bricks.next
		brick.render()
	}
}
