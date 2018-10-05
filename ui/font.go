package ui

import (
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var FontManager = &fontManager{Fonts: make(map[string]Font)}

type Font struct {
	Text  *text.Text
	Atlas *text.Atlas
	*Properties
}

func (f *Font) GetChildren() []Element { return nil }
func (f *Font) Update() bool           { return true }
func (f *Font) Draw(t pixel.Target)    {}
func (f *Font) Init()                  {}

type fontManager struct {
	Fonts map[string]Font
}

// LoadFont ...
func (fm *fontManager) LoadFont(filename string, size float64, runesets ...[]rune) *Font {
	if len(runesets) == 0 {
		runesets = [][]rune{text.ASCII}
	}
	face, err := loadtruetype(filename, size)
	if err != nil {
		panic(err)
	}
	atlas := makeatlas(face, runesets...)
	txt := maketext(atlas)
	return &Font{Text: txt, Atlas: atlas}
}

func loadtruetype(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}
	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func makeatlas(face font.Face, runesets ...[]rune) *text.Atlas {
	return text.NewAtlas(face, runesets...)
}

func maketext(atlas *text.Atlas) *text.Text {
	return text.New(
		pixel.ZV,
		atlas,
	)
}
