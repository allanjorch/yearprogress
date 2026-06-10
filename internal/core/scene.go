package core

import "fmt"

type Scene struct {
	DayLine     string
	YearPct     float64
	Year        int
	EvenDay     bool
	ShowEvenDay bool
}

func BuildScene(state TimeState, cfg Config) Scene {
	return Scene{
		DayLine:     FormatDayNumber(state.Day, state.Fraction),
		YearPct:     state.YearPct,
		Year:        state.Year,
		EvenDay:     state.EvenDay,
		ShowEvenDay: cfg.ShowEvenDayLabel,
	}
}

func (s Scene) EvenDayLabel() string {
	if s.EvenDay {
		return "EVEN DAY"
	}
	return "ODD DAY"
}

func FormatBarLabel(pct float64) string {
	return fmt.Sprintf("%.2f%%", pct*100)
}