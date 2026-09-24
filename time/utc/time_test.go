/* For license and copyright information please see the LEGAL file in the code repository */

package utc

import (
	"fmt"
	"memar/time/earth"
	"testing"
)

func TestRoundToHour(t *testing.T) {
	fmt.Print(9/3, "\n")
	fmt.Print(10/3, "\n")
	fmt.Print(11/3, "\n")
	fmt.Print(12/3, "\n")
	// fmt.Print(RoundToDay(1603174921), "\n")
	// fmt.Print(RoundToYear(1603174921), "\n")
}

func Test_NanoSecToWeekdays(t *testing.T) {
	var now Time
	now.Now()
	fmt.Println(now.Weekdays())
	now.AddDuration(earth.NanoSecondInDay)
	fmt.Println(now.Weekdays())
	now.AddDuration(earth.NanoSecondInDay)
	fmt.Println(now.Weekdays())
	now.AddDuration(earth.NanoSecondInDay)
	fmt.Println(now.Weekdays())
	now.AddDuration(earth.NanoSecondInDay)
	fmt.Println(now.Weekdays())
	now.AddDuration(earth.NanoSecondInDay)
	fmt.Println(now.Weekdays())
	now.AddDuration(earth.NanoSecondInDay)
	fmt.Println(now.Weekdays())
	now.AddDuration(earth.NanoSecondInDay)
	fmt.Println(now.Weekdays())
}

func Test_DayHours(t *testing.T) {
	var now Time
	now.Now()
	fmt.Printf("%b\n", now.DayHours())
	// 12 :: 1000000000000
}

func TestWeek(t *testing.T) {
	var now Time
	now.Now()
	fmt.Println(now)
	var weekDay = ((int64(now.NanoElapsed()) % int64(earth.NanoSecondInWeek)) / int64(earth.NanoSecondInDay))
	if weekDay < 4 {
		weekDay += 3
	} else {
		weekDay -= 4
	}
	fmt.Println(weekDay)
	var day uint8 = (1 << weekDay)
	fmt.Printf("%b\n", day)
}
