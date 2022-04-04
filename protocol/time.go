/* For license and copyright information please see LEGAL file in repository */

package protocol

// Time is the interface that must implement by any time object.
// It is base on Epoch and Second terms to work anywhere (in any planet in the universe).
// https://en.wikipedia.org/wiki/Epoch
// https://en.wikipedia.org/wiki/Second
type Time interface {
	Epoch() TimeEpoch
	SecondElapsed() int64     // From Epoch
	NanoSecondElapsed() int32 // From second

	Stringer
}

type TimeEpoch uint64

const (
	TimeEpoch_Unset TimeEpoch = iota
	TimeEpoch_Monotonic
	TimeEpoch_Unix
	TimeEpoch_UTC
)

// A Duration represents the elapsed time between two instants as an int64 nanosecond count.
// The representation limits the largest representable duration to approximately 290 earth years.
type Duration int64
