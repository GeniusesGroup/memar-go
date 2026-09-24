/* For license and copyright information please see the LEGAL file in the code repository */

package duration

import (
	error_p "memar/process/error/protocol"
	"memar/math/integer"
)

// A NanoInSecond represents the elapsed time between two seconds in nano-second count.
type NanoInSecond integer.S32

func (d *NanoInSecond) FromSec(nsec NanoSecond) {
	var sec Second
	sec.FromNanoSecond(nsec)
	*d = NanoInSecond(int64(nsec) % int64(sec))
}

//memar:impl memar/codec/string/protocol.Stringer
func (d *NanoInSecond) ToString() (str string, err error_p.Error) {
	// Go doesn't support to access derived type methods, So we MUST hack it this way!
	// var s32 = integer.S32(*d)
	// return s32.ToString()
	return (*integer.S32)(d).ToString()
}
