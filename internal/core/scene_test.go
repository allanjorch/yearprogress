package core

import "testing"

func TestFormatBarLabel(t *testing.T) {
	tests := []struct {
		pct  float64
		want string
	}{
		{0, "0.00%"},
		{0.44123, "44.12%"},
		{0.5, "50.00%"},
		{1, "100.00%"},
	}
	for _, tt := range tests {
		if got := FormatBarLabel(tt.pct); got != tt.want {
			t.Fatalf("FormatBarLabel(%v) = %q, want %q", tt.pct, got, tt.want)
		}
	}
}