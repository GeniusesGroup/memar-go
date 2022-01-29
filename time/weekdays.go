/* For license and copyright information please see LEGAL file in repository */

package time

import (
	"../protocol"
)

// Check given day exist in desire days
func CheckWeekdays(weekdays, day protocol.Weekdays) (exist bool) {
	return day&weekdays == day
}

// CheckReverse given days exist in desire day
func CheckWeekdaysReverse(day, weekdays protocol.Weekdays) (exist bool) {
	return day&weekdays == day
}
