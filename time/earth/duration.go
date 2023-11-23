/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/protocol"
)

const (
	// Common durations.
	Nanosecond  protocol.Duration = 1
	Microsecond                   = 1000 * Nanosecond
	Millisecond                   = 1000 * Microsecond
	Second                        = 1000 * Millisecond

	// Earth specific durations
	Minute = 60 * Second
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 2629743 * Second  // 30.44 days
	Year   = 31556926 * Second // 365.24 days
)
