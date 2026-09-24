/* For license and copyright information please see the LEGAL file in the code repository */

package duration

import (
	error_p "memar/process/error/protocol"
	"memar/math/integer"
)

// TODO::: Need to check overflow??

// A NanoSecond duration represents the elapsed time between two instants as an int64 nanosecond count.
// The representation limits the largest representable duration to approximately 290 earth years.
type NanoSecond integer.S64

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

//memar:impl memar/codec/string/protocol.Stringer
func (d *NanoSecond) ToString() (str string, err error_p.Error) {
	// Go doesn't support to access derived type methods, So we MUST hack it this way!
	// var s64 = integer.S64(*d)
	// return s64.ToString()
	return (*integer.S64)(d).ToString()
}

// func (d *NanoSecond) FromString(str string) (err error_p.Error) {
// 	strconv.ParseInt(str, 10, s.bitSize())
// 	return
// }
