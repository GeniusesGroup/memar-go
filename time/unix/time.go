/* For license and copyright information please see the LEGAL file in the code repository */

/*
Unix time specifies time(second) elapsed of January 1 of the absolute year.
January 1 of the absolute year(1970), like January 1 of 2001, was a Monday.
*/
package unix

import (
	"memar/protocol"
)

const Base = "00:00:00 UTC on 1 January 1970" // Thursday

type (
	DayElapsed   int64 // fast way: unix.Now().DayElapsed()
	HourElapsed  int64 // fast way: unix.Now().HourElapsed()
	SecElapsed   int64 // fast way: unix.Now().SecElapsed()		|| Go way: time.Now().Unix()
	MilliElapsed int64 // fast way: unix.Now().MilliElapsed()	|| Go way: time.Now().UnixMilli()
	MicroElapsed int64 // fast way: unix.Now().MicroElapsed()	|| Go way: time.Now().UnixMicro()
	NanoElapsed  int64 // fast way: unix.Now().NanoElapsed()	|| Go way: time.Now().UnixNano()
)

func Now() (t Time) { t.Now(); return }

// A Time specifies second elapsed of January 1 of the absolute year.
// January 1 of the absolute year(1970), like January 1 of 2001, was a Monday.
type Time struct {
	sec  int64
	nsec int32
}

//memar:impl memar/protocol.Time
func (t *Time) Epoch() protocol.TimeEpoch { return protocol.TimeEpoch_Unix }
func (t *Time) SecondElapsed() int64      { return t.sec }
func (t *Time) NanoSecondElapsed() int32  { return t.nsec }

//memar:impl memar/protocol.Stringer
func (t *Time) ToString() string {
	// TODO:::
	return ""
}
func (t *Time) FromString(s string) (err protocol.Error) {
	// TODO:::
	return
}

func (t *Time) ChangeTo(sec SecElapsed, nsecElapsed int32) { t.sec, t.nsec = int64(sec), nsecElapsed }
func (t *Time) Now()                                       { t.sec, t.nsec, _ = now() }
func (t Time) Pass(from Time) (pass bool) {
	if (t.sec > from.sec) || (t.sec == from.sec && t.nsec > from.nsec) {
		pass = true
	}
	return
}
func (t Time) PassNow() bool { return t.Pass(Now()) }
func (t *Time) Add(d protocol.Duration) {
	var sec, nsec = nsecToSec(d)
	t.sec += sec
	t.nsec += nsec
}
func (t Time) Since(baseTime Time) (d protocol.Duration) {
	d = protocol.Duration(baseTime.sec-t.sec) * Second
	d += protocol.Duration(baseTime.nsec - t.nsec)
	return
}
func (t Time) SinceNow() (d protocol.Duration) { return t.Since(Now()) }

func (t Time) NanoElapsed() NanoElapsed   { return NanoElapsed((t.sec * 1e9) + int64(t.nsec)) }
func (t Time) MicroElapsed() MicroElapsed { return MicroElapsed((t.sec * 1e6) + int64(t.nsec/1e3)) }
func (t Time) MilliElapsed() MilliElapsed { return MilliElapsed((t.sec * 1e3) + int64(t.nsec/1e6)) }
func (t Time) SecElapsed() SecElapsed     { return SecElapsed(t.sec) }
func (t Time) HourElapsed() HourElapsed   { return HourElapsed(t.sec / (60 * 60)) }
func (t Time) DayElapsed() DayElapsed     { return DayElapsed(t.sec / (24 * 60 * 60)) }

// ElapsedByDuration returns the result of rounding t to the nearest multiple of d.
func (t Time) ElapsedByDuration(d protocol.Duration) (period int64) {
	if d < 0 {
		return
	}
	if d == 0 {
		return int64(t.NanoElapsed())
	}
	var sec, nsec = nsecToSec(d)
	if sec > 0 {
		period = t.sec / sec
		if nsec > 0 {
			period += (int64(t.nsec/nsec) / int64(Second))
		}
	} else {
		period = (t.sec * int64(Second)) / int64(nsec)
		period += int64(t.nsec / nsec)
	}
	return
}
