package gfx

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text/v2"

	"yearprogress/internal/core"
)

const labelPillRadius = float32(5)

func Draw(screen *ebiten.Image, scene core.Scene, layout Layout, barPct float64) {
	screen.Fill(colorBg)
	drawDayNumber(screen, scene.DayLine, layout)
	if layout.HasEvenDay {
		drawEvenDayLabel(screen, scene, layout)
	}
	drawBar(screen, layout, barPct)
}

func drawDayNumber(screen *ebiten.Image, dayLine string, layout Layout) {
	f, err := displayFace(layout.DayFontSize)
	if err != nil {
		return
	}
	x := layout.OriginX
	y := layout.DayTop

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(colorBright)
	text.Draw(screen, dayLine, f, op)
}

func drawEvenDayLabel(screen *ebiten.Image, scene core.Scene, layout Layout) {
	label := scene.EvenDayLabel()
	size := layout.DayFontSize * 0.28
	if size < minLabelSize {
		size = minLabelSize
	}
	f, err := bodyFace(size)
	if err != nil {
		return
	}
	w, _ := text.Measure(label, f, 0)
	x := layout.OriginX + (layout.ContentW-w)/2
	y := layout.EvenDayTop

	fg := colorMid
	if scene.EvenDay {
		fg = colorBright
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(fg)
	text.Draw(screen, label, f, op)
}

func drawBar(screen *ebiten.Image, layout Layout, barPct float64) {
	x := float32(layout.BarX)
	y := float32(layout.BarTop)
	w := float32(layout.BarW)
	h := float32(layout.BarH)
	r := float32(layout.BarRadius)

	fillW := float32(layout.BarW * barPct)
	if fillW > w {
		fillW = w
	}

	fillRoundRect(screen, x, y, w, h, r, colorDim)
	if fillW > 0 {
		fillRoundRect(screen, x, y, fillW, h, r, colorBright)
	}

	drawBarLabel(screen, layout, barPct, float64(fillW))
}

type barLabelChar struct {
	ch     string
	x      float64
	topY   float64
	h      float64
	pillW  float64
	onFill bool
}

func drawBarLabel(screen *ebiten.Image, layout Layout, barPct, fillW float64) {
	label := core.FormatBarLabel(barPct)
	f, err := barLabelFace(layout.BarLabelSize)
	if err != nil {
		return
	}
	labelW, labelH := text.Measure(label, f, 0)
	labelX := layout.BarX + (layout.BarW-labelW)/2
	labelTop := layout.BarTop + (layout.BarH-labelH)/2

	var chars []barLabelChar
	charX := labelX
	for _, r := range label {
		ch := string(r)
		cw, chH := text.Measure(ch, f, 0)
		chars = append(chars, barLabelChar{
			ch:     ch,
			x:      charX,
			topY:   labelTop,
			h:      chH,
			pillW:  charPillWidth(ch, f, cw),
			onFill: charX+cw/2-layout.BarX < fillW,
		})
		charX += cw
	}

	for _, c := range chars {
		if !c.onFill {
			drawLabelCharBg(screen, c, colorDim)
		}
	}
	for _, c := range chars {
		fg := colorBright
		if c.onFill {
			fg = colorBg
		}
		drawLabelCharText(screen, c, f, fg)
	}
}

func charPillWidth(ch string, f *text.GoTextFace, advance float64) float64 {
	glyphs := text.AppendGlyphs(nil, ch, f, nil)
	maxRight := advance
	for _, g := range glyphs {
		if g.Image == nil {
			continue
		}
		b := g.Image.Bounds()
		right := g.X + float64(b.Max.X)
		if right > maxRight {
			maxRight = right
		}
	}
	return maxRight
}

func drawLabelCharBg(screen *ebiten.Image, c barLabelChar, bg color.Color) {
	pad := float32(3)
	fillRoundRect(
		screen,
		float32(c.x)-pad,
		float32(c.topY)-pad,
		float32(c.pillW)+pad*2,
		float32(c.h)+pad*2,
		labelPillRadius,
		bg,
	)
}

func drawLabelCharText(screen *ebiten.Image, c barLabelChar, f *text.GoTextFace, fg color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(c.x, c.topY)
	op.ColorScale.ScaleWithColor(fg)
	text.Draw(screen, c.ch, f, op)
}