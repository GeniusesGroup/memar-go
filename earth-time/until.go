/* For license and copyright information please see LEGAL file in repository */

package etime

// UntilRoundSeconds return second until round of given time in given period!
func UntilRoundSeconds(time int64, period int64) (duration int64) {
	return period - (time - (time/period)*period)
}
