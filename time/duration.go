/* For license and copyright information please see LEGAL file in repository */

package time

import (
	"../protocol"
)

// Common durations.
const (
	Nanosecond  protocol.Duration = 1
	Microsecond                   = 1000 * Nanosecond
	Millisecond                   = 1000 * Microsecond
	Second                        = 1000 * Millisecond
	Minute                        = 60 * Second
	Hour                          = 60 * Minute
	Day                           = 24 * Hour
	Week                          = 7 * Day
	Month                         = 2629743 * Second  // 30.44 days
	Year                          = 31556926 * Second // 365.24 days
)

const (
// secondsPerHour   = 60 * secondsPerMinute
// secondsPerDay    = 24 * secondsPerHour
// secondsPerWeek   = 7 * secondsPerDay
// daysPer400Years  = 365*400 + 97
// daysPer100Years  = 365*100 + 24
// daysPer4Years    = 365*4 + 1
)
