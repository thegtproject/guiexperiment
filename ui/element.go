package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Element interface {
	GetChildren() []Element
	Update() bool
	Draw(t pixel.Target)
	Init()
	SetIMD(imd *imdraw.IMDraw)
	Props() *Properties
}

type BasicElement struct {
	*Properties
	imd      *imdraw.IMDraw
	children []Element
}

func (e *BasicElement) GetChildren() []Element {
	return e.children
}

func (e *BasicElement) SetIMD(imd *imdraw.IMDraw) {
	e.imd = imd
}

func (e *BasicElement) Props() *Properties {
	return e.Properties
}
