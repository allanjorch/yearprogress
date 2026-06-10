package tui

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"yearprogress/internal/core"
	"yearprogress/internal/glyph"
)

func TestRender_containsExpectedElements(t *testing.T) {
	state := core.ComputeTimeState(time.Date(2026, 6, 11, 12, 0, 0, 0, time.Local))
	scene := core.BuildScene(state, core.DefaultConfig())
	out := Render(scene, 80, 24)

	plain := stripANSI(out)
	blockLines := 0
	for _, line := range strings.Split(plain, "\n") {
		if strings.Contains(line, "█") && strings.Contains(line, " ") {
			blockLines++
		}
	}
	if blockLines < glyph.Height {
		t.Fatalf("expected block digit lines in output, got %d", blockLines)
	}
	if strings.Contains(plain, "EVEN DAY") || strings.Contains(plain, "ODD DAY") {
		t.Fatalf("even day label should be hidden by default")
	}
	if !regexp.MustCompile(`\d+\.\d{2}%`).MatchString(plain) {
		t.Fatalf("expected percentage label inside bar")
	}
	if strings.Count(plain, barCharFill) == 0 {
		t.Fatalf("expected filled bar characters")
	}
	if strings.Count(plain, barCharTrack) == 0 {
		t.Fatalf("expected track bar characters")
	}
}

func TestRender_showEvenDayWhenEnabled(t *testing.T) {
	state := core.ComputeTimeState(time.Date(2026, 6, 11, 12, 0, 0, 0, time.Local))
	cfg := core.Config{ShowEvenDayLabel: true}
	out := stripANSI(Render(core.BuildScene(state, cfg), 80, 24))
	if !strings.Contains(out, "EVEN DAY") {
		t.Fatalf("expected even day label when ShowEvenDayLabel is true")
	}
}

func TestRender_fewResets(t *testing.T) {
	state := core.ComputeTimeState(time.Date(2026, 6, 11, 12, 0, 0, 0, time.Local))
	out := Render(core.BuildScene(state, core.DefaultConfig()), 80, 24)
	resets := strings.Count(out, ansiReset)
	if resets > 5 {
		t.Fatalf("expected few ANSI resets in full frame, got %d", resets)
	}
}

var ansiRe = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

func stripANSI(s string) string {
	return ansiRe.ReplaceAllString(s, "")
}