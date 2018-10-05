package ui

import (
	"fmt"
	"image/color"
)

type Properties struct {
	Width, Height, X, Y float64
	Color               color.NRGBA
}

func (p *Properties) Clone() *Properties {
	clone := *p
	return &clone
}

func (p *Properties) Modify(changes ...interface{}) *Properties {
	if len(changes)%2 != 0 {
		fmt.Println("argument count should be divisible by 2 (string")
		return nil
	}
	for i := 0; i < len(changes); i += 2 {
		name := changes[i+0].(string)
		val := changes[i+1]

		if changeFn, ok := propertyMap[name]; ok {
			changeFn(p, val)
		}
	}
	return p
}

var propertyMap = map[string]func(p *Properties, val interface{}){
	"X": func(p *Properties, val interface{}) {
		p.X = val.(float64)
	},
	"Y": func(p *Properties, val interface{}) {
		p.Y = val.(float64)
	},
	"Width": func(p *Properties, val interface{}) {
		p.Width = val.(float64)
	},
	"Height": func(p *Properties, val interface{}) {
		p.Height = val.(float64)
	},
	"Color": func(p *Properties, val interface{}) {
		p.Color = val.(color.NRGBA)
	},
}
