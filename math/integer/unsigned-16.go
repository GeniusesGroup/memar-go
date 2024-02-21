/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	"memar/math/boolean"
	errs "memar/math/errors"
	"memar/protocol"
)

// U16 is unsigned 16 bit integer
type U16 uint16

//memar:impl memar/protocol.DataType_Equal
func (u *U16) Equal(with U16) boolean.Boolean {
	return *u == with
}

//memar:impl memar/protocol.Stringer
func (u *U16) ToString() (str string, err protocol.Error) {
	str = strconv.FormatUint(uint64(*u), 10)
	return
}
func (u *U16) FromString(str string) (err protocol.Error) {
	err = u.FromString_Base10(str)
	return
}

func (u *U16) FromString_UnknownBase(str string) (err protocol.Error) {
	return
}

// FromString_Base10 parse given string and returns any error occur in the process.
func (u *U16) FromString_Base10(str string) (err protocol.Error) {
	if str == "" {
		return &errs.ErrEmptyValue
	}
	var stringLen = len(str)
	if stringLen > 5 {
		err = &errs.ErrValueOutOfRange
		return
	}

	var n uint32
	for i := 0; i < stringLen; i++ {
		var char = str[i]
		var numSegment byte
		switch {
		case '0' <= char && char <= '9':
			numSegment = char - '0'
		default:
			return &errs.ErrBadValue
		}
		n *= 10
		n += uint32(numSegment)
		if n > uint32(maxUint16Value) {
			return &errs.ErrValueOutOfRange
		}
	}
	*u = U16(n)
	return
}

func (u *U16) bitSize() int { return 16 }
