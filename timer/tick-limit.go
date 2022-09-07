/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

func NewLimitTicker(first, interval protocol.Duration, periodNumber int64) (t *LimitTicker, err protocol.Error) {
	if periodNumber < 1 {
		panic("timer - LimitTicker: periodNumber must be more than one.")
	}

	var timer LimitTicker
	timer.Init()
	timer.periodNumber = periodNumber
	err = timer.Tick(first, interval)
	t = &timer
	return
}

type LimitTicker struct {
	periodNumber int64 // -1 means no limit
	Sync
}

//libgo:impl protocol.Timer
func (t *LimitTicker) Init() {
	// Give the channel a 1-element buffer.
	// If the client falls behind while reading, we drop ticks
	// on the floor until the client catches up.
	t.signal = make(chan struct{}, 1)
	t.Async.Init(t)
}

func (t *LimitTicker) RemainingNumber() int64 { return t.periodNumber }

// TimerHandler or NotifyChannel does a non-blocking send the signal on t.signal
func (t *LimitTicker) TimerHandler() {
	select {
	case t.signal <- struct{}{}:
	default:
	}

	if t.periodNumber > 0 {
		t.periodNumber--
	} else {
		t.Stop()
	}
}
