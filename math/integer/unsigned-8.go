/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	"memar/math/boolean"
	errs "memar/math/errors"
	"memar/protocol"
)

// U8 is unsigned 8 bit integer
type U8 uint8

//memar:impl memar/protocol.DataType_Equal
func (u *U8) Equal(with U8) boolean.Boolean {
	return *u == with
}

//memar:impl memar/protocol.Stringer
func (u *U8) ToString() (str string, err protocol.Error) {
	str = strconv.FormatUint(uint64(*u), 10)
	return
}
func (u *U8) FromString(str string) (err protocol.Error) {
	err = u.FromString_Base10(str)
	return
}

func (u *U8) FromString_UnknownBase(str string) (err protocol.Error) {
	return
}

// FromString_Base10 parse given string and returns any error occur in the process.
func (u *U8) FromString_Base10(str string) (err protocol.Error) {
	if str == "" {
		err = &errs.ErrEmptyValue
		return
	}
	var stringLen = len(str)
	if stringLen > 3 {
		err = &errs.ErrValueOutOfRange
		return
	}

	var n uint16
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
		n += uint16(numSegment)
		if n > uint16(maxUint8Value) {
			return &errs.ErrValueOutOfRange
		}
	}
	*u = U8(n)
	return
}

func (u *U8) bitSize() int { return 8 }
