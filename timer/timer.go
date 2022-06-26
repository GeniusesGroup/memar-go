/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"../protocol"
)

// A Timer must be created with Init, After or AfterFunc.
type Timer struct {
	signal chan struct{}

	// callback function that call when reach
	// it is possible that callback will be called a little after the delay.
	callback func(arg any) // NOTE: must not be closure and must not block the caller.
	arg      any

	// Timer wakes up at when, and then at when+period, ... (period > 0 only)
	// each time calling callback(arg) in the timer goroutine, so callback must be
	// a well-behaved function and not block.
	// when must be positive on an active timer.
	when   int64
	period int64

	// The status field holds one of the values in status file.
	status status

	timers *Timers
}

// Init initialize the Timer with given callback function or make the channel and send signal on it
// Be aware that given function must not be closure and must not block the caller.
func (t *Timer) Init(callback func(arg any), arg any) {
	if t.callback != nil {
		panic("timer: Don't initialize a timer twice. Use Modify() method to change the timer.")
	}

	if callback == nil {
		// Give the channel a 1-element buffer.
		// If the client falls behind while reading, we drop ticks
		// on the floor until the client catches up.
		t.signal = make(chan struct{}, 1)
		t.callback = notifyTimerChannel
	} else {
		t.callback = callback
	}
}
func (t *Timer) Signal() <-chan struct{}                           { return t.signal }
func (t *Timer) Start(d protocol.Duration)                         { t.add(d) }
func (t *Timer) Stop() (alreadyStopped bool)                       { return t.delete() }
func (t *Timer) Reset(d protocol.Duration) (alreadyActivated bool) { return t.modify(d) }

// After waits for the duration to elapse and then sends signal on the returned channel.
// The underlying Timer is not recovered by the garbage collector
// until the timer fires. If efficiency is a concern, copy the body
// instead and call timer.Stop() if the timer is no longer needed.
func After(d protocol.Duration) <-chan struct{} {
	var timer Timer
	timer.Init(nil, nil)
	timer.Start(d)
	return timer.Signal()
}

// AfterFunc waits for the duration to elapse and then calls f
// in its own goroutine. It returns a Timer that can
// be used to cancel the call using its Stop method.
func AfterFunc(d protocol.Duration, cb func(arg any), arg any) *Timer {
	var timer Timer
	timer.Init(callback(cb).concurrentRun, arg)
	timer.Start(d)
	return &timer
}
