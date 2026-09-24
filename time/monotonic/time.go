/* For license and copyright information please see the LEGAL file in the code repository */

package monotonic

import (
	error_p "memar/process/error/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
)

// Now returns runtime monotonic clock in nanoseconds.
func Now() Time {
	return Time(now())
}

// A Time monotonic clock is for measuring time.
// time-measuring operations, specifically comparisons and subtractions, use the monotonic clock.
type Time duration.NanoSecond

//memar:impl memar/time/protocol.Time
func (self *Time) Epoch() time_p.Epoch { return &Epoch }
func (self *Time) SecondElapsed() (sec duration.Second) {
	sec.FromNanoSecond(duration.NanoSecond(*self))
	return
}
func (self *Time) NanoInSecondElapsed() (d duration.NanoInSecond) {
	d.FromSec(duration.NanoSecond(*self))
	return
}

//memar:impl memar/codec/string/protocol.Stringer
func (self *Time) ToString() (s string, err error_p.Error) {
	s = "TODO:::"
	return
}

//memar:impl memar/codec/string/protocol.Stringer
func (self *Time) FromString(str string) (err error_p.Error) {
	// TODO:::
	return
}

func (self *Time) Now()                      { *self = Now() }
func (self *Time) Add(d duration.NanoSecond) { *self += Time(d) }

// Equal reports whether given time represent the same time instant.
func (self Time) Equal(other Time) bool { return self == other }

// Pass reports whether the time instant is after given time.
func (self Time) Pass(from Time) bool { return self > from }

// PassNow reports whether the time instant is after now.
func (self Time) PassNow() bool { return self > Now() }

// Since returns the time elapsed since t.
func (self Time) Since(from Time) (d duration.NanoSecond) { return duration.NanoSecond(from - self) }

// SinceNow returns the time elapsed since now.
func (self Time) SinceNow() (d duration.NanoSecond) { return duration.NanoSecond(Now() - self) }

// Until returns the duration t(the future) until to (before the t).
func (self Time) Until(to Time) (d duration.NanoSecond) { return duration.NanoSecond(self - to) }

// UntilNow returns the duration t(the future) until now.
func (self Time) UntilNow() (d duration.NanoSecond) { return duration.NanoSecond(self - Now()) }
