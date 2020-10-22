/* For license and copyright information please see LEGAL file in repository */

package etime

import "time"

// Now returns earth time in second elapsed after ...!
func Now() (sec int64) {
	// TODO::: it is not so efficient
	return time.Now().Unix()
}

// NowNano returns earth time in nanosecond elapsed after ...!
func NowNano() (nano int64) {
	// TODO::: it is not so efficient
	return time.Now().UnixNano()
}
