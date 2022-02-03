package main

import (
	"github.com/je09/GoSnakeGo/snake"
	"sync"
)

// Renders game frames when necessary.
func render(pos chan snake.Positions, stop chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case pos := <-pos:
			// Clear game field.
			ctx.Call("clearRect", 0, 0, width*10+10, height*10+10)
			drawSnake(pos.Player)
			drawFruit(pos.FruitObj)
		case <-stop:
			wg.Done()
			break
		}
	}
}

func drawSnake(s *snake.Snake) {
	for {
		drawRect(s.Point.X*10, s.Point.Y*10, 10, "black")
		if s.Next == nil {
			break
		}
		s = s.Next
	}
}

func drawFruit(f *snake.Fruit) {
	drawRect(f.Point.X*10, f.Point.Y*10, 10, "red")
}

func drawRect(x int, y int, size int, clr string) {
	ctx.Set("fillStyle", clr)
	ctx.Call("fillRect", x, y, size, size)
}
