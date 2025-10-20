/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	error_p "memar/process/error/protocol"
	"memar/time/duration"
)

// After waits for the duration to elapse and then sends signal on the returned channel.
// The underlying Timer is not recovered by the garbage collector until the timer fires.
// If efficiency is a concern, copy the body instead and call timer.Stop() if the timer is no longer needed.
// It will **panic** if it can't start the timer due to any situation like not enough memory, ...
func After(d duration.NanoSecond) <-chan struct{} {
	var timer Sync
	timer.Init()
	var err = timer.Start(d)
	if err != nil {
		panic(err)
	}
	return timer.Signal()
}

// NewAsync waits for the duration to elapse and then calls callback.
// If callback need blocking operation it must do its logic in new thread(goroutine).
// It returns a SyncTimer that can be used to cancel the call using its Stop method.
func NewSync(d duration.NanoSecond) (t *Sync, err error_p.Error) {
	var timer Sync
	timer.Init()
	err = timer.Start(d)
	t = &timer
	return
}

// Sync Timer must be created with Init, After or AfterFunc.
type Sync struct {
	Async
	signal chan struct{}
}

// Init initialize the timer by make the channel and send signal on it
//
//memar:impl memar/time/timer/protocol.Timer
func (self *Sync) Init() (err error_p.Error) {
	// Give the channel a 1-element buffer.
	// If the client falls behind while reading, we drop ticks
	// on the floor until the client catches up.
	self.signal = make(chan struct{}, 1)
	err = self.Async.Init(self)
	return
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Sync) Reinit() (err error_p.Error) { err = self.Async.Reinit(self); return }
func (self *Sync) Deinit() (err error_p.Error) {
	err = self.Async.Deinit()
	if err != nil {
		return
	}
	close(self.signal)
	return
}

//memar:impl memar/time/timer/protocol.Timer_Sync
func (self *Sync) Signal() <-chan struct{} { return self.signal }

// TimerHandler or NotifyChannel does a non-blocking send the signal on self.signal
func (self *Sync) TimerHandler() {
	select {
	case self.signal <- struct{}{}:
	default:
	}
}
