/* For license and copyright information please see LEGAL file in repository */

package etime

// A Dayhours specifies a hour of a day.
type Dayhours uint32

// Hours
const (
	DayhoursNone  Dayhours = 0b00000000000000000000000000000000
	DayhoursOne   Dayhours = 0b00000000000000000000000000000001
	DayhoursTwo   Dayhours = 0b00000000000000000000000000000010
	DayhoursThree Dayhours = 0b00000000000000000000000000000100
	DayhoursAll   Dayhours = 0b11111111111111111111111111111111
)

// Set given hours to given Dayhours!
func (dh Dayhours) Set(hours Dayhours) {
	dh = hours
}

// Check given hour exist in given hours!
func (dh Dayhours) Check(checkHour Dayhours) (exist bool) {
	if checkHour&dh == checkHour {
		return true
	}
	return false
}

// CheckReverse check given hours exist in given hour!
func (dh Dayhours) CheckReverse(hours Dayhours) (exist bool) {
	if dh&hours == dh {
		return true
	}
	return false
}

// Dayhours return Hour of the day in the given time by Dayhours format.
func (t Time) Dayhours() (hour Dayhours) {
	var secPassLastDay = t % Day
	hour = (1 << (secPassLastDay / Hour))
	return
}
