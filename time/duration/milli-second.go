/* For license and copyright information please see the LEGAL file in the code repository */

package duration

// A MilliSecond duration represents the elapsed time between two instants as an int64 milli-second count.
// The representation limits the largest representable duration to approximately 290 earth years.
type MilliSecond int64

func (d *MilliSecond) FromSecAndNano(sec Second, nsec NanoInSecond) {
	*d = (MilliSecond(sec) * 1e3) + MilliSecond(nsec/1e6)
}
