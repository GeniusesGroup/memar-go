/* For license and copyright information please see LEGAL file in repository */

package time

import (
	_ "unsafe" // for go:linkname

	"../protocol"
)

// Provided by package runtime.
//go:linkname UnixNow time.now
func UnixNow() (sec int64, nsec int32, mono int64)

// RuntimeNano returns the current value of the runtime monotonic clock in nanoseconds.
// It isn't not wall clock, Use in tasks like timeout, ...
//go:linkname RuntimeNano runtime.nanotime
func RuntimeNano() int64

// RuntimeMonotonic returns runtime monotonic clock in nanoseconds.
func RuntimeMonotonic() protocol.RuntimeMonotonic {
	return protocol.RuntimeMonotonic(RuntimeNano())
}

// UnixNowNano returns earth time (UTC) in nano-second elapsed after protocol.TimeUnixBase
func UnixNowNano() protocol.TimeUnixNano {
	var sec, nsec, _ = UnixNow()
	return protocol.TimeUnixNano((sec * 1e9) + int64(nsec))
}

// UnixNowNano returns earth time (UTC) in micro-second elapsed after protocol.TimeUnixBase
func UnixNowMicro() protocol.TimeUnixMicro {
	var sec, nsec, _ = UnixNow()
	return protocol.TimeUnixMicro((sec * 1e6) + int64(nsec/1e3))
}

// UnixNowNano returns earth time (UTC) in milli-second elapsed after protocol.TimeUnixBase
func UnixNowMilli() protocol.TimeUnixMilli {
	var sec, nsec, _ = UnixNow()
	return protocol.TimeUnixMilli((sec * 1e3) + int64(nsec/1e6))
}

// UnixNowSecond returns earth time (UTC) in second elapsed after protocol.TimeUnixBase
func UnixNowSecond() protocol.TimeUnixSec {
	var sec, _, _ = UnixNow()
	return protocol.TimeUnixSec(sec)
}

// UnixNowNano returns earth time (UTC) in hour elapsed after protocol.TimeUnixBase
func UnixNowHour() protocol.TimeUnixHour {
	var sec, _, _ = UnixNow()
	return protocol.TimeUnixHour(sec / (60 * 60))
}

// UnixNowDay returns earth time (UTC) in day elapsed after protocol.TimeUnixBase
func UnixNowDay() protocol.TimeUnixDay {
	var sec, _, _ = UnixNow()
	return protocol.TimeUnixDay(sec / (24 * 60 * 60))
}
