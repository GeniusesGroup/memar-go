/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Timing observe Timers or Tickers and call TimerHandler() methods of them in desire time.
// All package provide default Timing mechanism for easily usage,
// But they should provide some other algorithms for other use-cases too.
// Packages can also break Init() methods of Timer or Ticker if they can't provide default Timing mechanism e.g. on TimingWheel
type Timing interface {
	// Depend on implementation but in most cases t can be a Ticker too.
	AddTimer(t Timer) (err Error)
}

// Timer is the interface that must implement by any timer.
//
// many other type than async timer can be implemented by libraries,
// like channel-based sync one that provide e.g. Signal() <-chan struct{}
// client use Signal() to block until timeout occur.
type Timer interface {
	// Init initialize the Timer with given callback function.
	// - **NOTE**: each time calling callback() in the timer goroutine, so callback must be
	// a well-behaved function and not block. If callback need blocking operation it must do its logic in new thread(goroutine).
	// - Be aware that given function must not be closure and must not block the caller.
	Init(callback TimerListener) (err Error)

	// TODO::: instead of Init we must force ObjectLifeCycle methods but
	// we have problem with optional args for method interface that Go not support like ECMA-script
	// ObjectLifeCycle

	// Start will add timer to default timing mechanism like TimingHeap, TimingWheel, ...
	Start(d Duration) (err Error)

	Reset(d Duration) (err Error)

	// Client must call Stop(), otherwise **"leaks"** occur, specially in Tick()
	Stop() (err Error)

	// Status return active status of the timer.
	// It is atomic operation and return a state at a particular time and
	// can be changed just after you get the status.
	Status() (status TimerStatus)

	Stringer
}

// Ticker is the interface that must implement by any ticker.
// Implement object of Timer can also be a ticker,
// just Start() method of timer is same as Tick(d, 0).
// Reset() just change the interval not first tick duration.
type Ticker interface {
	// Tick will add timer to default timing mechanism like TimingHeap, TimingWheel, ...
	Tick(first, interval Duration) (err Error)
}

// TimerListener or TimerCallBack
type TimerListener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	TimerHandler()
}
