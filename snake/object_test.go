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
	type testArgs struct {
		p Point
		f Field
	}

	type testCase struct {
		name string
		args testArgs
		want Point
	}

	tests := []testCase{
		// TODO: Change this tests, it doesn't work anymore because of move and spawn logic changes.
		{
			name: "down space test", args: testArgs{p: Point{0, 0}, f: Field{2, 2, 1}}, want: Point{0, 1},
		},
		{
			name: "right space test", args: testArgs{p: Point{0, 2}, f: Field{2, 2, 1}}, want: Point{1, 2},
		},
		{
			name: "up space test", args: testArgs{p: Point{2, 2}, f: Field{2, 2, 1}}, want: Point{2, 1},
		},
		// Can't think of the way to test it
		//{
		//	name: "left space test", args: testArgs{p: Point{1, 0}, f: Field{1, 1}}, want: Point{0, 0},
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewSnake(tt.args.f, tt.args.p)
			if err != nil {
				panic(err)
			}
			s.NewNode()
			if s.Next == nil {
				t.Errorf("Expected next pointer not nil got nil")
			}
			if s.Next.Point != tt.want {
				t.Errorf("Expected next point (%d, %d) got (%d, %d)",
					tt.want.X, tt.want.Y, s.Next.Point.X, s.Next.Point.Y)
			}
		})
	}
}

func TestSnake_NewNode2(t *testing.T) {
	s, err := NewSnake(Field{999, 999, 1}, Point{0, 0})
	if err != nil {
		panic(err)
	}
	if s.Next != nil {
		t.Errorf("Expected next pointer to be nil on the first iteration")
	}

	for i := 0; i < 99; i++ {
		s.NewNode()
	}

	if s.Length() != 100 {
		t.Errorf("Expected snake to be 100 nodes long, got: %d", s.Length())
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
