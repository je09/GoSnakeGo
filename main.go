package main

import (
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
	// We subtract here to prevent player go outside the render borders.
	f := snake.NewField(width-size, height-size, size)

	kCh := make(chan snake.Key)
	pos := make(chan snake.Positions)
	stop := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)
	go snake.Loop(f, stop, kCh, pos, &wg)
	go render(pos, stop, &wg)

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
