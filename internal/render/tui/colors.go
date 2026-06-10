package tui

import "strings"

const (
	ansiReset      = "\x1b[0m"
	ansiBold       = "\x1b[1m"
	ansiBgScreen   = "\x1b[48;2;20;19;16m"
	ansiFgBright   = "\x1b[38;2;232;220;168m"
	ansiFgMid      = "\x1b[38;2;184;168;130m"
	ansiFgDim      = "\x1b[38;2;92;86;72m"
	ansiBgBright   = "\x1b[48;2;232;220;168m"
	ansiFgOnBright = "\x1b[38;2;20;19;16m"
)

func ansiBright(s string) string {
	return ansiBold + ansiFgBright + s
}

func ansiMid(s string) string {
	return ansiFgMid + s
}

func ansiLabelOnFill(s string) string {
	return ansiBold + ansiBgBright + ansiFgOnBright + s
}

func ansiLabelOnTrack(s string) string {
	return ansiBold + ansiBgScreen + ansiFgBright + s
}

func bgPad(n int) string {
	if n <= 0 {
		return ""
	}
	return ansiBgScreen + strings.Repeat(" ", n)
}