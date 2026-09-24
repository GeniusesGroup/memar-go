/* For license and copyright information please see the LEGAL file in the code repository */

package scheduler

import (
	"runtime"
)

// https://zhdonas.medium.com/how-to-train-your-goroutine-e241bf0c2d3
func init() {
	runtime.LockOSThread()
}
