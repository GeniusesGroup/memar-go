/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	"memar/math/boolean"
	math_errs "memar/math/errors"
	error_p "memar/process/error/protocol"
)

// U32 is unsigned 32 bit integer
type U32 uint32

//memar:impl memar/math/logic/protocol.Equivalence
func (u *U32) Equal(with U32) boolean.Boolean {
	return *u == with
}

//memar:impl memar/codec/string/protocol.Stringer
func (u *U32) ToString() (str string, err error_p.Error) {
	str = strconv.FormatUint(uint64(*u), 10)
	return
}
func (u *U32) FromString(str string) (err error_p.Error) {
	err = u.FromString_Base10(str)
	return
}

func (u *U32) FromString_UnknownBase(str string) (err error_p.Error) {
	return
}

// FromString_Base10 parse given string and returns any error occur in the process.
func (u *U32) FromString_Base10(str string) (err error_p.Error) {
	if str == "" {
		return &math_errs.ErrEmptyValue
	}
	var stringLen = len(str)
	if stringLen > 10 {
		err = &math_errs.ErrValueOutOfRange
		return
	}

	var n uint64
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
		n += uint64(numSegment)
		if n > uint64(maxUint32Value) {
			return &math_errs.ErrValueOutOfRange
		}
	}
	*u = U32(n)
	return
}

func (u *U32) bitSize() int { return 32 }
