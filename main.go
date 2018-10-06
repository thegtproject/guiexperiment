package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var win *pixelgl.Window

func setup() {
	cfg := pixelgl.WindowConfig{
		Title:  "GUI Experiment",
		Bounds: pixel.R(0, 0, 600, 450),
		VSync:  true,
	}
	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win = window
	initgui()
	run()
}

func main() {
	pixelgl.Run(setup)
}
