package gfx

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"yearprogress/internal/core"
)

const (
	// Wide default — the day string spans ~59 cell columns; small windows shrink cells
	// to a few pixels and colors become imperceptible.
	defaultW = 1280
	defaultH = 720
	tick     = 50 * time.Millisecond
)

type App struct {
	config  core.Config
	scene   core.Scene
	last    time.Time
	barAnim barAnim
}

func newApp(cfg core.Config) *App {
	now := time.Now()
	scene := core.BuildScene(core.ComputeTimeState(now), cfg)
	return &App{
		config:  cfg,
		scene:   scene,
		last:    now,
		barAnim: newBarAnim(scene.YearPct, now),
	}
}

func (a *App) Update() error {
	if exitKeyPressed() {
		return ebiten.Termination
	}

	now := time.Now()
	if now.Sub(a.last) >= tick {
		a.scene = core.BuildScene(core.ComputeTimeState(now), a.config)
		a.last = now
	}
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	bounds := screen.Bounds()
	layout := ComputeLayout(a.scene, float64(bounds.Dx()), float64(bounds.Dy()))
	barPct := a.barAnim.pct(time.Now(), a.scene.YearPct)
	Draw(screen, a.scene, layout, barPct)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	if outsideWidth <= 0 {
		outsideWidth = defaultW
	}
	if outsideHeight <= 0 {
		outsideHeight = defaultH
	}
	return outsideWidth, outsideHeight
}

