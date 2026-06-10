# yearprogress — functional specification

## Overview

`yearprogress` displays how far through the current calendar year you are. It renders two synchronized views of the same data:

1. **Day counter** — day-of-year with live sub-day precision
2. **Year progress bar** — elapsed fraction of the year as a horizontal bar with percentage label

Two renderers share one core model:

| Renderer | Package | Technology |
|----------|---------|------------|
| GUI (default) | `internal/render/gfx` | Ebitengine v2 |
| TUI | `internal/render/tui` | Bubble Tea + Lip Gloss |

## Time model

All calculations use the system's local timezone (`time.Now().Location()`).

### Day counter

- **Integer part:** `time.YearDay()` — day 1 through 365/366
- **Fractional part:** seconds elapsed since local midnight ÷ 86,400
- **Display format:** `D.FFFFF` — one integer, dot, five decimal digits  
  Example: `161.50000` at local noon on day 161

### Year percentage

- **Start:** January 1, 00:00:00 local time
- **End:** January 1 next year, 00:00:00 local time (handles leap years via `AddDate`)
- **Value:** `0.0` at year start → `1.0` at year end
- **Bar label:** `%.2f%%` — two decimal places, no year suffix  
  Example: `44.12%`

### Even / odd day

- `EvenDay = (day % 2 == 0)`
- Label text: `EVEN DAY` or `ODD DAY`
- **Hidden by default** (`Config.ShowEvenDayLabel = false`)
- Logic remains in core; enabling the label is a config-only change

## Display elements

### 1. Day number

- Dominant visual element, centered horizontally
- GUI: Nunito semi-bold (weight 600), bright chalk yellow on dark background
- TUI: 7-row block glyphs (`█` patterns), one column gap between characters

### 2. Even-day label (optional)

- Shown below the day counter when `ShowEvenDayLabel` is true
- GUI: Nunito regular, ~28% of day font size
- Color: mid yellow when odd, bright yellow when even
- TUI: plain text line

### 3. Year progress bar

- Width matches the day counter width (GUI: measured text width; TUI: glyph column count)
- Three rows in TUI; single pill bar in GUI

#### Track

- Full-width rounded rectangle (pill shape)
- Dim chalk brown/yellow

#### Fill

- Rounded on both ends (capsule segment inside the track)
- Bright chalk yellow
- Width = `yearPct × barWidth`
- No partial cells in TUI — filled columns are whole `█` blocks only

#### Percentage label

- Centered inside the bar
- GUI: Nunito bold (weight 700), per-character coloring:
  - **On fill:** dark text directly on bright fill (no pill background)
  - **On track:** bright text on dim pill background behind each character
- TUI: inverted ANSI colors on fill (`bright bg + dark fg`), screen bg + bright fg on track

## GUI-specific behavior

### Window

- Default size: 1280 × 720
- Resizable
- Title: `yearprogress`
- Layout recomputes on resize; day font scales to fit available space

### Launch animation

- Plays once at startup
- Bar and label animate from `0%` to the current year percentage
- **Duration:** `currentPct × 2 seconds` (a full year = 2 seconds)
- Linear interpolation
- Label digits update during animation; fill/track text coloring follows the animated fill edge
- After completion, switches to live `YearPct` (updates every 50 ms)

### Input

- Exit keys: Esc, Enter, Space, Q, Backspace (see `core.ExitKeys`)

### Color palette (GUI)

| Name | RGB | Usage |
|------|-----|-------|
| Background | `#141310` | Screen fill |
| Dim | `#726242` | Bar track, off-fill label pills |
| Mid | `#CCB468` | Odd-day label |
| Bright | `#FFEC88` | Day number, bar fill, on-track label text |

## TUI-specific behavior

### Rendering

- Alternate screen buffer
- ANSI true-color (24-bit RGB)
- Batched color spans to minimize reset sequences

### Input

- Exit keys: Esc, Enter, Space, Q, Backspace (see `core.ExitKeys`)

### Update interval

- 50 ms tick, same as GUI scene refresh

## Configuration

```go
type Config struct {
    ShowEvenDayLabel bool  // default: false
}
```

No CLI flags or config file yet — defaults are set in `core.DefaultConfig()`.

## CLI

| Invocation | Behavior |
|------------|----------|
| `yearprogress` | GUI |
| `yearprogress --tui` | TUI |
| `yearprogress --gui` | GUI (alias, default) |

## Update rates and perceived motion

| Element | Typical change interval |
|---------|-------------------------|
| Day decimals (5 places) | ~0.86 seconds |
| Bar label 1st decimal (0.1%) | ~5.3 minutes |
| Bar label 2nd decimal (0.01%) | ~52.6 minutes (non-leap year) |
| Even/odd day label | At local midnight |

## Architecture

```
main.go
  └─ core.ComputeTimeState(now)
  └─ core.BuildScene(state, config)
       ├─ gfx.Draw(scene, layout, barPct)   [GUI]
       └─ tui.Render(scene, w, h)           [TUI]
```

Core is renderer-agnostic. Glyph data is TUI-only. Font embedding is GUI-only.

## Dependencies

- `github.com/hajimehoshi/ebiten/v2` — GUI
- `github.com/charmbracelet/bubbletea` — TUI framework
- `github.com/charmbracelet/lipgloss` — TUI layout
- `golang.org/x/image` — font rendering (via Ebitengine text/v2)