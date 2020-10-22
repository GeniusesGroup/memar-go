/* For license and copyright information please see LEGAL file in repository */

package etime

// RoundSeconds delete seconds parts of time to have unique time in each second period instead second
func RoundSeconds(time int64, period int64) (rounded int64) {
	return time / period
}

// RoundToHour delete seconds parts of time to have unique time in each hour instead second
// Mostly use in index proccess!
func RoundToHour(time int64) (rounded int64) {
	return time / (60 * 60)
}

// RoundToDay delete seconds parts of time to have unique time in each day instead second
// Mostly use in index proccess!
func RoundToDay(time int64) (rounded int64) {
	return time / (24 * 60 * 60)
}

// RoundToMonth delete seconds parts of time to have unique time in each month instead second
// Mostly use in index proccess!
func RoundToMonth(time int64) (rounded int64) {
	return time / (30 * 24 * 60 * 60)
}

// RoundToYear delete seconds parts of time to have unique time in each year instead second
// Mostly use in index proccess!
func RoundToYear(time int64) (rounded int64) {
	return time / (365 * 24 * 60 * 60)
}
