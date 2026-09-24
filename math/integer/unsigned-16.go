/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	"memar/math/boolean"
	math_errs "memar/math/errors"
	error_p "memar/process/error/protocol"
)

// U16 is unsigned 16 bit integer
type U16 uint16

//memar:impl memar/math/logic/protocol.Equivalence
func (u *U16) Equal(with U16) boolean.Boolean {
	return *u == with
}

//memar:impl memar/codec/string/protocol.Stringer
func (u *U16) ToString() (str string, err error_p.Error) {
	str = strconv.FormatUint(uint64(*u), 10)
	return
}
func (u *U16) FromString(str string) (err error_p.Error) {
	err = u.FromString_Base10(str)
	return
}

func (u *U16) FromString_UnknownBase(str string) (err error_p.Error) {
	return
}

// FromString_Base10 parse given string and returns any error occur in the process.
func (u *U16) FromString_Base10(str string) (err error_p.Error) {
	if str == "" {
		return &math_errs.ErrEmptyValue
	}
	var stringLen = len(str)
	if stringLen > 5 {
		err = &math_errs.ErrValueOutOfRange
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
			return &math_errs.ErrBadValue
		}
		n *= 10
		n += uint32(numSegment)
		if n > uint32(maxUint16Value) {
			return &math_errs.ErrValueOutOfRange
		}
	}
	*u = U16(n)
	return
}

func (u *U16) bitSize() int { return 16 }
