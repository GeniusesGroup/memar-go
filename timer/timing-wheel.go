/* For license and copyright information please see LEGAL file in repository */

package timer

import (
	"../protocol"
)

type TimingWheel struct {
	wheel    [][]*Timer
	interval protocol.Duration // same as tw.ticker.period
	pos      int
	ticker   Timer
	stop     chan struct{}
}

// if you make one sec interval for a earth day(60*60*24=86400), you need 2MB ram just to hold empty wheel without any timer.
func (tw *TimingWheel) Init(interval protocol.Duration, wheelSize int) {
	tw.stop = make(chan struct{})
	tw.wheel = make([][]*Timer, wheelSize)
	tw.ticker.Init(nil, nil)
	tw.interval = interval
}

func (tw *TimingWheel) AddTimer(t *Timer) {
	var addedPosition = tw.addedPosition(t)
	if addedPosition == tw.pos {
		// TODO::: run the timer??
		return
	}
	tw.wheel[addedPosition] = append(tw.wheel[addedPosition], t)
}

// call by go keyword if you don't want the current goroutine block.
func (tw *TimingWheel) Start() {
	tw.ticker.Tick(tw.interval, tw.interval, -1)
LOOP:
	for {
		select {
		case <-tw.ticker.Signal():
			var pos = tw.pos
			tw.incrementTickPosition()
			var timers = tw.wheel[pos]
			tw.wheel[pos] = timers[:0]
			for i := 0; i < len(timers); i++ {
				var timer = timers[i]
				timer.callback(timer.arg)
				tw.checkAndAddTimerAgain(timer)
			}
		case <-tw.stop:
			tw.stop = nil
			break LOOP
		}
	}
	tw.ticker.Stop()
}

// Not concurrent safe.
func (tw *TimingWheel) Stop() (alreadyStopped bool) {
	if tw.stop == nil {
		return true
	}

	select {
	case tw.stop <- struct{}{}:
	default:
		alreadyStopped = true
	}
	tw.ticker.Stop()
	return
}

func (tw *TimingWheel) incrementTickPosition() {
	var wheelLen = len(tw.wheel)
	if wheelLen-1 == tw.pos {
		tw.pos = 0
	} else {
		tw.pos++
	}
}

func (tw *TimingWheel) checkAndAddTimerAgain(t *Timer) {
	if t.periodNumber == 0 {
		t.reset()
	} else {
		var addedPosition = tw.addedPosition(t)
		tw.wheel[addedPosition] = append(tw.wheel[addedPosition], t)
		if t.periodNumber > 0 {
			t.periodNumber--
		}
	}
}

func (tw *TimingWheel) addedPosition(t *Timer) int {
	return int(t.period/tw.interval) + tw.pos
}
