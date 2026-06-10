package gfx

import (
	"bytes"
	_ "embed"
	"fmt"
	"sync"

	text "github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed fonts/Nunito.ttf
var nunitoTTF []byte

var (
	fontOnce   sync.Once
	fontInit   error
	nunitoSrc  *text.GoTextFaceSource
	weightTag  text.Tag
	weightSemi = float32(600)
	weightBold = float32(700)
	weightReg  = float32(400)
)

func InitFonts() error {
	fontOnce.Do(func() {
		nunitoSrc, fontInit = text.NewGoTextFaceSource(bytes.NewReader(nunitoTTF))
		if fontInit != nil {
			return
		}
		weightTag, fontInit = text.ParseTag("wght")
	})
	return fontInit
}

func displayFace(size float64) (*text.GoTextFace, error) {
	if err := InitFonts(); err != nil {
		return nil, err
	}
	f := &text.GoTextFace{Source: nunitoSrc, Size: size}
	f.SetVariation(weightTag, weightSemi)
	return f, nil
}

func barLabelFace(size float64) (*text.GoTextFace, error) {
	if err := InitFonts(); err != nil {
		return nil, err
	}
	f := &text.GoTextFace{Source: nunitoSrc, Size: size}
	f.SetVariation(weightTag, weightBold)
	return f, nil
}

func bodyFace(size float64) (*text.GoTextFace, error) {
	if err := InitFonts(); err != nil {
		return nil, err
	}
	f := &text.GoTextFace{Source: nunitoSrc, Size: size}
	f.SetVariation(weightTag, weightReg)
	return f, nil
}

func measureDisplay(s string, size float64) (w, h float64, err error) {
	f, err := displayFace(size)
	if err != nil {
		return 0, 0, err
	}
	w, h = text.Measure(s, f, 0)
	return w, h, nil
}

func measureBody(s string, size float64) (w, h float64, err error) {
	f, err := bodyFace(size)
	if err != nil {
		return 0, 0, err
	}
	w, h = text.Measure(s, f, 0)
	return w, h, nil
}

func mustInitFonts() {
	if err := InitFonts(); err != nil {
		panic(fmt.Sprintf("gfx fonts: %v", err))
	}
}