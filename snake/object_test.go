package snake

import (
	"testing"
	"time"
)

func TestFruit_Spawn(t *testing.T) {
	startPoint := Point{0, 0}
	field := Field{10, 10, 1}

	f := Fruit{Point: startPoint}
	f.Spawn(field)
	if f.Point == startPoint {
		t.Errorf("Expected new point, got x: %d. y: %d", f.Point.X, f.Point.Y)
	}

	// Because we set seed based on time.
	time.Sleep(time.Second)

	oldPoint := f.Point
	f.Spawn(field)
	if f.Point == oldPoint {
		t.Errorf("Expected new point, got x: %d. y: %d", f.Point.X, f.Point.Y)
	}
}

func TestSnake_NewNode(t *testing.T) {
	s, err := NewSnake(Field{999, 999, 1}, Point{0, 0})
	if err != nil {
		panic(err)
	}
	if s.Next != nil {
		t.Errorf("Expected next pointer to be nil on the first iteration")
	}

	for i := 0; i < 99; i++ {
		s.NewNode()
		s.Move(Point{1 + i, 0})
	}

	ss := s

	for i := 0; i < 99; i++ {
		if ss.Point == ss.Next.Point {
			t.Errorf("Expected next pointer to be different from a previous one, got: %d", ss.Next.Point)
		}
		ss = *ss.Next
	}

	s, err = NewSnake(Field{999, 999, 1}, Point{0, 0})

	for i := 0; i < 99; i++ {
		s.NewNode()
		s.Move(Point{0, 1 + i})
	}

	ss = s

	for i := 0; i < 99; i++ {
		if ss.Point == ss.Next.Point {
			t.Errorf("Expected next pointer to be different from a previous one, got: %d", ss.Next.Point)
		}
		ss = *ss.Next
	}

	if s.Length() != 100 {
		t.Errorf("Expected snake to be 100 nodes long, got: %d", s.Length())
	}
}

func TestSnake_Move(t *testing.T) {
	s, _ := NewSnake(Field{999, 999, 1}, Point{0, 0})
	s.NewNode()
	s.Move(Point{1, 0})
	for i := 0; i < 99; i++ {
		s.Move(Point{2 + i, 0})
	}
	s.NewNode()
	s.Move(Point{101, 0})

	ss := s

	for {
		if ss.Next == nil {
			break
		}
		if ss.Point == ss.Next.Point {
			t.Errorf("Expected next pointer to be different from a previous one, got: %d", ss.Next.Point)
		}
		ss = *ss.Next
	}
}

func TestSnakeMove(t *testing.T) {
	s, err := NewSnake(Field{10, 10, 1}, Point{0, 0})
	if err != nil {
		panic(err)
	}
	s.Move(Point{1, 0})
	if s.Point.X != 1 {
		t.Errorf("Expected snake to have Y=%d, got Y=%d", 1, s.Point.X)
	}

	s.NewNode()
	s.Move(Point{2, 0})
	if s.Point.X != 2 {
		t.Errorf("Expected head to have X=%d, got X=%d", 2, s.Point.X)
	}
	if s.Next.Point.X != 1 {
		t.Errorf("Expected next node to have X=%d, got X=%d", 1, s.Point.X)
	}
}

func TestSnake_eat(t *testing.T) {
	s, err := NewSnake(Field{10, 10, 1}, Point{0, 0})
	if err != nil {
		panic(err)
	}
	f := Fruit{Point{1, 0}}

	s.Move(Point{1, 0})
	s.eat(&f)

	if s.Next == nil {
		t.Errorf("Expected next node not to be nil")
	}
	// May be false alarms, cause data spawn location is random.
	if f.Point.X == 1 {
		t.Errorf("Expected fruit to change location got X=%d", f.Point.X)
	}
}

func TestSnake_die(t *testing.T) {
	//s, err := NewSnake(Field{10, 10, 1}, Point{0, 0})
	//if err != nil {
	//	panic(err)
	//}
}
