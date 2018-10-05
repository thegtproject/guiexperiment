package main

import (
	"image/color"

	"github.com/thegtproject/guiexperiment/ui"
)

var HexViewTextBox *ui.Component

func initHexDisplay() {

	props := &ui.Properties{
		Width:  300,
		Height: 300,
		X:      0,
		Y:      0,
		Color: color.NRGBA{
			R: 25,
			G: 25,
			B: 25,
			A: 255,
		},
	}

	HexViewTextBox = ui.NewComponent(
		"hexviewtextbox",
		ui.NewPanel(
			props,
			ui.NewVerticleScrollbar(
				props.Clone().Modify(
					"X", float64(290),
					"Width", float64(10),
					"Color", color.NRGBA{55, 35, 55, 255},
				)),
		),
	)

	HexViewTextBox.Init()
	DrawManager.Add(HexViewTextBox)
}
