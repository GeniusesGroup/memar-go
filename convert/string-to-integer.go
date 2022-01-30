/* For license and copyright information please see LEGAL file in repository */

package convert

import "../protocol"

const (
	maxUint8Value  uint8  = 255
	maxUint16Value uint16 = 65535
	maxUint32Value uint32 = 4294967295
)

// StringToUint8Base10 Parse given string and returns uint8
func StringToUint8Base10(str string) (num uint8, err protocol.Error) {
	if str == "" {
		err = ErrEmptyValue
		return
	}
	var stringLen = len(str)
	if stringLen > 3 {
		err = ErrValueOutOfRange
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
			return 0, ErrBadValue
		}
		n *= 10
		n += uint16(numSegment)
		if n > uint16(maxUint8Value) {
			return 0, ErrValueOutOfRange
		}
	}
	return uint8(n), nil
}

// StringToUint16Base10 Parse given string and returns uint16
func StringToUint16Base10(str string) (num uint16, err protocol.Error) {
	if str == "" {
		return 0, ErrEmptyValue
	}
	var stringLen = len(str)
	if stringLen > 5 {
		err = ErrValueOutOfRange
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
			return 0, ErrBadValue
		}
		n *= 10
		n += uint32(numSegment)
		if n > uint32(maxUint16Value) {
			return 0, ErrValueOutOfRange
		}
	}
	return uint16(n), nil
}

// StringToUint32Base10 Parse given string and returns uint32
func StringToUint32Base10(str string) (num uint32, err protocol.Error) {
	if str == "" {
		return 0, ErrEmptyValue
	}
	var stringLen = len(str)
	if stringLen > 10 {
		err = ErrValueOutOfRange
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
			return 0, ErrBadValue
		}
		n *= 10
		n += uint64(numSegment)
		if n > uint64(maxUint32Value) {
			return 0, ErrValueOutOfRange
		}
	}
	return uint32(n), nil
}

// StringToUint64Base10 Parse given string and returns uint64
func StringToUint64Base10(str string) (num uint64, err protocol.Error) {
	if str == "" {
		return 0, ErrEmptyValue
	}
	var stringLen = len(str)
	if stringLen > 20 {
		err = ErrValueOutOfRange
		return
	}

	for i := 0; i < stringLen; i++ {
		var char = str[i]
		var numSegment byte
		switch {
		case '0' <= char && char <= '9':
			numSegment = char - '0'
		default:
			return 0, ErrBadValue
		}
		var n1 = num * 10
		if n1 < num {
			return 0, ErrValueOutOfRange
		}
		num = n1 + uint64(numSegment)
	}
	return
}
