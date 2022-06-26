/* For license and copyright information please see LEGAL file in repository */

package timer

type callback func(arg any)

func (c callback) concurrentRun(arg any) {
	go c(arg)
}

// NotifyChannel does a non-blocking send the signal on t.signal
func notifyTimerChannel(t any) {
	select {
	case t.(*Timer).signal <- struct{}{}:
	default:
	}
}
