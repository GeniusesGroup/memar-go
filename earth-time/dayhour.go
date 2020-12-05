/* For license and copyright information please see LEGAL file in repository */

package etime

// A Dayhour specifies a hour of a day in normal number.
type Dayhour uint8

// RoundToHour delete seconds parts of time to have unique time in each hour instead second
// Mostly use in index proccess!
func (t Time) RoundToHour() (rounded int64) {
	return int64(t) / Hour
}

// Dayhour return Hour of the day in the given time by normal number.
func (t Time) Dayhour() (hour Dayhour) {
	var secPassLastDay = t % Day
	return Dayhour(secPassLastDay / Hour)
}
