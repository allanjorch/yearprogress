package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"

	"yearprogress/internal/core"
	"yearprogress/internal/glyph"
)

func TestRenderProgressBar_noPartialCells(t *testing.T) {
	scene := core.Scene{DayLine: "1234567890", YearPct: 0.456, Year: 2026}
	width := glyph.DisplayWidth(scene.DayLine)
	wantFilled := int(0.456 * float64(width))

	bar := renderProgressBar(scene)
	plain := stripANSI(bar)

	rows := strings.Split(plain, "\n")
	if len(rows) != glyph.BarRows {
		t.Fatalf("got %d rows, want %d", len(rows), glyph.BarRows)
	}

	for i, row := range rows {
		if i == glyph.BarLabel {
			if !strings.Contains(row, "%") {
				t.Fatalf("label row should contain percentage")
			}
			continue
		}
		filled := strings.Count(row, barCharFill)
		empty := strings.Count(row, barCharTrack)
		if filled+empty != width {
			t.Fatalf("row %d: got %d filled + %d empty = %d, want total %d", i, filled, empty, filled+empty, width)
		}
		if filled != wantFilled {
			t.Fatalf("row %d: got %d filled cells, want %d (no partial)", i, filled, wantFilled)
		}
	}
}

func TestRenderProgressBar_matchesCounterWidth(t *testing.T) {
	scene := core.Scene{DayLine: "161.50000", YearPct: 0.5, Year: 2026}
	want := glyph.DisplayWidth(scene.DayLine)
	bar := stripANSI(renderProgressBar(scene))
	firstRow := strings.Split(bar, "\n")[0]
	if lipgloss.Width(firstRow) != want {
		t.Fatalf("bar width %d, want %d", lipgloss.Width(firstRow), want)
	}
}

func TestRenderProgressBar_fewResets(t *testing.T) {
	scene := core.Scene{DayLine: strings.Repeat("0", 11), YearPct: 0.441, Year: 2026}
	bar := renderProgressBar(scene)
	if strings.Count(bar, ansiReset) > 2 {
		t.Fatalf("expected few ANSI resets in bar")
	}
}