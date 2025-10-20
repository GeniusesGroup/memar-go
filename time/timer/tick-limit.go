/* For license and copyright information please see the LEGAL file in the code repository */

package timer

import (
	error_p "memar/process/error/protocol"
	"memar/time/duration"
	errs "memar/time/timer/errors"
)

func NewLimitTicker(first, interval duration.NanoSecond, periodNumber int64) (t *LimitTicker, err error_p.Error) {
	if periodNumber < 1 {
		err = &errs.ErrNegativePeriodNumber
		return
	}

	var timer LimitTicker
	err = timer.Init()
	if err != nil {
		return
	}
	timer.periodNumber = periodNumber
	err = timer.Tick(first, interval)
	t = &timer
	return
}

type LimitTicker struct {
	periodNumber int64 // -1 means no limit
	Sync
}

//memar:impl memar/time/timer/protocol.Timer
func (self *LimitTicker) Init() (err error_p.Error) {
	// Give the channel a 1-element buffer.
	// If the client falls behind while reading, we drop ticks
	// on the floor until the client catches up.
	self.signal = make(chan struct{}, 1)
	err = self.Async.Init(self)
	return
}

func (self *LimitTicker) RemainingNumber() int64 { return self.periodNumber }

// TimerHandler or NotifyChannel does a non-blocking send the signal on t.signal
func (self *LimitTicker) TimerHandler() {
	select {
	case self.signal <- struct{}{}:
	default:
	}

	if self.periodNumber > 0 {
		self.periodNumber--
	} else {
		self.Stop()
	}
}
