/* For license and copyright information please see the LEGAL file in the code repository */

package duration

// A MicroSecond duration represents the elapsed time between two instants as an int64 micro-second count.
// The representation limits the largest representable duration to approximately 290 earth years.
type MicroSecond int64

func (d *MicroSecond) FromSecAndNano(sec Second, nsec NanoInSecond) {
	*d = (MicroSecond(sec) * 1e6) + MicroSecond(nsec/1e3)
}
