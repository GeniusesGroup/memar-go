/* For license and copyright information please see the LEGAL file in the code repository */

package utc

import (
	"memar/protocol"
)

// A Weekdays specifies a day of the week.
// can use by multiple state e.g. Weekdays_Saturday|Weekdays_Monday
type Weekdays uint8

// Weekdays
const (
	Weekdays_None      Weekdays = 0b00000000
	Weekdays_Monday    Weekdays = 0b00000001
	Weekdays_Tuesday   Weekdays = 0b00000010
	Weekdays_Wednesday Weekdays = 0b00000100
	Weekdays_Thursday  Weekdays = 0b00001000
	Weekdays_Friday    Weekdays = 0b00010000
	Weekdays_Saturday  Weekdays = 0b00100000
	Weekdays_Sunday    Weekdays = 0b01000000
	Weekdays_All       Weekdays = 0b11111111
)

// Check given day exist in desire days
func (w Weekdays) Check(day Weekdays) (exist bool) { return day&w != 0 }

// Check given day exist in desire Weekdays!
func (wd Weekdays) String() (day string) {
	switch protocol.AppLanguage {
	case protocol.LanguageEnglish:
		switch wd {
		case Weekdays_Monday:
			return "Monday"
		case Weekdays_Tuesday:
			return "Tuesday"
		case Weekdays_Wednesday:
			return "Wednesday"
		case Weekdays_Thursday:
			return "Thursday"
		case Weekdays_Friday:
			return "Friday"
		case Weekdays_Saturday:
			return "Saturday"
		case Weekdays_Sunday:
			return "Sunday"
		}
	case protocol.LanguagePersian:
		switch wd {
		case Weekdays_Monday:
			return "دوشنبه"
		case Weekdays_Tuesday:
			return "سه شنبه"
		case Weekdays_Wednesday:
			return "چهارشنبه"
		case Weekdays_Thursday:
			return "پنچ شنبه"
		case Weekdays_Friday:
			return "جمعه"
		case Weekdays_Saturday:
			return "شنبه"
		case Weekdays_Sunday:
			return "یکشنبه"
		}
	}
	return
}
