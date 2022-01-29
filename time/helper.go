/* For license and copyright information please see LEGAL file in repository */

package time

import (
	"../protocol"
)

// NanoSecToMilliSec delete nano-seconds parts of time to have unique time in each milli instead nano second elapsed after protocol.TimeUnixBase.
func NanoSecToMilliSec(nsec protocol.TimeUnixNano) (rounded protocol.TimeUnixMilli) {
	return protocol.TimeUnixMilli(int64(nsec) / int64(Millisecond))
}

// NanoSecToDay delete seconds parts of time to have unique time in each day instead nano second elapsed after protocol.TimeUnixBase.
func NanoSecToDay(nsec protocol.TimeUnixNano) (rounded protocol.TimeUnixDay) {
	return protocol.TimeUnixDay(int64(nsec) / int64(Day))
}

// NanoSecToWeekdays return Weekdays of given time.
func NanoSecToWeekdays(nsec protocol.TimeUnixNano) (day protocol.Weekdays) {
	var week = int64(nsec) % int64(Week)
	var weekDay = week / int64(Day)
	// weekDay index from Thursday so change it to Monday as protocol.Weekdays
	if weekDay < 4 {
		weekDay += 3
	} else {
		weekDay -= 4 // Due to WeekdaysNone must -4 instead -3
	}
	day = (1 << weekDay)
	return
}

// NanoSecToDayHours return Hour of the day in the given time by DayHours format.
func NanoSecToDayHours(nsec protocol.TimeUnixNano) (hour protocol.DayHours) {
	var nsecPassDay = int64(nsec) % int64(Day)
	var dayHour = nsecPassDay / int64(Hour)
	hour = (1 << dayHour)
	return
}
