package snake

import "testing"

func Test_move(t *testing.T) {
	type testArgs struct {
		s Snake
		k Key
	}

	type testCase struct {
		name string
		args testArgs
		want Point
	}

	p := Point{5, 5}
	f := Field{10, 10, 1}

	tests := []testCase{
		{
			name: "ArrowUp test", args: testArgs{Snake{Point: p, Next: nil, field: f}, "ArrowUp"}, want: Point{5, 4},
		},
		{
			name: "ArrowDown test", args: testArgs{Snake{Point: p, Next: nil, field: f}, "ArrowDown"}, want: Point{5, 6},
		},
		{
			name: "ArrowRight test", args: testArgs{Snake{Point: p, Next: nil, field: f}, "ArrowRight"}, want: Point{6, 5},
		},
		{
			name: "ArrowLeft test", args: testArgs{Snake{Point: p, Next: nil, field: f}, "ArrowLeft"}, want: Point{4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			move(&tt.args.s, tt.args.k)
			if tt.args.s.Point != tt.want {
				t.Errorf("Expected snake to have (%d, %d), got: (%d, %d)", tt.want.X, tt.want.Y, tt.args.s.Point.X, tt.args.s.Point.Y)
			}
		})
	}
}
