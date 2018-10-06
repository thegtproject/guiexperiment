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
	var ce []Element

	for _, child := range p.children {
		child.Attach(p)
		ce = append(ce, child)
		defer child.Init()
	}
}

func (p Panel) Buffer() {

	p.imd.Color = p.Properties.Color
	p.imd.Push(
		pixel.V(
			p.Properties.X,
			p.Properties.Y,
		),
		pixel.V(
			p.Properties.X+p.Properties.Width,
			p.Properties.Y+p.Properties.Height,
		),
	)
	p.imd.Rectangle(0)
	for _, child := range p.children {
		child.Buffer()
	}
}

func NewPanel(props *Properties, children ...Element) Element {
	p := &Panel{
		BasicElement: &BasicElement{
			Properties: props,
		},
	}
	for _, child := range children {
		p.children = append(p.children, child)

		child.Attach(p)
	}
	return p
}
