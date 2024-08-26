/* For license and copyright information please see the LEGAL file in the code repository */

package utc

import (
	"memar/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
	"memar/time/unix"
)

func Now() (t Time) { t.Now(); return }

type Time struct {
	sec  duration.Second
	nsec duration.NanoInSecond
}

//memar:impl memar/time/protocol.Time
func (t *Time) Epoch() time_p.Epoch                        { return &Epoch }
func (t *Time) SecondElapsed() duration.Second             { return t.sec }
func (t *Time) NanoInSecondElapsed() duration.NanoInSecond { return t.nsec }

//memar:impl memar/protocol.Stringer
func (t *Time) ToString() (str string, err protocol.Error) {
	// TODO:::
	return
}
func (t *Time) FromString(str string) (err protocol.Error) {
	// TODO:::
	return
}

func (t Time) NanoElapsed() (d duration.NanoSecond) {
	d.FromSecAndNano(t.sec, t.nsec)
	return
}
func (t Time) MicroElapsed() (d duration.MicroSecond) {
	d.FromSecAndNano(t.sec, t.nsec)
	return
}
func (t Time) MilliElapsed() (d duration.MilliSecond) {
	d.FromSecAndNano(t.sec, t.nsec)
	return
}
func (t Time) SecElapsed() duration.Second  { return t.sec }
func (t Time) MinuteElapsed() MinuteElapsed { return MinuteElapsed(t.sec / 60) }
func (t Time) HourElapsed() HourElapsed     { return HourElapsed(t.sec / (60 * 60)) }
func (t Time) DayElapsed() DayElapsed       { return DayElapsed(t.sec / (24 * 60 * 60)) }
func (t Time) WeekElapsed() WeekElapsed     { return WeekElapsed(t.sec / (7 * 24 * 60 * 60)) }
func (t Time) MonthElapsed() MonthElapsed   { return MonthElapsed(t.sec / (30 * 24 * 60 * 60)) }
func (t Time) TropicalYearElapsed() TropicalYearElapsed {
	return TropicalYearElapsed(t.sec / TropicalYear)
}
func (t Time) CalendarYearElapsed() CalendarYearElapsed { return 0 } // TODO:::

func (t Time) NanoElapsedSafe() bool {
	// TODO:::
	return false
}

func (t *Time) ChangeTo(sec duration.Second, nsecElapsed duration.NanoInSecond) {
	t.sec, t.nsec = sec, nsecElapsed
}
func (t *Time) Now() {
	var ut = unix.Now()
	t.sec, t.nsec = ut.SecondElapsed(), ut.NanoInSecondElapsed()
}

// func (t *Time) NowAtomic() {
// 	var sec, nsec = now()
// 	atomic.AddInt64(&t.sec, sec)
// 	atomic.AddInt32(&t.nsec, nsec)
// }

// Until return time duration until to given time!
func (t Time) Until(to Time) (until Time) {
	until = Time{
		sec:  t.sec - to.sec,
		nsec: t.nsec - to.nsec,
	}
	return
}

// UntilTo return second duration until to given time!
func (t Time) UntilTo(to Time) duration.NanoSecond {
	return duration.NanoSecond(t.Until(to).NanoElapsed())
}

// Pass check if time pass from given time
func (t Time) Pass(from Time) (pass bool) {
	if (t.sec > from.sec) || (t.sec == from.sec && t.nsec > from.nsec) {
		pass = true
	}
	return
}

// AddDuration return given time plus given duration
func (t *Time) AddDuration(d duration.NanoSecond) {
	var sec, nsec = d.ToSecAndNano()
	t.sec += sec
	t.nsec += nsec
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
