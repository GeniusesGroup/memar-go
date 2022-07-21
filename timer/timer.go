/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"../protocol"
)

// A Timer must be created with Init, After or AfterFunc.
type Timer struct {
	timer
	signal chan struct{}
}

// Init initialize the Timer with given callback function or make the channel and send signal on it
// Be aware that given function must not be closure and must not block the caller.
func (t *Timer) Init(callback protocol.TimerListener) {
	if t.callback != nil {
		panic("timer: Don't initialize a timer twice. Use Reset() method to change the timer.")
	}

	if callback == nil {
		// Give the channel a 1-element buffer.
		// If the client falls behind while reading, we drop ticks
		// on the floor until the client catches up.
		t.signal = make(chan struct{}, 1)
		t.callback = t
	} else {
		t.callback = callback
	}
}
func (t *Timer) Signal() <-chan struct{}                           { return t.signal }
func (t *Timer) Start(d protocol.Duration)                         { t.add(d) }
func (t *Timer) Stop() (alreadyStopped bool)                       { return t.delete() }
func (t *Timer) Reset(d protocol.Duration) (alreadyActivated bool) { return t.modify(d) }

// TimerHandler or NotifyChannel does a non-blocking send the signal on t.signal
func (t *Timer) TimerHandler() {
	select {
	case t.signal <- struct{}{}:
	default:
	}
}

// After waits for the duration to elapse and then sends signal on the returned channel.
// The underlying Timer is not recovered by the garbage collector
// until the timer fires. If efficiency is a concern, copy the body
// instead and call timer.Stop() if the timer is no longer needed.
func After(d protocol.Duration) <-chan struct{} {
	var timer Timer
	timer.Init(nil)
	timer.Start(d)
	return timer.Signal()
}

// AfterFunc waits for the duration to elapse and then calls callback.
// If callback need blocking operation it must do its logic in new thread(goroutine).
// It returns a Timer that can be used to cancel the call using its Stop method.
func AfterFunc(d protocol.Duration, callback protocol.TimerListener) *Timer {
	var timer Timer
	timer.Init(callback)
	timer.Start(d)
	return &timer
}
