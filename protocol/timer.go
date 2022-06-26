/* For license and copyright information please see LEGAL file in repository */

package protocol

// Timer is the interface that must implement by any timer.
type Timer interface {
	// if Init() called by nil for callback, client can use Signal() to block until timeout occur.
	Init(callback func(arg any), arg any)

	Start(d Duration)
	Tick(first, interval Duration)
	// if timer is a ticker, reset just change the interval not first tick duration.
	Reset(d Duration) (alreadyActivated bool)

	// Client must call Stop(), otherwise **"leaks"** occur
	Stop() (alreadyStopped bool)

	Signal() <-chan struct{}

	Stringer
}
