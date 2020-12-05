/* For license and copyright information please see LEGAL file in repository */

package etime

import (
	lang "../language"
)

// A Weekday specifies a day of the week.
type Weekday uint8

// Weekdays
const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Check given day exist in desire Weekdays!
func (wd Weekday) String() (day string) {
	switch lang.AppLanguage {
	case lang.EnglishLanguage:
		switch wd {
		case Monday:
			return "Monday"
		case Tuesday:
			return "Tuesday"
		case Wednesday:
			return "Wednesday"
		case Thursday:
			return "Thursday"
		case Friday:
			return "Friday"
		case Saturday:
			return "Saturday"
		case Sunday:
			return "Sunday"
		}
	case lang.PersianLanguage:
		switch wd {
		case Monday:
			return "دوشنبه"
		case Tuesday:
			return "سه شنبه"
		case Wednesday:
			return "چهارشنبه"
		case Thursday:
			return "پنچ شنبه"
		case Friday:
			return "جمعه"
		case Saturday:
			return "شنبه"
		case Sunday:
			return "یکشنبه"
		}
	}
	return
}

// Weekday return Weekday of given time.
// TODO::: not accurate yet!!??
func (t Time) Weekday() (day Weekday) {
	var secPassLastWeek = t % Week
	switch secPassLastWeek / Day {
	case 0:
		return Thursday
	case 1:
		return Friday
	case 2:
		return Saturday
	case 3:
		return Sunday
	case 4:
		return Monday
	case 5:
		return Tuesday
	case 6:
		return Wednesday
	}
	return
}
