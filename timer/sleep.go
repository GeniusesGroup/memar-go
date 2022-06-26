/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"unsafe"

	"../protocol"
	"../scheduler"
)

// Sleep pauses the execution of the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(d protocol.Duration) {
	if d <= 0 {
		return
	}

	var thread = scheduler.ActiveThread()
	var timer Timer
	timer.Init(goroutineReady, thread)
	timer.Start(d)

	// TODO::: Decide to park or sleep??
	// gopark(t.resetForSleep, unsafe.Pointer(t), waitReasonSleep, traceEvGoSleep, 1)
	thread.Sleep(scheduler.ThreadWaitReason_Sleep)
}

// resetForSleep is called after the goroutine is parked for timeSleep.
// We can't call resettimer in timeSleep itself because if this is a short
// sleep and there are many goroutines then the P can wind up running the
// timer function, goroutineReady, before the goroutine has been parked.
func resetForSleep(gp *g, ut unsafe.Pointer) bool {
	var t = (*Timer)(ut)
	resettimer(t, t.when)
	return true
}

// Ready the goroutine arg.
func goroutineReady(arg any) {
	var thread = arg.(*scheduler.Thread)
	thread.Ready(0)
}
