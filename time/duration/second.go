/* For license and copyright information please see the LEGAL file in the code repository */

package duration

import (
	error_p "memar/process/error/protocol"
	"memar/math/integer"
)

type Field_Second interface {
	// Epoch is not indicate here and MUST indicate by the user.
	SecondDuration() Second
}

// A Second duration represents the elapsed time between two instants as an int64 second count.
// https://en.wikipedia.org/wiki/Second
type Second integer.S64

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

//memar:impl memar/codec/string/protocol.Stringer
func (d *Second) ToString() (str string, err error_p.Error) {
	// Go doesn't support to access derived type methods, So we MUST hack it this way!
	// var s64 = integer.S64(*d)
	// return s64.ToString()
	return (*integer.S64)(d).ToString()
}
