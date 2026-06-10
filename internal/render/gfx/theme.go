package gfx

import "image/color"

// Saturated chalk palette — the subtle TUI tones were imperceptible at small pixel sizes.
var (
	colorBg     = color.RGBA{0x14, 0x13, 0x10, 0xff}
	colorDim    = color.RGBA{0x72, 0x62, 0x42, 0xff}
	colorMid    = color.RGBA{0xcc, 0xb4, 0x68, 0xff}
	colorBright = color.RGBA{0xff, 0xec, 0x88, 0xff}
)