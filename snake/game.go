package snake

import "time"

type Key string

func Loop(f Field, stop chan struct{}, kCh chan Key) {
	s, err := NewSnake(f, Point{0, 0})
	if err != nil {
		panic(err)
		return
	}

	score := 0
	d := time.Duration(1 * 60)
	var k Key

	for {
		select {
		case <-stop:
			break
		case k = <-kCh:
			move(s, k)
		}
		idleMove(s, k, d)
		if s.eat() {
			score++
		}
		if s.die() {
			break
		}
	}
}

func move(s Snake, k Key) {
	switch k {
	case "up":
		s.Move(Point{s.point.x, s.point.y - 1})
	case "down":
		s.Move(Point{s.point.x, s.point.y + 1})
	case "left":
		s.Move(Point{s.point.x - 1, s.point.y})
	case "right":
		s.Move(Point{s.point.x + 1, s.point.y})
	}
}

func idleMove(s Snake, k Key, t time.Duration) {
	move(s, k)
	time.Sleep(t)
}
