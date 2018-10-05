package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func input() {
	switch {
	case win.JustPressed(pixelgl.KeyLeft):

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
