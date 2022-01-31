package snake

import (
	"fmt"
	"sync"
	"time"
)

type Key string

type Positions struct {
	Player   *Snake
	FruitObj *Fruit
}

// Loop is a main game loop contains functions to handle player's actions and send to renderer.
func Loop(f Field, stop chan struct{}, kCh chan Key, pos chan Positions, wg *sync.WaitGroup) {
	fmt.Println("Game loop had started")
	s, err := NewSnake(f, Point{0, 0})
	fr := NewFruit(f)
	if err != nil {
		panic(err)
		return
	}

	score := 0
	d := time.Second / 2
	var k, key Key

	for {
		select {
		case <-stop:
			break
		case k = <-kCh:
			// Key buffer.
			key = k
		default:
			if s.eat(&fr) {
				score++
			}
			if s.die() {
				fmt.Println("End game")
				stop <- struct{}{}
				wg.Done()
			}
			idleMove(&s, key, d)
			sendUpdate(&s, &fr, pos)
		}
	}
}

func move(s *Snake, k Key) {
	fmt.Println(k)
	switch k {
	case "ArrowUp":
		s.Move(Point{s.Point.X, s.Point.Y - 1})
	case "ArrowDown":
		s.Move(Point{s.Point.X, s.Point.Y + 1})
	case "ArrowLeft":
		s.Move(Point{s.Point.X - 1, s.Point.Y})
	case "ArrowRight":
		s.Move(Point{s.Point.X + 1, s.Point.Y})
	}
}

func idleMove(s *Snake, k Key, t time.Duration) {
	fmt.Println(s.Point.X, s.Point.Y)
	move(s, k)
	time.Sleep(t)
}

func sendUpdate(s *Snake, f *Fruit, ch chan Positions) {
	fmt.Println("Update")
	ch <- Positions{s, f}
}
