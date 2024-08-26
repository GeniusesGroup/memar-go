/* For license and copyright information please see the LEGAL file in the code repository */

package timer_p

import (
	"memar/protocol"
	time_p "memar/time/protocol"
)

// Timing observe Timers or Tickers and call TimerHandler() methods of them in desire time.
// All package provide default Timing mechanism for easily usage,
// But they should provide some other algorithms for other use-cases too.
// Packages can also break Init() methods of Timer or Ticker if they can't provide default Timing mechanism e.g. on TimingWheel
type Timing[DUR time_p.Duration, TIME time_p.Time, ST TimerStatus] interface {
	// Depend on implementation but in most cases t can be a Ticker too.
	AddTimer(t Timer[DUR, TIME, ST]) (err protocol.Error)
}
