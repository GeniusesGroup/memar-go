/* For license and copyright information please see the LEGAL file in the code repository */

package monotonic

import (
	"sync/atomic"

	error_p "memar/process/error/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
)

// Atomic same as Time is monotonic clock is for measuring time.
// Just due to 32bit hardwares alignment problem,
// we suggest don't use Time for atomic operation and Use this type for any atomic purposes.
type Atomic struct {
	dur atomic.Int64
}

func (self *Atomic) Load() Time               { return Time(self.dur.Load()) }
func (self *Atomic) Store(t Time)             { self.dur.Store(int64(t)) }
func (self *Atomic) Swap(new Time) (old Time) { return Time(self.dur.Swap(int64(new))) }
func (self *Atomic) CompareAndSwap(old, new Time) (swapped bool) {
	return self.dur.CompareAndSwap(int64(old), int64(new))
}
func (self *Atomic) Add(d duration.NanoSecond) (new Time) { return Time(self.dur.Add(int64(d))) }

//memar:impl memar/time/protocol.Time
func (self *Atomic) Epoch() time_p.Epoch { return &Epoch }
func (self *Atomic) SecondElapsed() duration.Second {
	var t = self.Load()
	return t.SecondElapsed()
}
func (self *Atomic) NanoInSecondElapsed() duration.NanoInSecond {
	var t = self.Load()
	return t.NanoInSecondElapsed()
}

//memar:impl memar/codec/string/protocol.Stringer
func (self *Atomic) ToString() (s string, err error_p.Error) {
	var t = self.Load()
	return t.ToString()
}
func (self *Atomic) FromString(str string) (err error_p.Error) {
	var t Time
	err = t.FromString(str)
	self.Store(t)
	return
}
