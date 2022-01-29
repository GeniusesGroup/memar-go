/* For license and copyright information please see LEGAL file in repository */

package time

import (
	"../protocol"
)

// Check given hour exist in given hours
func CheckDayHours(dh, checkHour protocol.DayHours) (exist bool) {
	return checkHour&dh == checkHour
}

// CheckDayHoursReverse check given hours exist in given hour
func CheckDayHoursReverse(dh, hours protocol.DayHours) (exist bool) {
	return dh&hours == dh
}
