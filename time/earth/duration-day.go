/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/time/duration"
)

// fast way: unix.Now().DayElapsed()
type Day int64

// Common day durations
const (
	NanoSecondInDay duration.NanoSecond = 24 * NanoSecondInHour
	SecondInDay     duration.Second     = 24 * SecondInHour

	DayInWeek Day = 7
)

func (day *Day) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*day = Day(d / NanoSecondInDay)
}

func (day *Day) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*day = Day(d / SecondInDay)
}
