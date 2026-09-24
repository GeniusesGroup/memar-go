/* For license and copyright information please see the LEGAL file in the code repository */

package float

import (
	"strconv"

	error_p "memar/process/error/protocol"
	"memar/math/boolean"
)

// F64 is 64-bit floating point number.
type F64 float64

//memar:impl memar/math/logic/protocol.Equivalence
func (f *F64) Equal(with F64) boolean.Boolean {
	return *f == with
}

//memar:impl memar/codec/string/protocol.Stringer
func (f *F64) ToString() (str string, err error_p.Error) {
	str = strconv.FormatFloat(float64(*f), 'g', -1, f.bitSize())
	return
}
func (f *F64) FromString(str string) (err error_p.Error) {
	// TODO::: error handling
	var f64, _ = strconv.ParseFloat(str, f.bitSize())
	*f = F64(f64)
	return
}

func (f *F64) bitSize() int { return 64 }
