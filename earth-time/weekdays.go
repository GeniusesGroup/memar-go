/* For license and copyright information please see LEGAL file in repository */

package etime

import (
	lang "../language"
)

// A Weekdays specifies a day of the week.
type Weekdays uint8

// Weekdays
const (
	WeekdaysNone      Weekdays = 0b00000000
	WeekdaysMonday    Weekdays = 0b00000001
	WeekdaysTuesday   Weekdays = 0b00000010
	WeekdaysWednesday Weekdays = 0b00000100
	WeekdaysThursday  Weekdays = 0b00001000
	WeekdaysFriday    Weekdays = 0b00010000
	WeekdaysSaturday  Weekdays = 0b00100000
	WeekdaysSunday    Weekdays = 0b01000000
	WeekdaysAll       Weekdays = 0b11111111
)

// Set desire weekdays to given Weekdays!
// e.g. d.Set(WeekdaysSaturday|WeekdaysMonday)
func (wd Weekdays) Set(weekdays Weekdays) {
	wd = weekdays
}

// Check given day exist in desire days!
func (wd Weekdays) Check(day Weekdays) (exist bool) {
	if day&wd == day {
		return true
	}
	return false
}

// CheckReverse given days exist in desire day!
func (wd Weekdays) CheckReverse(days Weekdays) (exist bool) {
	if wd&days == wd {
		return true
	}
	return false
}

// Check given day exist in desire Weekdays!
func (wd Weekdays) String() (day string) {
	switch lang.AppLanguage {
	case lang.EnglishLanguage:
		switch wd {
		case WeekdaysMonday:
			return "Monday"
		case WeekdaysTuesday:
			return "Tuesday"
		case WeekdaysWednesday:
			return "Wednesday"
		case WeekdaysThursday:
			return "Thursday"
		case WeekdaysFriday:
			return "Friday"
		case WeekdaysSaturday:
			return "Saturday"
		case WeekdaysSunday:
			return "Sunday"
		}
	case lang.PersianLanguage:
		switch wd {
		case WeekdaysMonday:
			return "دوشنبه"
		case WeekdaysTuesday:
			return "سه شنبه"
		case WeekdaysWednesday:
			return "چهارشنبه"
		case WeekdaysThursday:
			return "پنچ شنبه"
		case WeekdaysFriday:
			return "جمعه"
		case WeekdaysSaturday:
			return "شنبه"
		case WeekdaysSunday:
			return "یکشنبه"
		}
	}
	return
}

// Weekdays return Weekdays of given time.
// TODO::: not accurate yet!!??
func (t Time) Weekdays() (day Weekdays) {
	var secPassLastWeek = t % Week
	day = (1 << (secPassLastWeek / Day))
	return
}
