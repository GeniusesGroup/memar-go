/* For license and copyright information please see LEGAL file in repository */

package utc

import (
	"../../protocol"
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
	// Month                         = 2629743 * Second  // 30.44 days
	// Year                          = 31556926 * Second // 365.24 days

	// TropicalYear also known as a solar year - https://en.wikipedia.org/wiki/Tropical_year
	TropicalYear = (365 * 24 * 60 * 60) + (5 * 60 * 60) + (48 * 60) + 46 // 365.24219 * 24 * 60 * 60 = 31,556,925.216
)
