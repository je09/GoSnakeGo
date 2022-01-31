package main

import (
	"fmt"
	"github.com/je09/GoSnakeGo/snake"
	"sync"
	"syscall/js"
)

const (
	size = 10
)

var (
	ctx    js.Value
	height int
	width  int
)

func main() {
	bootstrapApp()
}

func bootstrapApp() {
	// Getting window objects.
	window := js.Global()
	document := js.Global().Get("document")
	snakeGame := document.Call("getElementById", "snake-game")
	ctx = snakeGame.Call("getContext", "2d")

	// Setting game field size based on canvas area.
	height = snakeGame.Get("height").Int() / size
	width = snakeGame.Get("width").Int() / size
	f := snake.NewField(width, height, size)

	kCh := make(chan snake.Key)
	pos := make(chan snake.Positions)
	stop := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)
	go snake.Loop(f, stop, kCh, pos, &wg)
	go animate(pos, stop, &wg)

	// Default position for snake to crawl to.
	kCh <- "ArrowRight"

	allowedKeys := []snake.Key{
		"ArrowUp",
		"ArrowDown",
		"ArrowRight",
		"ArrowLeft",
	}

	keyboardListener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		k := snake.Key(args[0].Get("code").String())
		// Test if this an arrow key.
		for _, a := range allowedKeys {
			if a == k {
				kCh <- snake.Key(args[0].Get("code").String())
			}
		}
		return nil
	})
	window.Call("addEventListener", "keydown", keyboardListener)

	wg.Wait()
}

// Renders game frames when necessary.
func animate(pos chan snake.Positions, stop chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case pos := <-pos:
			// Clear game field.
			ctx.Call("clearRect", 0, 0, width*10, height*10)
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
		fmt.Println(s.Length(), s.Point.X, s.Point.Y)
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
