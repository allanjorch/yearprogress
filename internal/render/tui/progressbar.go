package tui

import (
	"strings"

	"yearprogress/internal/core"
	"yearprogress/internal/glyph"
)

const (
	barCharFill  = "█"
	barCharTrack = "░"
)

func renderProgressBar(scene core.Scene) string {
	width := glyph.DisplayWidth(scene.DayLine)
	if width < 1 {
		return ""
	}

	filled := int(scene.YearPct * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}

	label := core.FormatBarLabel(scene.YearPct)
	labelStart := (width - len(label)) / 2
	if labelStart < 0 {
		labelStart = 0
	}

	lines := make([]string, glyph.BarRows)
	for row := range lines {
		withLabel := row == glyph.BarLabel
		lines[row] = renderBarRow(width, filled, label, labelStart, withLabel)
	}
	return strings.Join(lines, "\n")
}

type barCellKind int

const (
	barCellFill barCellKind = iota
	barCellTrack
	barCellLabelFill
	barCellLabelTrack
)

func renderBarRow(width, filled int, label string, labelStart int, withLabel bool) string {
	kinds := make([]barCellKind, width)
	for col := 0; col < width; col++ {
		onFill := col < filled
		if withLabel && col >= labelStart && col < labelStart+len(label) {
			if onFill {
				kinds[col] = barCellLabelFill
			} else {
				kinds[col] = barCellLabelTrack
			}
			continue
		}
		if onFill {
			kinds[col] = barCellFill
		} else {
			kinds[col] = barCellTrack
		}
	}

	var b strings.Builder
	b.Grow(width * 8)
	for col := 0; col < width; {
		kind := kinds[col]
		end := col + 1
		for end < width && kinds[end] == kind {
			end++
		}
		n := end - col
		switch kind {
		case barCellFill:
			b.WriteString(ansiFgBright + strings.Repeat(barCharFill, n))
		case barCellTrack:
			b.WriteString(ansiFgDim + strings.Repeat(barCharTrack, n))
		case barCellLabelFill:
			b.WriteString(ansiLabelOnFill(label[col-labelStart : col-labelStart+n]))
		case barCellLabelTrack:
			b.WriteString(ansiLabelOnTrack(label[col-labelStart : col-labelStart+n]))
		}
		col = end
	}
	return b.String()
}