/* For license and copyright information please see LEGAL file in repository */

package earth

// A DayHours specifies a hour of a day.
type DayHours uint64

// Hours
const (
	DayHours_None DayHours = 0
	DayHours_0    DayHours = (1 << iota)
	DayHours_0_Half
	DayHours_1
	DayHours_1_Half
	DayHours_2
	DayHours_2_Half
	DayHours_3
	DayHours_3_Half
	DayHours_4
	DayHours_4_Half
	DayHours_5
	DayHours_5_Half
	DayHours_6
	DayHours_6_Half
	DayHours_7
	DayHours_7_Half
	DayHours_8
	DayHours_8_Half
	DayHours_9
	DayHours_9_Half
	DayHours_10
	DayHours_10_Half
	DayHours_11
	DayHours_11_Half
	DayHours_12
	DayHours_12_Half
	DayHours_13
	DayHours_13_Half
	DayHours_14
	DayHours_14_Half
	DayHours_15
	DayHours_15_Half
	DayHours_16
	DayHours_16_Half
	DayHours_17
	DayHours_17_Half
	DayHours_18
	DayHours_18_Half
	DayHours_19
	DayHours_19_Half
	DayHours_20
	DayHours_20_Half
	DayHours_21
	DayHours_21_Half
	DayHours_22
	DayHours_22_Half
	DayHours_23
	DayHours_23_Half
	DayHours_All = ^(DayHours(0))
)

// Check given hour exist in given day hours
func (dh DayHours) Check(hour DayHours) (exist bool) { return hour&dh != 0 }
