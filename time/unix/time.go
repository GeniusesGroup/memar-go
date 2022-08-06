/* For license and copyright information please see LEGAL file in repository */

package unix

import (
	"sync/atomic"
	_ "unsafe" // for go:linkname

	"github.com/GeniusesGroup/libgo/protocol"
)

const Base = "00:00:00 UTC on 1 January 1970" // Thursday

// UnixTime specifies time elapsed of January 1 of the absolute year.
// January 1 of the absolute year(1970), like January 1 of 2001, was a Monday.
type (
	DayElapsed   int64 // fast way: unix.Now().DayElapsed()
	HourElapsed  int64 // fast way: unix.Now().HourElapsed()
	SecElapsed   int64 // fast way: unix.Now().SecElapsed()		|| Go way: time.Now().Unix()
	MilliElapsed int64 // fast way: unix.Now().MilliElapsed()	|| Go way: time.Now().UnixMilli()
	MicroElapsed int64 // fast way: unix.Now().MicroElapsed()	|| Go way: time.Now().UnixMicro()
	NanoElapsed  int64 // fast way: unix.Now().NanoElapsed()	|| Go way: time.Now().UnixNano()
)

func Now() (t Time) { t.Now(); return }

// Provided by package runtime.
//
//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)

// A Time specifies second elapsed of January 1 of the absolute year.
// January 1 of the absolute year(1970), like January 1 of 2001, was a Monday.
type Time struct {
	sec  int64
	nsec int32
}

func (t *Time) Epoch() protocol.TimeEpoch { return protocol.TimeEpoch_Unix }
func (t *Time) SecondElapsed() int64      { return t.sec }
func (t *Time) NanoSecondElapsed() int32  { return t.nsec }
func (t *Time) ToString() string {
	// TODO:::
	return ""
}

func (t *Time) ChangeTo(sec SecElapsed, nsecElapsed int32) { t.sec, t.nsec = int64(sec), nsecElapsed }
func (t *Time) Now()                                       { t.sec, t.nsec, _ = now() }
func (t *Time) NowAtomic() {
	var sec, nsec, _ = now()
	atomic.AddInt64(&t.sec, sec)
	atomic.AddInt32(&t.nsec, nsec)
}

func (t Time) Pass(from Time) (pass bool) {
	if (t.sec > from.sec) || (t.sec == from.sec && t.nsec > from.nsec) {
		pass = true
	}
	return
}
func (t Time) AddDuration(d protocol.Duration) (new Time) {
	var secPass = d / Second
	new.sec += int64(secPass)
	new.nsec += int32(d % (secPass * Second))
	return
}

func (t Time) NanoElapsed() NanoElapsed   { return NanoElapsed((t.sec * 1e9) + int64(t.nsec)) }
func (t Time) MicroElapsed() MicroElapsed { return MicroElapsed((t.sec * 1e6) + int64(t.nsec/1e3)) }
func (t Time) MilliElapsed() MilliElapsed { return MilliElapsed((t.sec * 1e3) + int64(t.nsec/1e6)) }
func (t Time) SecElapsed() SecElapsed     { return SecElapsed(t.sec) }
func (t Time) HourElapsed() HourElapsed   { return HourElapsed(t.sec / (60 * 60)) }
func (t Time) DayElapsed() DayElapsed     { return DayElapsed(t.sec / (24 * 60 * 60)) }
