/* For license and copyright information please see the LEGAL file in the code repository */

package unix

import (
	"sync/atomic"

	"memar/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
)

// Atomic same as Time is unix clock is for measuring time.
// Just appear due to 32bit hardwares alignment problem,
// we suggest don't use Time for atomic operation and Use this type for any atomic purposes.
type Atomic struct {
	sec  atomic.Int64
	nsec atomic.Int32
}

//memar:impl memar/time/protocol.Time
func (t *Atomic) Epoch() time_p.Epoch { return &Epoch }
func (t *Atomic) SecondElapsed() duration.Second {
	return duration.Second(t.sec.Load())
}
func (t *Atomic) NanoInSecondElapsed() duration.NanoInSecond {
	return duration.NanoInSecond(t.nsec.Load())
}

func (a *Atomic) Now() {
	var t Time
	t.Now()
	a.Store(t)
}

func (a *Atomic) Load() Time {
	return Time{duration.Second(a.sec.Load()), duration.NanoInSecond(a.nsec.Load())}
}

// TODO::: below methods not work logically and they have problem. use mutex?

func (a *Atomic) Store(t Time) {
	var oldNSec = a.nsec.Load()
	a.sec.Store(int64(t.sec))
	a.nsec.CompareAndSwap(oldNSec, int32(t.nsec))
}
func (a *Atomic) Swap(new Time) (old Time) {
	old = Time{
		sec:  duration.Second(a.sec.Swap(int64(new.sec))),
		nsec: duration.NanoInSecond(a.nsec.Swap(int32(new.nsec))),
	}
	return
}
func (a *Atomic) CompareAndSwap(old, new Time) (swapped bool) {
	swapped = a.sec.CompareAndSwap(int64(old.sec), int64(new.sec))
	if !swapped {
		return
	}
	swapped = a.nsec.CompareAndSwap(int32(old.nsec), int32(new.nsec))
	return
}
func (a *Atomic) Add(d duration.NanoSecond) {
	var sec, nsec = d.ToSecAndNano()
	a.sec.Add(int64(sec))
	a.nsec.Add(int32(nsec))
}

//memar:impl memar/protocol.Stringer
func (a *Atomic) ToString() (str string, err protocol.Error) {
	// TODO:::
	return
}
func (a *Atomic) FromString(str string) (err protocol.Error) {
	// TODO:::
	return
}
