/* For license and copyright information please see LEGAL file in repository */

package protocol

type Timing interface {
	AddTimer(t Timer)
}

// Timer is the interface that must implement by any timer.
type Timer interface {
	// if Init() called by nil for callback, client can use Signal() to block until timeout occur.
	// **NOTE**: each time calling callback() in the timer goroutine, so callback must be
	// a well-behaved function and not block.
	// If callback need blocking operation it must do its logic in new thread(goroutine).
	Init(callback TimerListener)

	Start(d Duration)
	Tick(first, interval Duration, periodNumber int64)
	// if timer is a ticker, reset just change the interval not first tick duration.
	Reset(d Duration) (alreadyActivated bool)

	// Client must call Stop(), otherwise **"leaks"** occur, specially in Tick()
	Stop() (alreadyStopped bool)

	// Suggest to implement not force here to get the timer channel to receive timer-out signal
	// Signal() <-chan struct{}

	Stringer
}

// TimerListener or TimerCallBack
type TimerListener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	TimerHandler()
}
