/* For license and copyright information please see the LEGAL file in the code repository */

package boolean

import (
	"memar/protocol"
)

// Boolean algebra is a branch of mathematics that deals with operations on logical values with binary variables.
// The Boolean variables are represented as binary numbers to represent truths: 1 = true and 0 = false.
// https://en.wikipedia.org/wiki/Boolean_datatype
// https://en.wikipedia.org/wiki/Boolean_algebra
// https://en.wikipedia.org/wiki/Boolean_algebra_(structure)
type Boolean bool

//memar:impl memar/protocol.DataType_Equal
func (b *Boolean) Equal(with Boolean) Boolean {
	return *b == with
}

//memar:impl memar/protocol.Stringer
func (b *Boolean) ToString() (str string, err protocol.Error) {
	if *b {
		str = "true"
	} else {
		str = "false"
	}
	return
}
func (b *Boolean) FromString(str string) (err protocol.Error) {
	switch str {
	case "true", "TRUE", "True":
		*b = true
	case "false", "FALSE", "False":
		*b = false
	default:
		// TODO:::
		// err =
	}
	return
}
