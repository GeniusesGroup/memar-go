/* For license and copyright information please see LEGAL file in repository */

package convert

import "../giti"

const (
	maxUint8Value  uint8  = 255
	maxUint16Value uint16 = 65535
	maxUint32Value uint32 = 4294967295
)

// Base10StringToUint8 Parse given string and returns uint8
func Base10StringToUint8(str string) (num uint8, err giti.Error) {
	if str == "" {
		return 0, ErrEmptyValue
	}

	var n uint16
	for _, char := range []byte(str) {
		var numSegment byte
		// var char = strAsByteSlice[i]
		switch {
		case '0' <= char && char <= '9':
			numSegment = char - '0'
		default:
			return 0, ErrBadValue
		}
		n *= 10
		n += uint16(numSegment)
		if n > uint16(maxUint8Value) {
			return 0, ErrBadValue
		}
	}
	return uint8(n), nil
}

// Base10StringToUint16 Parse given string and returns uint16
func Base10StringToUint16(str string) (num uint16, err giti.Error) {
	if str == "" {
		return 0, ErrEmptyValue
	}

	var n uint32
	for _, char := range []byte(str) {
		var numSegment byte
		// var char = strAsByteSlice[i]
		switch {
		case '0' <= char && char <= '9':
			numSegment = char - '0'
		default:
			return 0, ErrBadValue
		}
		n *= 10
		n += uint32(numSegment)
		if n > uint32(maxUint16Value) {
			return 0, ErrBadValue
		}
	}
	return uint16(n), nil
}

// Base10StringToUint32 Parse given string and returns uint32
func Base10StringToUint32(str string) (num uint32, err giti.Error) {
	if str == "" {
		return 0, ErrEmptyValue
	}

	var n uint64
	for _, char := range []byte(str) {
		var numSegment byte
		// var char = strAsByteSlice[i]
		switch {
		case '0' <= char && char <= '9':
			numSegment = char - '0'
		default:
			return 0, ErrBadValue
		}
		n *= 10
		n += uint64(numSegment)
		if n > uint64(maxUint32Value) {
			return 0, ErrBadValue
		}
	}
	return uint32(n), nil
}

// Base10StringToUint64 Parse given string and returns uint64
func Base10StringToUint64(str string) (num uint64, err giti.Error) {
	if str == "" {
		return 0, ErrEmptyValue
	}

	for _, char := range []byte(str) {
		var numSegment byte
		// var char = strAsByteSlice[i]
		switch {
		case '0' <= char && char <= '9':
			numSegment = char - '0'
		default:
			return 0, ErrBadValue
		}
		num *= 10
		num += uint64(numSegment)
	}
	return
}
