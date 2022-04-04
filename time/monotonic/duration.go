/* For license and copyright information please see LEGAL file in repository */

package monotonic

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
)
