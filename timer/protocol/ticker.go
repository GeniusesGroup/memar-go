/* For license and copyright information please see the LEGAL file in the code repository */

package timer_p

import (
	"memar/protocol"
	time_p "memar/time/protocol"
)

// Ticker is the interface that must implement by any ticker.
// Implement object of Timer can also be a ticker,
// just Start() method of timer is same as Tick(d, 0).
// Reset() just change the interval not first tick duration.
type Ticker[DUR time_p.Duration] interface {
	// Tick will add timer to default timing mechanism like TimingHeap, TimingWheel, ...
	Tick(first, interval DUR) (err protocol.Error)
}
