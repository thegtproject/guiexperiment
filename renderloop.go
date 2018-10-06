package main

import (
	"image/color"

	"github.com/faiface/pixel"
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

func movehex(to pixel.Vec) {

	HexViewTextBox.Props().X = to.X - HexViewTextBox.Props().Width/2
	HexViewTextBox.Props().Y = to.Y - HexViewTextBox.Props().Height + 10
	HexViewTextBox.Buffer()
}
