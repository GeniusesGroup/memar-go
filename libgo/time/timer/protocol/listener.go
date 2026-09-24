/* For license and copyright information please see the LEGAL file in the code repository */

package timer_p

// TimerListener or TimerCallBack
type TimerListener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	// Be aware that given callback must not be closure too. TODO::: Why??
	TimerHandler()
}
