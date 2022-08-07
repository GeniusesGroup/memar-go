/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/scheduler"
)

// Sleep pauses the execution of the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(d protocol.Duration) {
	if d <= 0 {
		return
	}

	var thread = scheduler.ActiveThread()
	thread.Yield(scheduler.Thread_WaitReason_Sleep)

	var timer timer
	timer.init(thread)
	timer.start(d)
}
