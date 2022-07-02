//go:build !time-wheel

/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"runtime"
	"sync/atomic"
)

var poolByCores = make([]TimingHeap, runtime.NumCPU())

func init() {
	// var coreNumbers = runtime.GOMAXPROCS(0)
	// TODO:::
}

// destroy releases all of the resources associated with timers in specific CPU core and
// move them to other core
func (th *TimingHeap) destroy() {
	if len(th.timers) > 0 {
		th.timersLock.Lock()
		th.moveTimers(plocal, th.timers)
		th.timers = nil
		th.numTimers = 0
		th.deletedTimers = 0
		atomic.StoreInt64(&th.timer0When, 0)
		th.timersLock.Unlock()
	}
}
