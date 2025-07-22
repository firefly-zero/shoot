package game

import "github.com/firefly-zero/firefly-go/firefly"

type Level struct {
	bricks  *Set[Brick]
	letters *Set[Letter]
}

func loadLevel() *Level {
	bricks := newSet[Brick]()
	letters := newSet[Letter]()
	var fileName string
	if hub {
		playersCount := firefly.GetPeers().Len()
		fileName = "hub" + string(rune('0'+playersCount))
	} else {
		lvl := firefly.GetRandom() % 6
		fileName = "lvl" + string(rune('1'+lvl))
	}
	file := firefly.LoadFile(fileName, nil)
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
		default:
			letters.add(newLetter(c, x, y))
			x += brickSize.W
		}
	}
	return &Level{bricks: bricks, letters: letters}
}

// Check if the given bounding box collides with any static object.
func (l Level) collides(b BBox) bool {
	bricks := l.bricks.iter()
	for {
		brick := bricks.next()
		if brick == nil {
			break
		}
		if b.collides(brick.bbox()) {
			return true
		}
	}
	return false
}

// Collide the given bounding box with static objects on the level.
func (l Level) collide(oldPos firefly.Point, b BBox) firefly.Point {
	bricks := l.bricks.iter()
	for {
		brick := bricks.next()
		if brick == nil {
			break
		}
		b.Point = b.collide(oldPos, brick.bbox())
	}
	return b.Point
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
	letters := l.letters.iter()
	for {
		letter := letters.next()
		if letter == nil {
			break
		}
		letter.render()
	}
}
