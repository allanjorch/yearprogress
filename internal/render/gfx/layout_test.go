package gfx

import (
	"testing"

	"yearprogress/internal/core"
)

func TestMain(m *testing.M) {
	if err := InitFonts(); err != nil {
		panic(err)
	}
	m.Run()
}

func TestComputeLayout_scalesWithWindow(t *testing.T) {
	scene := core.Scene{DayLine: "161.50000", YearPct: 0.44, Year: 2026}
	small := ComputeLayout(scene, 320, 240)
	large := ComputeLayout(scene, 1280, 800)

	if large.DayFontSize <= small.DayFontSize {
		t.Fatalf("expected larger font for bigger window: small=%f large=%f", small.DayFontSize, large.DayFontSize)
	}
	if large.ContentW <= 0 || large.ContentH <= 0 {
		t.Fatal("expected positive content dimensions")
	}
}

func TestComputeLayout_evenDayAddsHeight(t *testing.T) {
	base := core.Scene{DayLine: "161.50000", YearPct: 0.44, Year: 2026}
	withLabel := base
	withLabel.ShowEvenDay = true

	l0 := ComputeLayout(base, 800, 600)
	l1 := ComputeLayout(withLabel, 800, 600)

	if !l1.HasEvenDay {
		t.Fatal("expected HasEvenDay")
	}
	if l1.ContentH <= l0.ContentH {
		t.Fatal("even day label should increase content height")
	}
}

func TestComputeLayout_defaultWindowReadable(t *testing.T) {
	scene := core.Scene{DayLine: "161.50000", YearPct: 0.44, Year: 2026}
	l := ComputeLayout(scene, defaultW, defaultH)

	if l.DayFontSize < 32 {
		t.Fatalf("day font size %f too small at default window", l.DayFontSize)
	}
	if l.BarLabelSize < minLabelSize {
		t.Fatalf("label size %f below minimum %f", l.BarLabelSize, minLabelSize)
	}
	if l.BarRadius != l.BarH/2 {
		t.Fatalf("bar radius %f want %f", l.BarRadius, l.BarH/2)
	}
	if l.BarW != l.DayTextW {
		t.Fatalf("bar width should match day text width")
	}
}