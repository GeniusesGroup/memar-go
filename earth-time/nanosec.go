/* For license and copyright information please see LEGAL file in repository */

package etime

import "time"

// A NanoSec specifies time elapsed of .
type NanoSec int64

// NowNano returns earth time in nanosecond elapsed after ...!
func NowNano() (nano NanoSec) {
	// TODO::: it is not so efficient
	return NanoSec(time.Now().UnixNano())
}
