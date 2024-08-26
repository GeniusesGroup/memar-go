/* For license and copyright information please see the LEGAL file in the code repository */

package unix

import (
	"memar/time/duration"
)

// Common durations.
const (
	Minute duration.NanoSecond = 60 * duration.OneSecond
	Hour                       = 60 * Minute
	Day                        = 24 * Hour
	Week                       = 7 * Day
	Month                      = 2629743 * duration.OneSecond  // 30.44 days
	Year                       = 31556926 * duration.OneSecond // 365.24 days
)

type (
	DayElapsed   int64 // fast way: unix.Now().DayElapsed()
	HourElapsed  int64 // fast way: unix.Now().HourElapsed()
)
