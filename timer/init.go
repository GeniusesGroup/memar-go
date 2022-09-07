/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"runtime"
)

var poolByCores = make([]TimingHeap, runtime.NumCPU())

func init() {
	// var coreNumbers = runtime.GOMAXPROCS(0)
	// TODO:::
}
