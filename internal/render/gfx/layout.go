package gfx

import (
	"math"

	"yearprogress/internal/core"
)

const (
	marginFrac   = 0.08
	minDaySize   = 24.0
	maxDaySize   = 200.0
	minBarH      = 24.0
	minLabelSize = 14.0
)

type Layout struct {
	DayFontSize  float64
	DayTextW     float64
	DayTextH     float64
	BarW         float64
	BarH         float64
	BarRadius    float64
	BarLabelSize float64
	ContentW     float64
	ContentH     float64
	OriginX      float64
	DayTop       float64
	EvenDayTop   float64
	BarTop       float64
	BarX         float64
	HasEvenDay   bool
}

func ComputeLayout(scene core.Scene, screenW, screenH float64) Layout {
	mustInitFonts()

	marginX := screenW * marginFrac
	marginY := screenH * marginFrac
	availW := screenW - marginX*2
	availH := screenH - marginY*2

	dayBudget := availH * 0.58
	if scene.ShowEvenDay {
		dayBudget = availH * 0.5
	}

	daySize := fitDayFontSize(scene.DayLine, availW, dayBudget)
	dayW, dayH := mustMeasureDisplay(scene.DayLine, daySize)

	barH := math.Max(minBarH, daySize*0.34)
	barRadius := barH / 2
	labelSize := math.Max(minLabelSize, barH*0.44)

	gap := daySize * 0.18
	evenDayH := daySize * 0.28

	contentW := dayW
	contentH := dayH + gap + barH
	if scene.ShowEvenDay {
		contentH += evenDayH + gap
	}

	originX := (screenW - contentW) / 2
	dayTop := (screenH - contentH) / 2
	barX := originX

	l := Layout{
		DayFontSize:  daySize,
		DayTextW:     dayW,
		DayTextH:     dayH,
		BarW:         dayW,
		BarH:         barH,
		BarRadius:    barRadius,
		BarLabelSize: labelSize,
		ContentW:     contentW,
		ContentH:     contentH,
		OriginX:      originX,
		DayTop:       dayTop,
		BarX:         barX,
		HasEvenDay:   scene.ShowEvenDay,
	}

	if scene.ShowEvenDay {
		l.EvenDayTop = dayTop + dayH + gap
		l.BarTop = l.EvenDayTop + evenDayH + gap
	} else {
		l.BarTop = dayTop + dayH + gap
	}

	return l
}

func fitDayFontSize(dayLine string, maxW, maxH float64) float64 {
	lo, hi := minDaySize, maxDaySize
	for hi-lo > 0.5 {
		mid := (lo + hi) / 2
		w, h, err := measureDisplay(dayLine, mid)
		if err != nil {
			return minDaySize
		}
		if w <= maxW && h <= maxH {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

func mustMeasureDisplay(s string, size float64) (float64, float64) {
	w, h, err := measureDisplay(s, size)
	if err != nil {
		panic(err)
	}
	return w, h
}