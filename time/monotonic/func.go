/* For license and copyright information please see LEGAL file in repository */

package monotonic

import (
	_ "unsafe" // for go:linkname
)

type (
	// A Time monotonic clock is for measuring time.
	// time-measuring operations, specifically comparisons and subtractions, use the monotonic clock.
	NanoSecElapsed int64
)

// RuntimeNano returns the current value of the runtime monotonic clock in nanoseconds.
// It isn't not wall clock, Use in tasks like timeout, ...
//go:linkname RuntimeNano runtime.nanotime
func RuntimeNano() int64

// Now returns runtime monotonic clock in nanoseconds.
func Now() NanoSecElapsed {
	return NanoSecElapsed(RuntimeNano())
}
