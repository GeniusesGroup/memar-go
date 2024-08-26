/* For license and copyright information please see the LEGAL file in the code repository */

package utc

// Elapsed types specified time elapsed of January 1 of the absolute year.
// January 1 of the absolute year(0001), like January 1 of 2001, was a Monday.
type (
	CalendarYearElapsed int64
	TropicalYearElapsed int64
	MonthElapsed        int64 // utc.Now().MonthElapsed()
	WeekElapsed         int64 // utc.Now().WeekElapsed()
	DayElapsed          int64 // utc.Now().DayElapsed()
	HourElapsed         int64 // utc.Now().HourElapsed()
	MinuteElapsed       int64 // utc.Now().MinuteElapsed()
)
