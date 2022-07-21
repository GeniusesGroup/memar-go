/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"../protocol"
	"../time/monotonic"
)

// when is a helper function for setting the 'when' field of a Timer.
// It returns what the time will be, in nanoseconds, Duration d in the future.
// If d is negative, it is ignored. If the returned value would be less than
// zero because of an overflow, MaxInt64 is returned.
func when(d protocol.Duration) (t int64) {
	t = monotonic.RuntimeNano()
	if d <= 0 {
		return
	}
	t += int64(d)
	// check for overflow.
	if t < 0 {
		// N.B. monotonic.RuntimeNano() and d are always positive, so addition
		// (including overflow) will never result in t == 0.
		t = maxWhen
	}
	return
}

// badTimer is called if the timer data structures have been corrupted,
// presumably due to racy use by the program. We panic here rather than
// panicing due to invalid slice access while holding locks.
// See issue #25686.
func badTimer() {
	panic("timers: data corruption")
}
