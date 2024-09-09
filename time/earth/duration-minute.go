/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/time/duration"
)

type Minute int64

// Common durations.
const (
	SecondInMinute     duration.Second     = 60
	NanoSecondInMinute duration.NanoSecond = 60 * duration.OneSecond

	MinuteInHour Minute = 60
	MinuteInDay  Minute = 24 * MinuteInHour
	MinuteInWeek Minute = 7 * MinuteInDay
)

func (m *Minute) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*m = Minute(d / NanoSecondInMinute)
}

func (m *Minute) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*m = Minute(d / SecondInMinute)
}
