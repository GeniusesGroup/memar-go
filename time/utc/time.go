/* For license and copyright information please see LEGAL file in repository */

package utc

import (
	"sync/atomic"

	"../../protocol"
	"../unix"
)

// Elapsed types specified time elapsed of January 1 of the absolute year.
// January 1 of the absolute year(0001), like January 1 of 2001, was a Monday.
type (
	MonthElapsed int64 // utc.Now().MonthElapsed()
	DayElapsed   int64 // utc.Now().DayElapsed()
	HourElapsed  int64 // utc.Now().HourElapsed()
	SecElapsed   int64 // utc.Now().SecondElapsed()
	MilliElapsed int64
	MicroElapsed int64
	NanoElapsed  int64
)

func Now() (t Time) { t.Now(); return }

type Time struct {
	sec  int64
	nsec int32
}

func (t *Time) Epoch() protocol.TimeEpoch { return protocol.TimeEpoch_UTC }
func (t *Time) SecondElapsed() int64      { return t.sec }
func (t *Time) NanoSecondElapsed() int32  { return t.nsec }
func (t *Time) ToString() string {
	// TODO:::
	return ""
}

func (t Time) NanoElapsed() NanoElapsed           { return NanoElapsed((t.sec * 1e9) + int64(t.nsec)) }
func (t Time) MicroElapsed() MicroElapsed         { return MicroElapsed((t.sec * 1e6) + int64(t.nsec/1e3)) }
func (t Time) MilliElapsed() MilliElapsed         { return MilliElapsed((t.sec * 1e3) + int64(t.nsec/1e6)) }
func (t Time) SecElapsed() SecElapsed             { return SecElapsed(t.sec) }
func (t Time) HourElapsed() (hour HourElapsed)    { return HourElapsed(t.sec / (60 * 60)) }
func (t Time) DayElapsed() (day DayElapsed)       { return DayElapsed(t.sec / (24 * 60 * 60)) }
func (t Time) MonthElapsed() (month MonthElapsed) { return MonthElapsed(t.sec / (30 * 24 * 60 * 60)) }
func (t Time) TropicalYearElapsed() (year int64)  { return t.sec / TropicalYear }
func (t Time) CallendarYearElapsed() (year int64) { return } // TODO:::

func (t *Time) ChangeTo(sec SecElapsed, nsecElapsed int32) { t.sec, t.nsec = int64(sec), nsecElapsed }
func (t *Time) Now() {
	t.sec, t.nsec, _ = unix.HardwareNow()
	// TODO:::
}
func (t *Time) NowAtomic() {
	var sec, nsec, _ = unix.HardwareNow()
	// TODO:::
	atomic.AddInt64(&t.sec, sec)
	atomic.AddInt32(&t.nsec, nsec)
}

// Until return time duration until to given time!
func (t Time) Until(to Time) (until Time) {
	until = Time{
		sec:  t.sec - to.sec,
		nsec: t.nsec - to.nsec,
	}
	return
}

// UntilTo return second duration until to given time!
func (t Time) UntilTo(to Time) (duration protocol.Duration) {
	return protocol.Duration(t.Until(to).NanoElapsed())
}

// Pass check if time pass from given time
func (t Time) Pass(from Time) (pass bool) {
	if (t.sec > from.sec) || (t.sec == from.sec && t.nsec > from.nsec) {
		pass = true
	}
	return
}

// AddDuration return given time plus given duration
func (t Time) AddDuration(d protocol.Duration) (new Time) {
	var secPass = d / Second
	new.sec += int64(secPass)
	new.nsec += int32(d % (secPass * Second))
	return
}

// Local change given time to local time by OS set time zone
func (t *Time) Local() (loc Time) {
	// TODO:::
	return
}

func (t *Time) DayHours() (hour DayHours) {
	var secPassDay = t.sec % (24 * 60 * 60)
	var dayHour = secPassDay / (60 * 60)
	hour = (1 << dayHour)
	return
}
func (t *Time) Weekdays() (day Weekdays) {
	var week = t.sec % (7 * 24 * 60 * 60)
	var weekDay = week / (24 * 60 * 60)
	// weekDay index from Thursday so change it to Monday as Weekdays
	if weekDay < 4 {
		weekDay += 3
	} else {
		weekDay -= 4 // Due to WeekdaysNone must -4 instead -3
	}
	day = (1 << weekDay)
	return
}
