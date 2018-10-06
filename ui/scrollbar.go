package ui

import (
	"fmt"

	"github.com/faiface/pixel"
)

type VerticleScrollbar struct {
	*BasicElement
}

// tests that Panel implements Element
var _ Element = Panel{}

func (vs VerticleScrollbar) Draw(t pixel.Target) {
	vs.imd.Draw(t)
	fmt.Println("asd")
}

func (vs VerticleScrollbar) Update() bool {

	return true
}

func (vs *VerticleScrollbar) Init() {
	var ce []Element

	for _, child := range vs.children {
		child.Attach(vs)
		ce = append(ce, child)
		defer child.Init()
	}
}
func (vs *VerticleScrollbar) Buffer() {

	vs.imd.Color = vs.Properties.Color
	vs.imd.Push(
		pixel.V(
			vs.Properties.X+vs.parent.Props().X,
			vs.Properties.Y+vs.parent.Props().Y,
		),
		pixel.V(
			vs.Properties.X+vs.Properties.Width+vs.parent.Props().X,
			vs.Properties.Y+vs.Properties.Height+vs.parent.Props().Y,
		),
	)
	vs.imd.Rectangle(0)
}

func NewVerticleScrollbar(props *Properties) Element {

	return &VerticleScrollbar{
		BasicElement: &BasicElement{
			Properties: props,
		},
	}
}
