/* For license and copyright information please see LEGAL file in repository */

package etime

import "time"

// A Duration store
type Duration int64

// ConvertToTimeDuration add some data as time.Duration base is nanosecond not second.
func (d Duration) ConvertToTimeDuration() (duration time.Duration) {
	return time.Duration(d) * time.Second
}

// Common durations.
const (
	Nanosecond  Duration = 1000 * Microsecond
	Microsecond          = 1000 * Millisecond
	Millisecond          = 1000
	Second               = 1
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	Day                  = 24 * Hour
	Week                 = 7 * Day
)

const (
// secondsPerHour   = 60 * secondsPerMinute
// secondsPerDay    = 24 * secondsPerHour
// secondsPerWeek   = 7 * secondsPerDay
// daysPer400Years  = 365*400 + 97
// daysPer100Years  = 365*100 + 24
// daysPer4Years    = 365*4 + 1
)
