/* For license and copyright information please see the LEGAL file in the code repository */

package unix

import (
	"sync/atomic"

	"memar/protocol"
)

// Atomic same as Time is unix clock is for measuring time.
// Just appear due to 32bit hardwares alignment problem,
// we suggest don't use Time for atomic operation and Use this type for any atomic purposes.
type Atomic struct {
	sec  atomic.Int64
	nsec atomic.Int32
}

//memar:impl memar/protocol.Time
func (t *Atomic) Epoch() protocol.TimeEpoch { return protocol.TimeEpoch_Unix }
func (t *Atomic) SecondElapsed() int64      { return t.sec.Load() }
func (t *Atomic) NanoSecondElapsed() int32  { return t.nsec.Load() }

func (a *Atomic) Now() {
	var t Time
	t.Now()
	a.Store(t)
}

func (a *Atomic) Load() Time { return Time{a.sec.Load(), a.nsec.Load()} }

// TODO::: below methods not work logically and they have problem. use mutex?

func (a *Atomic) Store(t Time) {
	var oldNSec = a.nsec.Load()
	a.sec.Store(t.sec)
	a.nsec.CompareAndSwap(oldNSec, t.nsec)
}
func (a *Atomic) Swap(new Time) (old Time) {
	old = Time{a.sec.Swap(new.sec), a.nsec.Swap(new.nsec)}
	return
}
func (a *Atomic) CompareAndSwap(old, new Time) (swapped bool) {
	swapped = a.sec.CompareAndSwap(old.sec, new.sec)
	if !swapped {
		return
	}
	swapped = a.nsec.CompareAndSwap(old.nsec, new.nsec)
	return
}
func (a *Atomic) Add(d protocol.Duration) {
	var sec, nsec = nsecToSec(d)
	a.sec.Add(sec)
	a.nsec.Add(nsec)
}

//memar:impl memar/protocol.Stringer
func (a *Atomic) ToString() string {
	// TODO:::
	return ""
}
func (a *Atomic) FromString(s string) (err protocol.Error) {
	// TODO:::
	return
}
