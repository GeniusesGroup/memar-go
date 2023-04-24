/* For license and copyright information please see the LEGAL file in the code repository */

package monotonic

import (
	"libgo/protocol"
)

// Now returns runtime monotonic clock in nanoseconds.
func Now() Time {
	return Time(now())
}

// A Time monotonic clock is for measuring time.
// time-measuring operations, specifically comparisons and subtractions, use the monotonic clock.
type Time int64

//libgo:impl /libgo/protocol.Time
func (t *Time) Epoch() protocol.TimeEpoch { return protocol.TimeEpoch_Monotonic }
func (t *Time) SecondElapsed() int64      { return int64(*t) / int64(Second) }
func (t *Time) NanoSecondElapsed() int32  { return int32(int64(*t) % t.SecondElapsed()) }

//libgo:impl /libgo/protocol.Stringer
func (t *Time) ToString() string {
	return "TODO:::"
}

//libgo:impl /libgo/protocol.Stringer
func (t *Time) FromString(s string) (err protocol.Error) {
	// TODO:::
	return
}

func (t *Time) Now()                    { *t = Now() }
func (t *Time) Add(d protocol.Duration) { *t += Time(d) }

// Equal reports whether t and other represent the same time instant.
func (t Time) Equal(other Time) bool { return t == other }

// Pass reports whether the time instant t is after from.
func (t Time) Pass(from Time) bool { return t > from }

// PassNow reports whether the time instant t is after now.
func (t Time) PassNow() bool { return t > Now() }

// Since returns the time elapsed since t.
func (t Time) Since(from Time) (d protocol.Duration) { return protocol.Duration(from - t) }

// SinceNow returns the time elapsed since now.
func (t Time) SinceNow() (d protocol.Duration) { return protocol.Duration(Now() - t) }

// Until returns the duration until to.
func (t Time) Until(to Time) (d protocol.Duration) { return protocol.Duration(t - to) }

// UntilNow returns the duration until now.
func (t Time) UntilNow() (d protocol.Duration) { return protocol.Duration(t - Now()) }
