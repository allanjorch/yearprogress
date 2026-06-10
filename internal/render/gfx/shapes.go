package gfx

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func fillRoundRect(dst *ebiten.Image, x, y, w, h, r float32, clr color.Color) {
	if w <= 0 || h <= 0 {
		return
	}
	if r > w/2 {
		r = w / 2
	}
	if r > h/2 {
		r = h / 2
	}

	var path vector.Path
	path.MoveTo(x+r, y)
	path.LineTo(x+w-r, y)
	path.ArcTo(x+w, y, x+w, y+r, r)
	path.LineTo(x+w, y+h-r)
	path.ArcTo(x+w, y+h, x+w-r, y+h, r)
	path.LineTo(x+r, y+h)
	path.ArcTo(x, y+h, x, y+h-r, r)
	path.LineTo(x, y+r)
	path.ArcTo(x, y, x+r, y, r)
	path.Close()

	op := &vector.DrawPathOptions{}
	op.ColorScale.ScaleWithColor(clr)
	vector.FillPath(dst, &path, nil, op)
}
