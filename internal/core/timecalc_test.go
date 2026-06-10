package core

import (
	"testing"
	"time"
)

func TestComputeTimeState_evenDay(t *testing.T) {
	loc := time.Local

	tests := []struct {
		name    string
		at      time.Time
		evenDay bool
	}{
		{"even day", time.Date(2026, 6, 11, 12, 0, 0, 0, loc), true},
		{"odd day", time.Date(2026, 6, 10, 12, 0, 0, 0, loc), false},
		{"day 1 odd", time.Date(2026, 1, 1, 12, 0, 0, 0, loc), false},
		{"day 2 even", time.Date(2026, 1, 2, 12, 0, 0, 0, loc), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := ComputeTimeState(tt.at)
			if state.EvenDay != tt.evenDay {
				t.Fatalf("day %d: got evenDay=%v, want %v", state.Day, state.EvenDay, tt.evenDay)
			}
		})
	}
}

func TestComputeTimeState_dayFraction(t *testing.T) {
	loc := time.Local
	at := time.Date(2026, 6, 10, 12, 0, 0, 0, loc)
	state := ComputeTimeState(at)

	if state.Day != 161 {
		t.Fatalf("got day %d, want 161", state.Day)
	}
	if state.Fraction < 0.49 || state.Fraction > 0.51 {
		t.Fatalf("got fraction %f, want ~0.5", state.Fraction)
	}
}

func TestComputeTimeState_yearPct(t *testing.T) {
	loc := time.Local
	at := time.Date(2026, 7, 1, 0, 0, 0, 0, loc)
	state := ComputeTimeState(at)

	if state.YearPct <= 0 || state.YearPct >= 1 {
		t.Fatalf("got yearPct %f, want between 0 and 1", state.YearPct)
	}
}

func TestFormatDayNumber(t *testing.T) {
	got := FormatDayNumber(161, 0.5)
	want := "161.50000"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}