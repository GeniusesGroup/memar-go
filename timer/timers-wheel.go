//go:build time-wheel

/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"runtime"
)

var poolByCores = make([]TimingWheel, runtime.NumCPU())

func init() {
	// var coreNumbers = runtime.GOMAXPROCS(0)
	// TODO:::
}
