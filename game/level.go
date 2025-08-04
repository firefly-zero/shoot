package game

import "github.com/firefly-zero/firefly-go/firefly"

// When in hub, close the hub and start the real game if all letters are active.
func maybeStartGame() {
	if !hub {
		return
	}
	for _, letter := range level.letters.iter() {
		if letter == nil {
			continue
		}
		if !letter.active {
			return
		}
	}
	hub = false
	resetGame()
}

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
func (level Level) collides(box BBox) bool {
	for _, brick := range level.bricks.iter() {
		if brick == nil {
			continue
		}
		if box.collides(brick.bbox()) {
			return true
		}
	}
	return false
}

// Collide the given bounding box with static objects on the level.
func (level Level) collide(oldPos firefly.Point, b BBox) firefly.Point {
	c := BCircle{BBox: b}
	for _, brick := range level.bricks.iter() {
		if brick == nil {
			continue
		}
		c.Point = c.collide(oldPos, brick.bbox())
	}
	return c.Point
}

func (level Level) render() {
	for _, brick := range level.bricks.iter() {
		if brick == nil {
			continue
		}
		brick.render()
	}
	for _, letter := range level.letters.iter() {
		if letter == nil {
			continue
		}
		letter.render()
	}
}
