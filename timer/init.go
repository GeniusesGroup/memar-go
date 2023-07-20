/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"libgo/cpu"
	"runtime"
)

var poolByCores = make(timing, cpu.CoreNum())

func init() {
	var coreNumbers = runtime.GOMAXPROCS(0)
	for id := 0; id < coreNumbers; id++ {
		poolByCores[id].Init()
	}
}

func getActiveTiming() *Timing { return poolByCores.activeTiming() }

type timing []Timing

func (tg timing) activeTiming() *Timing {
	return &tg[cpu.ActiveCoreID()]
}
