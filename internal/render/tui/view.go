package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"yearprogress/internal/core"
	"yearprogress/internal/glyph"
)

func Render(scene core.Scene, width, height int) string {
	contentWidth := glyph.DisplayWidth(scene.DayLine)
	bar := renderProgressBar(scene)

	var lines []string
	lines = append(lines, strings.Split(renderBigNumber(scene.DayLine), "\n")...)
	lines = append(lines, "")

	if scene.ShowEvenDay {
		if scene.EvenDay {
			lines = append(lines, ansiBright(scene.EvenDayLabel()))
		} else {
			lines = append(lines, ansiMid(scene.EvenDayLabel()))
		}
		lines = append(lines, "")
	}

	lines = append(lines, strings.Split(bar, "\n")...)

	if width > 0 && height > 0 {
		return paintScreen(width, height, lines, contentWidth)
	}

	return strings.Join(lines, "\n")
}

func paintScreen(termW, termH int, contentLines []string, contentWidth int) string {
	var blockLines []string
	for _, line := range contentLines {
		if line == "" {
			blockLines = append(blockLines, bgPad(contentWidth))
		} else {
			blockLines = append(blockLines, centerInWidth(contentWidth, line))
		}
	}

	contentH := len(blockLines)
	topPad := max(0, (termH-contentH)/2)

	rows := make([]string, termH)
	for i := range rows {
		switch {
		case i < topPad, i >= topPad+contentH:
			rows[i] = bgPad(termW)
		default:
			rows[i] = centerOnScreen(termW, blockLines[i-topPad])
		}
	}

	return ansiBgScreen + strings.Join(rows, "\n") + ansiReset
}

func centerInWidth(w int, line string) string {
	lineW := lipgloss.Width(line)
	left := (w - lineW) / 2
	right := w - lineW - left
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	return bgPad(left) + line + bgPad(right)
}

func centerOnScreen(termW int, line string) string {
	lineW := lipgloss.Width(line)
	left := (termW - lineW) / 2
	right := termW - lineW - left
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	return bgPad(left) + line + bgPad(right)
}