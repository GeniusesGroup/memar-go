/* For license and copyright information please see LEGAL file in repository */

package monotonic

import (
	_ "unsafe" // for go:linkname

	"github.com/GeniusesGroup/libgo/protocol"
)

// RuntimeNano returns the current value of the runtime monotonic clock in nanoseconds.
// It isn't not wall clock, Use in tasks like timeout, ...
//go:linkname RuntimeNano runtime.nanotime
func RuntimeNano() int64

// Now returns runtime monotonic clock in nanoseconds.
func Now() Time {
	return Time(RuntimeNano())
}

// A Time monotonic clock is for measuring time.
// time-measuring operations, specifically comparisons and subtractions, use the monotonic clock.
type Time int64

// protocol.Time interface methods
func (t *Time) Epoch() protocol.TimeEpoch { return protocol.TimeEpoch_Monotonic }
func (t *Time) SecondElapsed() int64      { return int64(*t) / int64(Second) }
func (t *Time) NanoSecondElapsed() int32  { return int32(int64(*t) % t.SecondElapsed()) }
func (t *Time) ToString() string {
	// TODO:::
	return ""
}

func (t *Time) Now()                        { *t = Now() }
func (t Time) Pass(baseTime Time) bool      { return baseTime > t }
func (t Time) PassNow() bool                { return Now() > t }
func (t *Time) Add(d protocol.Duration)     { *t += Time(d) }
func (t Time) Since() (d protocol.Duration) { return protocol.Duration(Now() - t) }
