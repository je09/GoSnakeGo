package snake

import "testing"

func TestFruit_Spawn(t *testing.T) {
	startPoint := Point{0, 0}
	field := Field{10, 10}

	f := Fruit{point: startPoint}
	f.Spawn(field)
	if f.point == startPoint {
		t.Errorf("Expected new point, got x: %d. y: %d", f.point.x, f.point.y)
	}
	oldPoint := f.point
	f.Spawn(field)
	if f.point == oldPoint {
		t.Errorf("Expected new point, got x: %d. y: %d", f.point.x, f.point.y)
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
		{
			name: "down space test", args: testArgs{p: Point{0, 0}, f: Field{2, 2}}, want: Point{0, 1},
		},
		{
			name: "right space test", args: testArgs{p: Point{0, 2}, f: Field{2, 2}}, want: Point{1, 2},
		},
		{
			name: "up space test", args: testArgs{p: Point{2, 2}, f: Field{2, 2}}, want: Point{2, 1},
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
			if s.next == nil {
				t.Errorf("Expected next pointer not nil got nil")
			}
			if s.next.point != tt.want {
				t.Errorf("Expected next point (%d, %d) got (%d, %d)",
					tt.want.x, tt.want.y, s.next.point.x, s.next.point.y)
			}
		})
	}

}

func TestSnake_Move(t *testing.T) {

}
