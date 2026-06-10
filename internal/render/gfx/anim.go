package gfx

import "time"

const fullYearAnimDuration = 2 * time.Second

type barAnim struct {
	start    time.Time
	target   float64
	duration time.Duration
	done     bool
}

func newBarAnim(target float64, start time.Time) barAnim {
	d := time.Duration(target * float64(fullYearAnimDuration))
	if d <= 0 {
		return barAnim{start: start, target: target, done: true}
	}
	return barAnim{start: start, target: target, duration: d}
}

func (b *barAnim) pct(now time.Time, live float64) float64 {
	if b.done {
		return live
	}
	elapsed := now.Sub(b.start)
	if elapsed >= b.duration {
		b.done = true
		return live
	}
	return b.target * (float64(elapsed) / float64(b.duration))
}