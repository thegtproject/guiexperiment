package ui

import (
	"github.com/faiface/pixel"
)

type Panel struct {
	*BasicElement
}

// tests that Panel implements Element
var _ Element = Panel{}

func (p Panel) Draw(t pixel.Target) {
	p.imd.Draw(t)
}

func (p Panel) Update() bool {

	return true
}

func (p Panel) Init() {
	p.imd.Color = p.Properties.Color
	p.imd.Push(
		pixel.V(
			p.Properties.X,
			p.Properties.Y,
		),
		pixel.V(
			p.Properties.Width,
			p.Properties.Height,
		),
	)
	p.imd.Rectangle(0)

	var ce []Element
	for _, child := range p.children {
		child.SetIMD(p.imd)
		ce = append(ce, child)
		defer child.Init()
	}
}

func NewPanel(props *Properties, children ...Element) Element {
	return &Panel{
		BasicElement: &BasicElement{
			Properties: props,
			children:   children,
		},
	}
}
