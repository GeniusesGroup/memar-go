/* For license and copyright information please see the LEGAL file in the code repository */

package scheduler

import (
	"memar/time/duration"
)

func ActiveThread() *Thread {
	return nil
}

// Yield call active thread Yield()
func Yield(wr Thread_WaitReason) {
	var thread = ActiveThread()
	thread.Yield(wr)
}

// Sleep pauses the execution of the current user-space thread for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(d duration.NanoSecond) {
	if d <= 0 {
		return
	}

	var thread = ActiveThread()
	thread.Sleep(d)
}

// SetRaceCtx set active thread race context.
func SetRaceCtx(ctx uintptr) {
	var thread = ActiveThread()
	thread.SetRaceCtx(ctx)
}

// SetRaceCtx release active thread race context.
func ReleaseRaceCtx() {
	var thread = ActiveThread()
	thread.ReleaseRaceCtx()
}
