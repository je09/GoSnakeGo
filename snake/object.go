package snake

import (
	"errors"
	"math/rand"
	"time"
)

// Snake is a node, contains one of the snake elements
// And a reference to another node if exists.
type Snake struct {
	Point Point
	// TODO: make last element reference.
	Next *Snake
	// TODO: get rid of the field inside of this struct
	field Field
}

func NewSnake(f Field, p Point) (Snake, error) {
	if bordersOut(f, p) {
		return Snake{}, errors.New(errorPointArea)
	}
	return Snake{p, nil, f}, nil
}

// Move receives coordinates to move snake to and moves
// all the snake nodes to new location.
func (s *Snake) Move(p Point) {
	if bordersOut(s.field, p) {
		s.die()
		return
	}

	// Go until we have nodes
	for {
		// Temporary point of previous node
		tP := s.Point
		s.Point = p
		p = tP

		if s.Next != nil {
			s.Next.Point = tP
		}

		// Works until the last element of the snake.
		if s.Next == nil {
			break
		}

		s = s.Next
	}
}

// NewNode spawns new node of snake, when it ate a fruit.
// Note that if there is no other places to spawn a new node except
// the place where another node stands, it'll spawn a node here.
func (s *Snake) NewNode() {
	for {
		if s.Next == nil {
			break
		}

		s = s.Next
	}
	var newPoint Point

	switch {
	// There is a space on the right side.
	case s.Point.X < s.field.maxX:
		newPoint.X = s.Point.X + s.field.CellSize
		newPoint.Y = s.Point.Y
	// There is a space on the left side.
	case s.Point.X >= 0:
		newPoint.X = s.Point.X - s.field.CellSize
		newPoint.Y = s.Point.Y
	// There is a space bellow.
	case s.Point.Y < s.field.maxY:
		newPoint.X = s.Point.X
		newPoint.Y = s.Point.Y + s.field.CellSize
	// There is a space above.
	case s.Point.Y >= 0:
		newPoint.X = s.Point.X
		newPoint.Y = s.Point.Y - s.field.CellSize
	}
	newS, _ := NewSnake(s.field, newPoint)
	s.Next = &newS
}

// Checks if a snake can eat a fruit on the same cell as the head.
// If true, then snake grows and fruit spawns in a new location.
func (s *Snake) eat(f *Fruit) bool {
	if s.Point == f.Point {
		s.NewNode()
		f.Spawn(s.field)
		return true
	}

	return false
}

// TODO: if snake eats it's tail it dies.
// If snake goes outside the field returns true.
func (s *Snake) die() bool {
	return bordersOut(s.field, s.Point)
}

// Length returns how many nodes does snake have.
func (s *Snake) Length() int {
	i := 0
	for {
		i++
		if s.Next == nil {
			break
		}

		s = s.Next
	}

	return i
}

type Fruit struct {
	Point Point
}

func NewFruit(f Field) Fruit {
	// To create a new fruit it has to be somewhere, so let it be outside the game field.
	fr := Fruit{Point{-1, -1}}
	fr.Spawn(f)
	return fr
}

func (f *Fruit) Spawn(field Field) {
	rand.Seed(time.Now().Unix())
	p := Point{
		X: rand.Intn(field.maxX),
		Y: rand.Intn(field.maxY),
	}

	f.Point = p
}

type Point struct {
	X int
	Y int
}

type Field struct {
	maxX     int
	maxY     int
	CellSize int
}

func NewField(mX int, mY int, size int) Field {
	return Field{mX, mY, size}
}

func bordersOut(f Field, p Point) bool {
	return p.X > f.maxX || p.Y > f.maxY || p.X < 0 || p.Y < 0
}
