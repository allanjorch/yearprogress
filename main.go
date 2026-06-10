package main

import (
	"flag"

	"yearprogress/internal/core"
	"yearprogress/internal/render/gfx"
	"yearprogress/internal/render/tui"
)

func main() {
	tuiMode := flag.Bool("tui", false, "open terminal UI instead of graphical window")
	_ = flag.Bool("gui", false, "open graphical window (default)")
	flag.Parse()

	cfg := core.DefaultConfig()
	if *tuiMode {
		tui.RunOrExit(cfg)
	} else {
		gfx.RunOrExit(cfg)
	}
}