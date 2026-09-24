/* For license and copyright information please see the LEGAL file in the code repository */

package scheduler

import (
	"memar/process/race"
	"memar/time/duration"
	"memar/time/monotonic"
)

// Thread is user level threads and respect PersiaOS way not other OSs kernels.
// virtual processor is used instead of kernel thread.
type Thread struct {
	process            *Process
	allowChangeCPUCore bool

	wakeAt   monotonic.Time // sleepUntil
	priority uint8

	// Program counter
	// Registers
	// Stack

	status ThreadStatus
	// state chan ThreadStatus
	waitReason Thread_WaitReason

	raceCTX uintptr
}

func (t *Thread) Active() {
	t.status.Store(ThreadState_Ready)
	// TODO:::
}

func (t *Thread) Sleep(d duration.NanoSecond) {
	if d <= 0 {
		return
	}

	t.wakeAt.Now()
	t.wakeAt.Add(d)
	// TODO::: if timer is fire before we yield its thread??
	t.Yield(Thread_WaitReason_Sleep)
}

//memar:impl memar/time/timer/protocol.Timer
func (t *Thread) TimerHandler() { t.Active() }

// SetRaceCtx set the thread race context.
// panic in call in non race detection enabled.
func (t *Thread) SetRaceCtx(ctx uintptr) {
	if !race.DetectorEnabled {
		panic("scheduler - race: unexpected call to SetRaceCtx in non race detector condition")
	}
	if t.raceCTX != 0 {
		panic("scheduler - race: unexpected raceCTX")
	}
	t.raceCTX = ctx
}

// ReleaseRaceCtx set the thread race context to zero.
// panic in call in non race detection enabled.
func (t *Thread) ReleaseRaceCtx() {
	if !race.DetectorEnabled {
		panic("scheduler - race: unexpected call to ReleaseRaceCtx in non race detector condition")
	}
	t.raceCTX = 0
}
