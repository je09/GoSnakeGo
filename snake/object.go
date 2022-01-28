package snake

import (
	"errors"
	"math/rand"
)

// Snake is a node, contains one of the snake elements
// And a reference to another node if exists.
type Snake struct {
	point Point
	next  *Snake
	field Field
	// If true, show this node on the field.
	// We need it because of the situations when a new node needs to be spawned, but there is no space left.
	show bool
}

func bordersOut(f Field, p Point) bool {
	return p.x > f.maxX || p.y > f.maxY || p.x < 0 || p.y < 0
}

func NewSnake(f Field, p Point) (Snake, error) {
	if bordersOut(f, p) {
		return Snake{}, errors.New(errorPointArea)
	}
	return Snake{p, nil, f, true}, nil
}

// Move receives coordinates to move snake to and moves
// all the snake nodes to new location.
func (s *Snake) Move(p Point) {
	if bordersOut(s.field, p) {
		panic(errorMoveArea)
		return
	}

	ss := s

	// Go until we have nodes
	for ss.next != nil {
		// Temporary point of previous node
		tP := ss.point
		ss.point = p
		p = tP
		// TODO: Think, do we need this if?
		if ss.next != nil {
			ss = ss.next
		}
	}
}

// NewNode spawns new node of snake, when it ate a fruit.
// Note that if there is no other places to spawn a new node except
// the place where another node stands, it'll spawn a node here.
func (s *Snake) NewNode() {
	lastS := s
	for lastS.next != nil {
		lastS = s.next
	}
	var newPoint Point

	switch {
	// There is a space bellow. \/
	case lastS.point.y < lastS.field.maxY:
		{
			newPoint.x = lastS.point.x
			newPoint.y = lastS.point.y + 1
			break
		}
	// There is a space on the right side. >
	case lastS.point.x < lastS.field.maxX:
		{
			newPoint.x = lastS.point.x + 1
			newPoint.y = lastS.point.y
			break
		}
	// There is a space above. /\
	case lastS.point.y >= 0:
		{
			newPoint.x = lastS.point.x
			newPoint.y = lastS.point.y - 1
			break
		}
	// There is a space on the left side. <
	case lastS.point.x >= 0:
		{
			newPoint.x = lastS.point.x - 1
			newPoint.y = lastS.point.y
			break
		}
	}
	newS, _ := NewSnake(lastS.field, newPoint)
	lastS.next = &newS
}

// Checks if a snake can eat a fruit on the same square as the head.
// If true, then snake grows and fruit spawns in a new location.
func (s *Snake) eat(f *Fruit) bool {
	if s.point == f.point {
		s.NewNode()
		f.Spawn(s.field)
		return true
	}

	return false
}

// If snake goes outside the field it might die. =(
func (s *Snake) die() bool {
	return bordersOut(s.field, s.point)
}

type Fruit struct {
	point Point
}

func (f *Fruit) Spawn(field Field) {
	p := Point{
		x: rand.Intn(field.maxX),
		y: rand.Intn(field.maxY),
	}

	f.point = p
}

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

type Field struct {
	maxX int
	maxY int
}

func NewField(mX int, mY int) Field {
	return Field{mX, mY}
}
