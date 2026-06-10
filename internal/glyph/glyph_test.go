package glyph

import "testing"

func TestDisplayWidth(t *testing.T) {
	if got := DisplayWidth("161.50000"); got != 53 {
		t.Fatalf("got %d, want 53", got)
	}
	narrow := DisplayWidth("1.00000")
	wide := DisplayWidth("365.99999")
	if wide <= narrow {
		t.Fatalf("expected wider display for more digits")
	}
}

func TestIsFilled(t *testing.T) {
	if !IsFilled('0', 0, 0) {
		t.Fatal("top-left of 0 should be filled")
	}
	if !IsFilled('0', 0, 2) {
		t.Fatal("top edge of 0 should be filled at col 2")
	}
	if IsFilled('0', 1, 2) {
		t.Fatal("interior of 0 should be empty at row 1 col 2")
	}
}