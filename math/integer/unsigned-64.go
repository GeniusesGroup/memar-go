/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	"memar/math/boolean"
	errs "memar/math/errors"
	"memar/protocol"
)

// U64 is unsigned 64 bit integer
type U64 uint64

//memar:impl memar/protocol.DataType_Equal
func (u *U64) Equal(with U64) boolean.Boolean {
	return *u == with
}

//memar:impl memar/protocol.Stringer
func (u *U64) ToString() (str string, err protocol.Error) {
	str = strconv.FormatUint(uint64(*u), 10)
	return
}
func (u *U64) FromString(str string) (err protocol.Error) {
	err = u.FromString_Base10(str)
	return
}

func (u *U64) FromString_UnknownBase(str string) (err protocol.Error) {
	return
}

// FromString_Base10 parse given string and returns any error occur in the process.
func (u *U64) FromString_Base10(str string) (err protocol.Error) {
	if str == "" {
		return &errs.ErrEmptyValue
	}
	var stringLen = len(str)
	if stringLen > 20 {
		err = &errs.ErrValueOutOfRange
		return
	}

	var num uint64
	for i := 0; i < stringLen; i++ {
		var char = str[i]
		var numSegment byte
		switch {
		case '0' <= char && char <= '9':
			numSegment = char - '0'
		default:
			return &errs.ErrBadValue
		}
		var n1 = num * 10
		if n1 < num {
			return &errs.ErrValueOutOfRange
		}
		num = n1 + uint64(numSegment)
	}

	*u = U64(num)
	return
}

func (u *U64) bitSize() int { return 64 }
