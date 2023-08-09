/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"memar/protocol"
	"memar/scheduler"
)

// Sleep pauses the execution of the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(d protocol.Duration) {
	if d <= 0 {
		return
	}

	var thread = scheduler.ActiveThread()
	var timer Async
	timer.Init(thread)
	timer.Start(d)

	// TODO::: if timer is fire before we yield its thread??
	thread.Yield(scheduler.Thread_WaitReason_Sleep)
}
