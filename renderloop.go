package main

import (
	"image/color"
)

var DrawManager = DrawList{}

func run() {

	for !win.Closed() {
		win.Clear(color.NRGBA{0, 0, 0, 255})
		input()
		DrawManager.Update()
		DrawManager.Draw(win)
		win.Update()
	}
}
