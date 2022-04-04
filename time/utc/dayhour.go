/* For license and copyright information please see LEGAL file in repository */

package utc

// A DayHours specifies a hour of a day.
type DayHours uint32

// Hours
const (
	DayHours_None DayHours = 0
	DayHours_0 DayHours = (1 << iota)
	DayHours_1
	DayHours_2
	DayHours_3
	DayHours_4
	DayHours_5
	DayHours_6
	DayHours_7
	DayHours_8
	DayHours_9
	DayHours_10
	DayHours_11
	DayHours_12
	DayHours_13
	DayHours_14
	DayHours_15
	DayHours_16
	DayHours_17
	DayHours_18
	DayHours_19
	DayHours_20
	DayHours_21
	DayHours_22
	DayHours_23
	DayHours_All DayHours = 0b11111111111111111111111111111111
)

// Check given hour exist in given day hours
func (dh DayHours) Check(hour DayHours) (exist bool) { return hour&dh != 0 }
