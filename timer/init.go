/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"runtime"
)

var poolByCores = make([]TimingHeap, runtime.NumCPU())

func init() {
	// var coreNumbers = runtime.GOMAXPROCS(0)
	// TODO:::
}
