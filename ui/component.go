package ui

import "github.com/faiface/pixel/imdraw"

type Component struct {
	Element
	Name string

	NeedRedraw bool
}

func NewComponent(name string, rootelement Element) *Component {
	rootelement.SetIMD(imdraw.New(nil))
	return &Component{
		Name:       name,
		NeedRedraw: true,
		Element:    rootelement,
	}
}

func (c *Component) NeedsRedraw() bool {
	return c.NeedRedraw
}

func (c *Component) Init() {
	c.Element.Init()

}
