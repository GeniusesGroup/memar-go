/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/time/duration"
)

type Week int64

// Common durations.
const (
	NanoSecondInWeek duration.NanoSecond = 7 * NanoSecondInDay
	SecondInWeek     duration.Second     = 7 * SecondInDay
)

func (w *Week) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*w = Week(d / NanoSecondInWeek)
}

func (w *Week) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*w = Week(d / SecondInWeek)
}
