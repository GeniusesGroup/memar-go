/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"memar/protocol"
	"memar/time/duration"
	t_p "memar/timer/protocol"
)

// NewSyncTick is a convenience wrapper for SyncTimer.Tick() providing access to the ticking.
// Unlike After() that providing access to (<-chan struct{}),
// due to client need a way to Deinit the underlying
// Ticker to recovered by the garbage collector; to prevent **"leaks"**.
func NewSyncTick(first, interval duration.NanoSecond) (t *Sync, err protocol.Error) {
	var timer Sync
	timer.Init()
	err = timer.Tick(first, interval)
	t = &timer
	return
}

// NewAsyncTick or Schedule waits for the duration to elapse and then calls callback in each duration elapsed
// If callback need blocking operation it must do its logic in new thread(goroutine).
// It returns a Timer that can be used to cancel the call using its Stop method.
// Schedule an execution at a given time, then once per interval. A typical use case is to execute code once per day at 12am.
func NewAsyncTick(first, interval duration.NanoSecond, callback t_p.TimerListener) (t *Async, err protocol.Error) {
	var timer Async
	timer.Init(callback)
	err = timer.Tick(first, interval)
	t = &timer
	return
}
