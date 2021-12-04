/* For license and copyright information please see LEGAL file in repository */

package log

import "runtime"

func CallerInfo(calldepth int) (file string, line int) {
	var ok bool
	_, file, line, ok = runtime.Caller(calldepth)
	if !ok {
		file = "?"
		line = 0
	}
	return
}
