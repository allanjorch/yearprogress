package gfx

import (
	"testing"
	"time"
)

func TestBarAnim_pct(t *testing.T) {
	start := time.Date(2026, 6, 10, 12, 0, 0, 0, time.UTC)
	target := 0.5
	anim := newBarAnim(target, start)

	if got := anim.pct(start, target); got != 0 {
		t.Fatalf("at start got %v, want 0", got)
	}

	half := start.Add(500 * time.Millisecond)
	if got := anim.pct(half, target); got != 0.25 {
		t.Fatalf("at 0.5s got %v, want 0.25", got)
	}

	end := start.Add(1 * time.Second)
	live := 0.51
	if got := anim.pct(end, live); got != live {
		t.Fatalf("at end got %v, want live %v", got, live)
	}
	if !anim.done {
		t.Fatal("animation should be done")
	}
}

func TestBarAnim_zeroTarget(t *testing.T) {
	start := time.Now()
	anim := newBarAnim(0, start)
	if !anim.done {
		t.Fatal("zero target should skip animation")
	}
	if got := anim.pct(start, 0.001); got != 0.001 {
		t.Fatalf("got %v, want live value", got)
	}
}