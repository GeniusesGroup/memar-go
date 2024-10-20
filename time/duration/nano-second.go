/* For license and copyright information please see the LEGAL file in the code repository */

package duration

// TODO::: Need to check overflow??

// A NanoSecond duration represents the elapsed time between two instants as an int64 nanosecond count.
// The representation limits the largest representable duration to approximately 290 earth years.
type NanoSecond int64

// Common durations.
const (
	OneNanosecond NanoSecond = 1
)

func (d NanoSecond) ToSecAndNano() (sec Second, nsec NanoInSecond) {
	sec.FromNanoSecond(d)
	// TODO::: Is it worth to uncomment below logic?
	// if sec == 0 {
	// 	nsec = NanoInSecond(d)
	// 	return
	// }
	var secPass = sec.ToNanoSecond()
	nsec = NanoInSecond(d - secPass)
	return
}

func (d *NanoSecond) FromSecAndNano(sec Second, nsec NanoInSecond) {
	*d = sec.ToNanoSecond() + NanoSecond(nsec)
}
