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
	Next  *Snake
	Last  *Snake
	// TODO: get rid of the field inside of this struct or make it a reference
	field  Field
	active bool
}

func NewSnake(f Field, p Point) (Snake, error) {
	if bordersOut(f, p) {
		return Snake{}, errors.New(errorPointArea)
	}
	return Snake{p, nil, nil, f, true}, nil
}

// Move receives coordinates to move snake to and moves
// all the snake nodes to new location.
func (s *Snake) Move(p Point) {
	// TODO: Make a loop instead of dying.
	if bordersOut(s.field, p) {
		s.die()
		return
	}

	// Go until we change all the nodes.
	for {
		// If a node was inactive, we do nothing 'cause it's already on the right location.
		if !s.active {
			s.active = true
			break
		}

		// Remember current point of the head node.
		// Change it to a location we're heading to (assuming this location is 1 cell away).
		// Set "current point of the head node" as a point we're heading to.
		tmp := s.Point
		s.Point = p
		p = tmp

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
	// if Last is nil, that means it's the head of the snake.
	if s.Last == nil {
		s.Last = s
	}

	// We don't use NewSnake method because we need to manually set active value.
	snake := Snake{s.Last.Point, nil, s.Last, s.field, false}
	s.Last.Next = &snake
	s.Last = &snake
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
	p := s.Point
	ss := s.Next
	if ss == nil {
		return false
	}

	for {
		if ss.Point == p && ss.active {
			return true
		}

		if ss.Next == nil {
			break
		}

		ss = ss.Next
	}

	return false
}

// Length returns how many nodes does snake have.
// As long as we don't need to know this information on the regular basis
// I think it's good enough to use a method like this.
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
