package gfx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"yearprogress/internal/core"
)

// exitEbitenKeys mirrors core.ExitKeys for the GUI backend.
var exitEbitenKeys = []ebiten.Key{
	ebiten.KeyEscape,
	ebiten.KeyEnter,
	ebiten.KeyNumpadEnter,
	ebiten.KeySpace,
	ebiten.KeyQ,
	ebiten.KeyBackspace,
}

func exitKeyPressed() bool {
	for _, r := range ebiten.AppendInputChars(nil) {
		if core.IsExitKey(string(r)) {
			return true
		}
	}
	for _, k := range exitEbitenKeys {
		if inpututil.IsKeyJustPressed(k) {
			return true
		}
	}
	return false
}