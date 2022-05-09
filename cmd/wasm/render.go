package main

import (
	"github.com/je09/GoSnakeGo/internal/snake"
	"sync"
)

// Renders game frames when necessary.
func render(pos chan snake.Positions, stop chan struct{}, wg *sync.WaitGroup) {
	var score int
	for {
		select {
		case pos := <-pos:
			// Clear game field.
			ctx.Call("clearRect", 0, 0, width*10+10, height*10+10)
			drawSnake(pos.Player)
			drawFruit(pos.FruitObj)
			score = pos.Player.Length()
		case <-stop:
			window.Call("removeEventListener", "keydown", keyboardListener)
			wg.Done()
			drawGameOver(score)
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

func drawGameOver(score int) {
	ctx.Set("font", "30px Arial")
	ctx.Set("textAlign", "center")
	ctx.Call("fillText", "Game Over", height*size/2, width*size/2)
}
