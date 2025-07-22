package game

import "github.com/firefly-zero/firefly-go/firefly"

type Level struct {
	bricks *Set[Brick]
}

func loadLevel() *Level {
	bricks := newSet[Brick]()
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
		case '#', 'x', 'X':
			bricks.add(newBrick(x, y))
			x += brickSize.W
		}
	}
	return &Level{bricks: bricks}
}

func (l Level) render() {
	bricks := l.bricks.iter()
	for {
		brick := bricks.next()
		if brick == nil {
			break
		}
		brick.render()
	}
}
