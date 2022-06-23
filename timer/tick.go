/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"../protocol"
)

// Tick will send a signal on the t.Signal() channel after each tick on initialized Timer.
// The period of the ticks is specified by the duration argument.
// The ticker will adjust the time interval or drop ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, Tick() will panic.
// Stop the ticker to release associated resources.
func (t *Timer) Tick(d protocol.Duration) {
	if d <= 0 {
		panic("non-positive interval to tick")
	}
	t.period = int64(d)
	t.Start(d)
}

// Tick is a convenience wrapper for Timer.Tick() providing access to the ticking
// channel only. While Tick is useful for clients that have no need to shut down
// the Ticker, be aware that without a way to shut it down the underlying
// Ticker cannot be recovered by the garbage collector; it "leaks".
func Tick(d protocol.Duration) <-chan struct{} {
	var timer Timer
	timer.Init(nil)
	timer.Tick(d)
	return timer.Signal()
}

// TickFunc waits for the duration to elapse and then calls f in each duration elapsed
// in its own goroutine. It returns a Timer that can
// be used to cancel the call using its Stop method.
func TickFunc(d protocol.Duration, f func()) *Timer {
	var timer Timer
	timer.Init(goFunc(f).concurrentRun)
	timer.Tick(d)
	return &timer
}
