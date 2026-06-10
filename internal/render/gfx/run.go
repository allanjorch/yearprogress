package gfx

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"yearprogress/internal/core"
)

// Run opens the graphical window.
func Run(cfg core.Config) error {
	if err := InitFonts(); err != nil {
		return fmt.Errorf("gfx fonts: %w", err)
	}

	ebiten.SetWindowTitle("yearprogress")
	ebiten.SetWindowSize(defaultW, defaultH)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(newApp(cfg)); err != nil {
		return fmt.Errorf("gfx: %w", err)
	}
	return nil
}

func RunOrExit(cfg core.Config) {
	if err := Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}