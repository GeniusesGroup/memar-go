/* For license and copyright information please see LEGAL file in repository */

package etime

// RoundTo delete seconds parts of time to have unique time in each second period instead second
func (t Time) RoundTo(second Duration) (rounded int64) {
	return int64(t) / int64(second)
}

// UntilRoundTo return second duration until round of given time in given period!
func (t Time) UntilRoundTo(period Duration) (duration Duration) {
	return period - (Duration(t) - (Duration(t)/period)*period)
}

// UntilTo return second duration until to given time!
func (t Time) UntilTo(time Time) (duration Duration) {
	return Duration(time - t)
}
