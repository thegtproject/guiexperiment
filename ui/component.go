package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Component struct {
	Element
	Name string

	NeedRedraw bool
	imd        *imdraw.IMDraw
}

func NewComponent(name string, rootelement Element) *Component {

	cmp := &Component{
		Element:    rootelement,
		Name:       name,
		NeedRedraw: true,
	}
	cmp.imd = imdraw.New(nil)

	cmp.Element.SetImd(cmp.imd)

	return cmp
}

func (c *Component) NeedsRedraw() bool {
	return c.NeedRedraw
}

func (c *Component) Buffer() {
	c.imd.Reset()
	c.imd.Clear()
	c.Element.Buffer()
}

func (c *Component) Init() {
	c.Element.Init()
	c.Buffer()
}

func (c *Component) Draw(t pixel.Target) {
	c.imd.Draw(t)
}
