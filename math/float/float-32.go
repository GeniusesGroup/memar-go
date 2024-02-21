/* For license and copyright information please see the LEGAL file in the code repository */

package float

import (
	"strconv"

	"memar/math/boolean"
	"memar/protocol"
)

// F32 is 32-bit floating point number.
type F32 float32

//memar:impl memar/protocol.DataType_Equal
func (f *F32) Equal(with F32) boolean.Boolean {
	return *f == with
}

//memar:impl memar/protocol.Stringer
func (f *F32) ToString() (str string, err protocol.Error) {
	str = strconv.FormatFloat(float64(*f), 'g', -1, f.bitSize())
	return
}
func (f *F32) FromString(str string) (err protocol.Error) {
	// TODO::: error handling
	var f32, _ = strconv.ParseFloat(str, f.bitSize())
	*f = F32(f32)
	return
}

func (f *F32) bitSize() int { return 32 }
