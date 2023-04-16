/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"libgo/cpu"
	"runtime"
)

var poolByCores = make(timingsHeap, cpu.CoreNum())

func init() {
	var coreNumbers = runtime.GOMAXPROCS(0)
	for id := 0; id < coreNumbers; id++ {
		poolByCores[id].Init()
	}
}

type timingsHeap []TimingHeap

func (tsh timingsHeap) activeTiming() *TimingHeap {
	return &tsh[cpu.ActiveCoreID()]
}

func getActiveTiming() *TimingHeap { return poolByCores.activeTiming() }
