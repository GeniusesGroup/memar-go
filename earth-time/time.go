/* For license and copyright information please see LEGAL file in repository */

package etime

import "time"

// A Time specifies second elapsed of January 1 of the absolute year.
// January 1 of the absolute year(1970), like January 1 of 2001, was a Monday.
type Time int64

// Now returns earth time in second elapsed after ...!
func Now() (sec Time) {
	// TODO::: it is not so efficient
	return Time(time.Now().Unix())
}

// Pass check if time pass from given time
func (t Time) Pass(baseTime Time) (pass bool) {
	if baseTime > t {
		return true
	}
	return false
}

// AddDuration return given time plus given duration
func (t Time) AddDuration(d Duration) (new Time) {
	return t + Time(d)
}

// Local change given time to local time by OS set time zone
func (t Time) Local() (loc Time) {
	// TODO:::
	return t
}

// RoundToDay delete seconds parts of time to have unique time in each day instead second
// Mostly use in index proccess!
func (t Time) RoundToDay() (rounded int64) {
	return int64(t) / (24 * 60 * 60)
}

// RoundToMonth delete seconds parts of time to have unique time in each month instead second
// Mostly use in index proccess!
func (t Time) RoundToMonth() (rounded int64) {
	return int64(t) / (30 * 24 * 60 * 60)
}

// RoundToYear delete seconds parts of time to have unique time in each year instead second
// Mostly use in index proccess!
func (t Time) RoundToYear() (rounded int64) {
	return int64(t) / (365 * 24 * 60 * 60)
}
