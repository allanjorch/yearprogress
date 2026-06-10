package core

import (
	"fmt"
	"time"
)

type TimeState struct {
	Year     int
	Day      int
	Fraction float64
	YearPct  float64
	EvenDay  bool
}

func ComputeTimeState(now time.Time) TimeState {
	loc := now.Location()
	year := now.Year()

	startOfDay := time.Date(year, now.Month(), now.Day(), 0, 0, 0, 0, loc)
	day := now.YearDay()
	fraction := now.Sub(startOfDay).Seconds() / 86400.0

	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, loc)
	yearEnd := yearStart.AddDate(1, 0, 0)
	yearPct := now.Sub(yearStart).Seconds() / yearEnd.Sub(yearStart).Seconds()

	return TimeState{
		Year:     year,
		Day:      day,
		Fraction: fraction,
		YearPct:  yearPct,
		EvenDay:  day%2 == 0,
	}
}

func FormatDayNumber(day int, fraction float64) string {
	frac := int(fraction * 1e5)
	return fmt.Sprintf("%d.%05d", day, frac)
}