package ui

import (
	"github.com/faiface/pixel"
)

type VerticleScrollbar struct {
	*BasicElement
}

// tests that Panel implements Element
var _ Element = Panel{}

func (vs VerticleScrollbar) Draw(t pixel.Target) {
	vs.imd.Draw(t)
}

func (vs VerticleScrollbar) Update() bool {

	return true
}

func (vs VerticleScrollbar) Init() {
	vs.imd.Color = vs.Properties.Color
	vs.imd.Push(
		pixel.V(
			vs.Properties.X,
			vs.Properties.Y,
		),
		pixel.V(
			vs.Properties.Width,
			vs.Properties.Height,
		),
	)
	vs.imd.Rectangle(0)

	var ce []Element
	for _, child := range vs.children {
		child.SetIMD(vs.imd)
		ce = append(ce, child)
		defer child.Init()
	}
}

func NewVerticleScrollbar(props *Properties) Element {
	return &VerticleScrollbar{
		BasicElement: &BasicElement{
			Properties: props,
		},
	}
}
