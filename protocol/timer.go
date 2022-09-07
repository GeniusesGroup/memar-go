/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Timing observe Timers and
type Timing interface {
	AddTimer(t Timer)
}

// Timer is the interface that must implement by any timer.
type Timer interface {
	Start(d Duration) (err Error)

	Reset(d Duration) (alreadyActivated bool)

	// Client must call Stop(), otherwise **"leaks"** occur, specially in Tick()
	Stop() (alreadyStopped bool)

	Stringer
}

// Timer_Async is the interface that must implement by any timer.
type Timer_Async interface {
	// Init initialize the Timer with given callback function.
	// - **NOTE**: each time calling callback() in the timer goroutine, so callback must be
	// a well-behaved function and not block. If callback need blocking operation it must do its logic in new thread(goroutine).
	// - Be aware that given function must not be closure and must not block the caller.
	Init(t Timing, callback TimerListener)

	Timer
}

type Timer_Sync interface {
	// Init initialize the Timer and make ready the timer channel and send signal on it
	// client use Signal() to block until timeout occur.
	Init(t Timing)

	// Suggest to implement not force here to get the timer channel to receive timer-out signal
	Signal() <-chan struct{}

	Timer
}

// Ticker is the interface that must implement by any ticker.
// Both Timer_Async and Timer_Sync can also be a ticker, just
// Start method of timer is same as Tick(d, 0)
// Reset just change the interval not first tick duration.
type Ticker interface {
	Tick(first, interval Duration) (err Error)
}

// TimerListener or TimerCallBack
type TimerListener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	TimerHandler()
}
