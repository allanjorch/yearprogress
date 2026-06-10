# yearprogress

A live day-of-year and year-progress display with a chalk-yellow dark theme. Shows the current day as a large counter with sub-day decimals, plus a pill-shaped progress bar with an animated fill on launch.

**Repository:** [github.com/allanjorch/yearprogress](https://github.com/allanjorch/yearprogress)

## Features

- Large day-of-year counter with five fractional digits (e.g. `161.50000`)
- Year progress bar with percentage label (`44.12%`)
- GUI launch animation: bar fills left-to-right to the current year percentage
- Graphical window (default) or terminal UI (`--tui`)
- Optional even/odd day label (off by default)

## Requirements

- Go 1.25 or later
- **GUI:** a graphics stack supported by [Ebitengine](https://ebitengine.org/) (Windows, macOS, Linux)
- **TUI:** a terminal with ANSI color support (e.g. Windows Terminal, Ghostty, iTerm)

## Run

```bash
git clone https://github.com/allanjorch/yearprogress.git
cd yearprogress
go run .          # GUI (default)
go run . --tui    # terminal UI
```

Press **Esc**, **Enter**, **Space**, **Q**, or **Backspace** to exit.

## Build

**Linux / macOS:**

```bash
go build -o yearprogress .
./yearprogress
```

**Windows (from Linux):**

```bash
GOOS=windows GOARCH=amd64 go build -o yearprogress.exe .
```

**Windows (native):** install Go and a C compiler (Ebitengine uses CGO), then `go build -o yearprogress.exe .`

## Flags

| Flag | Description |
|------|-------------|
| *(none)* | Open the graphical window |
| `--tui` | Open the terminal UI instead |
| `--gui` | Graphical window (default; kept for compatibility) |

## Project layout

```
main.go                 Entry point, mode selection
internal/core/          Time calculations, scene model, config
internal/glyph/         Block-digit patterns (TUI only)
internal/render/gfx/    Ebitengine GUI renderer
internal/render/tui/    Bubble Tea terminal renderer
```

## Fonts

The GUI embeds [Nunito](https://fonts.google.com/specimen/Nunito) (SIL Open Font License 1.1) from `internal/render/gfx/fonts/Nunito.ttf`.

## License

Source code: no license file is included yet — add one if you plan to distribute binaries publicly.