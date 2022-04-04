/* For license and copyright information please see LEGAL file in repository */

package monotonic

import (
	"sync/atomic"

	"../../protocol"
)

type Time struct {
	sec  int64
	nsec int32
}

func (t *Time) Epoch() protocol.TimeEpoch { return protocol.TimeEpoch_Monotonic }
func (t *Time) SecondElapsed() int64      { return t.sec }
func (t *Time) NanoSecondElapsed() int32  { return t.nsec }
func (t *Time) ToString() string {
	// TODO:::
	return ""
}

func (t *Time) Now() {
	var nsec = RuntimeNano()
	var secElapsed = nsec / int64(Second)
	t.sec = secElapsed
	t.nsec = int32(nsec % (secElapsed * int64(Second)))
}
func (t *Time) NowAtomic() {
	var nsec = RuntimeNano()
	var secElapsed = nsec / int64(Second)
	var nsecElapsed = nsec % (secElapsed * int64(Second))
	atomic.AddInt64(&t.sec, secElapsed)
	atomic.AddInt32(&t.nsec, int32(nsecElapsed))
}

func (t *Time) Pass(baseTime Time) (pass bool)             { return }
func (t *Time) AddDuration(d protocol.Duration) (new Time) { return }
