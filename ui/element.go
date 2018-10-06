package ui

import (
	"github.com/faiface/pixel/imdraw"
)

type Attachable interface {
	Attach(parent Element)
}

type Element interface {
	Attachable
	GetChildren() []Element
	Update() bool
	Buffer()
	Init()
	SetImd(imd *imdraw.IMDraw)
	Props() *Properties
	Parent() Element

	IMD() *imdraw.IMDraw
}

type BasicElement struct {
	Element
	*Properties
	parent   Element
	imd      *imdraw.IMDraw
	children []Element
}

func (e *BasicElement) SetImd(imd *imdraw.IMDraw) {
	e.imd = imd
}
func (e *BasicElement) Attach(parent Element) {
	e.parent = parent
	e.imd = parent.IMD()

}

func (e *BasicElement) Parent() Element {
	return e.parent
}

func (e *BasicElement) GetChildren() []Element {
	return e.children
}

func (e *BasicElement) IMD() *imdraw.IMDraw {
	return e.imd
}

func (e *BasicElement) Props() *Properties {
	return e.Properties
}
