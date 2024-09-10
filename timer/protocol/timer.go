/* For license and copyright information please see the LEGAL file in the code repository */

package timer_p

import (
	error_p "memar/error/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
)

// Timer is the interface that must implement by any timer.
//
// many other type than async timer can be implemented by libraries,
// like channel-based sync one that provide e.g. Signal() <-chan struct{}
// client use Signal() to block until timeout occur.
type Timer[ /*DUR time_p.Duration,*/ TIME time_p.Time, ST TimerStatus] interface {
	// Init initialize the Timer with given callback function.
	// - **NOTE**: each time calling callback() in the timer goroutine, so callback must be
	// a well-behaved function and not block. If callback need blocking operation it must do its logic in new thread(goroutine).
	// - Be aware that given function must not be closure and must not block the caller.
	Init(callback TimerListener) (err error_p.Error)

	// TODO::: instead of Init we must force object_p.LifeCycle methods but
	// we have problem with optional args for method interface that Go not support like ECMA-script
	// object_p.LifeCycle

	// Start will add timer to default timing mechanism like TimingHeap, TimingWheel, ...
	Start(d duration.NanoSecond) (err error_p.Error)

	Reset(d duration.NanoSecond) (err error_p.Error)

	// Client must call Stop(), otherwise **"leaks"** occur, specially in Tick()
	Stop() (err error_p.Error)

	When() TIME

	// Status return active status of the timer.
	// It is atomic operation and return a state at a particular time and
	// can be changed just after you get the status.
	Status() ST

	// Stringer[String]
}
