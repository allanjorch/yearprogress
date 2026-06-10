package tui

import (
	"strings"

	"yearprogress/internal/glyph"
)

func renderBigNumber(s string) string {
	return ansiBright(strings.Join(glyph.Lines(s), "\n"))
}