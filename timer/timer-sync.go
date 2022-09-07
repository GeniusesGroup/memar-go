/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// After waits for the duration to elapse and then sends signal on the returned channel.
// The underlying Timer is not recovered by the garbage collector
// until the timer fires. If efficiency is a concern, copy the body
// instead and call timer.Stop() if the timer is no longer needed.
// It will **panic** if it can't start the timer due to any situation like not enough memory, ...
func After(d protocol.Duration) <-chan struct{} {
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
func NewSync(d protocol.Duration) (t *Sync, err protocol.Error) {
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

func (t *Sync) Reinit() { t.Stop(); close(t.signal); t.Async.Reinit() }
func (t *Sync) Deinit() { t.Stop(); close(t.signal) }

//libgo:impl protocol.Timer
func (t *Sync) Init() {
	// Give the channel a 1-element buffer.
	// If the client falls behind while reading, we drop ticks
	// on the floor until the client catches up.
	t.signal = make(chan struct{}, 1)
	t.Async.Init(t)
}
func (t *Sync) Signal() <-chan struct{}                           { return t.signal }
func (t *Sync) Start(d protocol.Duration) (err protocol.Error)    { return t.Async.Start(d) }
func (t *Sync) Stop() (alreadyStopped bool)                       { return t.Async.Stop() }
func (t *Sync) Reset(d protocol.Duration) (alreadyActivated bool) { return t.Modify(d) }

// TimerHandler or NotifyChannel does a non-blocking send the signal on t.signal
func (t *Sync) TimerHandler() {
	select {
	case t.signal <- struct{}{}:
	default:
	}
}
