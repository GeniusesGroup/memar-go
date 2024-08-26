/* For license and copyright information please see the LEGAL file in the code repository */

package utc

import (
	"memar/time/duration"
)

// Common durations.
const (
	Minute duration.NanoSecond = 60 * duration.OneSecond
	Hour                          = 60 * Minute
	Day                           = 24 * Hour
	Week                          = 7 * Day
	// Month                         = 2629743 * Second  // 30.44 days
	// Year                          = 31556926 * Second // 365.24 days

	// TropicalYear also known as a solar year - https://en.wikipedia.org/wiki/Tropical_year
	TropicalYear = (365 * 24 * 60 * 60) + (5 * 60 * 60) + (48 * 60) + 46 // 365.24219 * 24 * 60 * 60 = 31,556,925.216
)

const (
// secondsPerHour   = 60 * secondsPerMinute
// secondsPerDay    = 24 * secondsPerHour
// secondsPerWeek   = 7 * secondsPerDay
// daysPer400Years  = 365*400 + 97
// daysPer100Years  = 365*100 + 24
// daysPer4Years    = 365*4 + 1
)
