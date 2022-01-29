/* For license and copyright information please see LEGAL file in repository */

package protocol

// Time is the interface that must implement by any time object
type Time interface {
	// Unix() TimeUnixSec
	Pass(baseTime Time) (pass bool)
	AddDuration(d Duration) (new Time)
}

type (
	// A monotonic clock is for measuring time.
	// time-measuring operations, specifically comparisons and subtractions, use the monotonic clock.
	RuntimeMonotonic int64

	// A Duration represents the elapsed time between two instants
	// as an int64 nanosecond count. The representation limits the
	// largest representable duration to approximately 290 years.
	Duration int64
)

const TimeUnixBase = "00:00:00 UTC on 1 January 1970"

// UnixTime specifies time elapsed of January 1 of the absolute year.
// January 1 of the absolute year(1970), like January 1 of 2001, was a Monday.
type (
	TimeUnixDay   int64 // fast way: time.UnixNowDay() this repo
	TimeUnixHour  int64 // fast way: time.UnixNowHour() this repo
	TimeUnixSec   int64 // fast way: time.UnixNowMicro() this repo || regular way time.Now().Unix()
	TimeUnixMilli int64 // fast way: time.UnixNowMilli() this repo || regular way time.Now().UnixMilli()
	TimeUnixMicro int64 // fast way: time.UnixNowNano() this repo || regular way time.Now().UnixMicro()
	TimeUnixNano  int64 // fast way: time.UnixNowNano() this repo || regular way time.Now().UnixNano()
)
