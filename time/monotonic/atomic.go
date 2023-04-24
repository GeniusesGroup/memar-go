/* For license and copyright information please see the LEGAL file in the code repository */

package monotonic

import (
	"sync/atomic"

	"libgo/protocol"
)

// Atomic same as Time is monotonic clock is for measuring time.
// Just due to 32bit hardwares alignment problem,
// we suggest don't use Time for atomic operation and Use this type for any atomic purposes.
type Atomic struct {
	atomic.Int64
}

func (a *Atomic) Load() Time               { return Time(a.Int64.Load()) }
func (a *Atomic) Store(t Time)             { a.Int64.Store(int64(t)) }
func (a *Atomic) Swap(new Time) (old Time) { return Time(a.Int64.Swap(int64(new))) }
func (a *Atomic) CompareAndSwap(old, new Time) (swapped bool) {
	return a.Int64.CompareAndSwap(int64(old), int64(new))
}
func (a *Atomic) Add(d protocol.Duration) { a.Int64.Add(int64(d)) }

//libgo:impl /libgo/protocol.Time
func (a *Atomic) Epoch() protocol.TimeEpoch { return protocol.TimeEpoch_Monotonic }
func (a *Atomic) SecondElapsed() int64      { return int64(a.Load()) / int64(Second) }
func (a *Atomic) NanoSecondElapsed() int32  { var t = a.Load(); return int32(t % (t / Time(Second))) }

//libgo:impl /libgo/protocol.Stringer
func (a *Atomic) ToString() string {
	var t = a.Load()
	return t.ToString()
}
func (a *Atomic) FromString(s string) (err protocol.Error) {
	var t Time
	err = t.FromString(s)
	a.Store(t)
	return
}
