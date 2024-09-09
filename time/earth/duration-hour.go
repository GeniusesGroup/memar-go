/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/time/duration"
)

type Hour int64

// Common durations.
const (
	NanoSecondInHour duration.NanoSecond = 60 * NanoSecondInMinute
	SecondInHour     duration.Second     = 60 * SecondInMinute

	HourInDay  Hour = 24
	HourInWeek Hour = 7 * HourInDay
)

func (h *Hour) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*h = Hour(d / NanoSecondInHour)
}

func (h *Hour) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*h = Hour(d / SecondInHour)
}
