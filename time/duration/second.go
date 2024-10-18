/* For license and copyright information please see the LEGAL file in the code repository */

package duration

// A Second duration represents the elapsed time between two instants as an int64 second count.
// https://en.wikipedia.org/wiki/Second
type Second int64

// Common durations.
const (
	OneSecond Second = 1

	NanoSecondInSecond NanoSecond = coefficientUnit_1000 * NanoSecondInMilliSecond // 1e9

	MicroSecondInSecond MicroSecond = coefficientUnit_1000 * OneMicroSecond

	MilliSecondInSecond MilliSecond = coefficientUnit_1000 * OneMilliSecond
)

func (d *Second) ToNanoSecond() (nsec NanoSecond) {
	// TODO::: check overflow??
	nsec = NanoSecond(*d) * NanoSecondInSecond
	return
}

func (d *Second) FromNanoSecond(nsec NanoSecond) {
	*d = Second(nsec / NanoSecondInSecond)
	// TODO::: return overflow nanosecond??
}
