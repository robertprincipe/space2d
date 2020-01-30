package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/robertprincipe/spacepixel/internal"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "My first game",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		log.Fatal(err)
	}

	world := internal.NewWorld(windowWidth, windowHeight)

	if err := world.AddBackground("resources/purple.png"); err != nil {
		log.Fatal(err)
	}

	world.Draw(win)

	for !win.Closed() {
		win.Update()
	}

	// win.Clear(colornames.Aqua)

	// for !win.Closed() {
	// 	win.Update()
	// }
}
