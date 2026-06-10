package glyph

import "strings"

const (
	Height   = 7
	Width    = 5
	Gap      = 1
	BarRows  = 3
	BarLabel = 1
)

var patterns = map[rune][Height]string{
	'0': {"█████", "█   █", "█   █", "█   █", "█   █", "█   █", "█████"},
	'1': {"  █  ", " ██  ", "  █  ", "  █  ", "  █  ", "  █  ", " ███ "},
	'2': {"█████", "    █", "    █", "█████", "█    ", "█    ", "█████"},
	'3': {"█████", "    █", "    █", "█████", "    █", "    █", "█████"},
	'4': {"█   █", "█   █", "█   █", "█████", "    █", "    █", "    █"},
	'5': {"█████", "█    ", "█    ", "█████", "    █", "    █", "█████"},
	'6': {"█████", "█    ", "█    ", "█████", "█   █", "█   █", "█████"},
	'7': {"█████", "    █", "    █", "   █ ", "  █  ", "  █  ", "  █  "},
	'8': {"█████", "█   █", "█   █", "█████", "█   █", "█   █", "█████"},
	'9': {"█████", "█   █", "█   █", "█████", "    █", "    █", "█████"},
	'.': {"     ", "     ", "     ", "     ", "     ", "  █  ", "  █  "},
}

// DisplayWidth returns the character-column width of a glyph string in the TUI grid.
func DisplayWidth(s string) int {
	if len(s) == 0 {
		return 0
	}
	return len([]rune(s))*(Width+Gap) - Gap
}

func IsFilled(ch rune, row, col int) bool {
	if row < 0 || row >= Height || col < 0 || col >= Width {
		return false
	}
	pattern, ok := patterns[ch]
	if !ok {
		return false
	}
	runes := []rune(pattern[row])
	if col >= len(runes) {
		return false
	}
	return runes[col] == '█'
}

// Lines returns TUI text rows for a glyph string.
func Lines(s string) []string {
	lines := make([]string, Height)
	for _, ch := range s {
		pattern, ok := patterns[ch]
		if !ok {
			continue
		}
		for i, row := range pattern {
			lines[i] += row + strings.Repeat(" ", Gap)
		}
	}
	return lines
}