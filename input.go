package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var oldmouse pixel.Vec

func input() {

	switch {
	case win.Pressed(pixelgl.MouseButton1):

		movehex(win.MousePosition())
	}

	switch {
	case win.JustPressed(pixelgl.KeyLeft):
		HexViewTextBox.IMD().SetMatrix(pixel.IM.Moved(pixel.V(15, 15)))
		HexViewTextBox.Buffer()
	case win.JustPressed(pixelgl.KeyRight):

	}
	switch {
	case win.JustPressed(pixelgl.KeyUp):

	case win.JustPressed(pixelgl.KeyDown):

	}

	switch {
	case win.JustPressed(pixelgl.KeyEscape):
		win.SetClosed(true)
	}
}
