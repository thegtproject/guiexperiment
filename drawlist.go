package main

import "github.com/faiface/pixel"

type Drawable interface {
	NeedsRedraw() bool
	Update() bool
	Draw(t pixel.Target)
}

type DrawList struct {
	Items []Drawable

	drawqueue []Drawable
}

func (dl *DrawList) Draw(t pixel.Target) {
	for _, obj := range dl.Items {
		obj.Draw(t)
	}
	dl.drawqueue = nil
}

func (dl *DrawList) Update() {
	for _, obj := range dl.Items {
		obj.Update()
		if obj.NeedsRedraw() {
			dl.PushDrawQueue(obj)
		}
	}
}

func (dl *DrawList) Add(obj Drawable) {
	dl.Items = append(dl.Items, obj)
}

func (dl *DrawList) PushDrawQueue(obj Drawable) {
	dl.drawqueue = append(dl.drawqueue, obj)
}
