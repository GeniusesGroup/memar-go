/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/math/integer"
	"memar/time/duration"
)

type Field_Week interface {
	Week() Week
}

type Week integer.S64

// Common durations.
const (
	NanoSecondInWeek duration.NanoSecond = 7 * NanoSecondInDay
	SecondInWeek     duration.Second     = 7 * SecondInDay
	MinuteInWeek     Minute              = 7 * MinuteInDay
	HourInWeek       Hour                = 7 * HourInDay
	DayInWeek        Day                 = 7
)

func (w *Week) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*w = Week(d / NanoSecondInWeek)
}

func (w *Week) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*w = Week(d / SecondInWeek)
}
