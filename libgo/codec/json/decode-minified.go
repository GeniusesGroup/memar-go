/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"encoding/base64"
	"strconv"

	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	errs "memar/codec/data_exchange/json/errors"
	error_p "memar/process/error/protocol"
	"memar/math/integer"
)

// DecoderMinified store data to decode data by each method!
type DecoderMinified[BUF buffer_p.Buffer] struct {
	Decoder[BUF]
}

// DecodeKey return json key. pass d.buf start from after {||, and receive from after :
func (d *DecoderMinified[BUF]) DecodeKey() string {
	d.Offset(2)
	var loc, _ = d.buf.Index('"', 0, 0)
	var slice []byte = d.buf[:loc]
	d.Offset(loc + 2) // +2 due to have '":' after key name end!
	return string(slice)
}

// DecodeBool convert 64bit integer number string to number. pass d.buf start from after : and receive from after ,
func (d *DecoderMinified[BUF]) DecodeBool() (b bool, err error_p.Error) {
	if d.buf[0] == 't' {
		b = true
		d.Offset(5) // true,
	} else {
		// b = false
		d.Offset(6) // false,
	}
	return
}

// Decode_Integer_U8 convert 8bit integer number string to number. pass d.buf start from number and receive from after ,
func (d *DecoderMinified[BUF]) Decode_Integer_U8() (ui integer.U8, err error_p.Error) {
	d.FindEndToken()
	err = ui.FromString_Base10(d.LastItem())
	if err != nil {
		// err = &errs.EncodedIntegerCorrupted
		return
	}
	return
}

// Decode_Integer_U16 convert 16bit integer number string to number. pass d.buf start from number and receive from after ,
func (d *DecoderMinified[BUF]) Decode_Integer_U16() (ui integer.U16, err error_p.Error) {
	d.FindEndToken()
	err = ui.FromString_Base10(d.LastItem())
	if err != nil {
		// err = &errs.EncodedIntegerCorrupted
		return
	}
	return
}

// Decode_Integer_U32 convert 32bit integer number string to number. pass d.buf start from number and receive from after ,
func (d *DecoderMinified[BUF]) Decode_Integer_U32() (ui integer.U32, err error_p.Error) {
	d.FindEndToken()
	err = ui.FromString_Base10(d.LastItem())
	if err != nil {
		// err = &errs.EncodedIntegerCorrupted
		return
	}
	return
}

// Decode_Integer_U64 convert 64bit integer number string to number. pass d.buf start from number and receive from after ,
func (d *DecoderMinified[BUF]) Decode_Integer_U64() (ui integer.U64, err error_p.Error) {
	d.FindEndToken()
	err = ui.FromString_Base10(d.LastItem())
	if err != nil {
		// err = &errs.EncodedIntegerCorrupted
		return
	}
	return
}

// DecodeInt32 convert 32bit number string to number. pass d.buf start from number and receive from after end of number
func (d *DecoderUnsafe[BUF]) DecodeInt32() (i int32, err error_p.Error) {
	d.FindEndToken()
	var goErr error
	var num int64
	num, goErr = strconv.ParseInt(d.LastItem(), 10, 32)
	if goErr != nil {
		return 0, &errs.EncodedIntegerCorrupted
	}
	i = int32(num)
	return
}

// DecodeInt64 convert 64bit number string to number. pass d.buf start from number and receive from after ,
func (d *DecoderMinified[BUF]) DecodeInt64() (i int64, err error_p.Error) {
	d.FindEndToken()
	var goErr error
	i, goErr = strconv.ParseInt(d.LastItem(), 10, 64)
	if goErr != nil {
		return 0, &errs.EncodedStringCorrupted
	}
	return
}

// DecodeFloat64AsNumber convert float64 number string to float64 number. pass d.buf start from number and receive from ,
func (d *DecoderMinified[BUF]) DecodeFloat64AsNumber() (f float64, err error_p.Error) {
	d.FindEndToken()
	var goErr error
	f, goErr = strconv.ParseFloat(d.LastItem(), 64)
	if goErr != nil {
		return 0, &errs.EncodedStringCorrupted
	}
	return
}

// DecodeString return string. pass d.buf start from after " and receive from from after "
func (d *DecoderMinified[BUF]) DecodeString() (s string) {
	var loc container_p.ElementIndex // Coma, Colon, bracket, ... location
	loc, _ = d.buf.Index('"')
	if loc < 0 {
		// Reach last item of d.buf!
		loc = len(d.buf) - 1
	}

	var slice []byte = d.buf[:loc]
	d.Offset(loc + 1)
	return string(slice)
}

/*
	Array part
*/

// DecodeByteArrayAsBase64 convert base64 string to [n]byte
func (d *DecoderMinified[BUF]) DecodeByteArrayAsBase64(array []byte) (err error_p.Error) {
	d.Offset(1) // due to have " at start

	var loc, _ = d.buf.Index('"')
	if loc < 0 {
		// err = &errs.EncodedArrayCorrupted
		return
	}

	var goErr error
	_, goErr = base64.RawStdEncoding.Decode(array, d.buf[:loc])
	if goErr != nil {
		return &errs.EncodedArrayCorrupted
	}

	d.Offset(loc + 1)
	return
}

// DecodeByteArrayAsNumber convert number array to [n]byte
func (d *DecoderMinified[BUF]) DecodeByteArrayAsNumber(array []byte) (err error_p.Error) {
	var value uint8
	for i := 0; i < len(array); i++ {
		d.Offset(1) // due to have [ or ,
		value, err = d.Decode_Integer_U8()
		if err != nil {
			// err = &errs.EncodedArrayCorrupted
			return
		}
		array[i] = value
	}
	if d.buf[0] != ']' {
		// err = &errs.EncodedArrayCorrupted
	}
	d.Offset(1)
	return
}

/*
	Slice as Number
*/

// DecodeByteSliceAsNumber convert number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderMinified[BUF]) DecodeByteSliceAsNumber() (slice []byte, err error_p.Error) {
	d.Offset(1)                // due to have [ at start
	slice = make([]byte, 0, 8) // TODO::: Is cap efficient enough?

	var num integer.U8
	for !d.CheckToken(']') {
		num, err = d.Decode_Integer_U8()
		if err != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, byte(num))
		d.Offset(1)
	}
	return
}

// Decode_Integer_U16SliceAsNumber convert uint16 number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderMinified[BUF]) Decode_Integer_U16SliceAsNumber() (slice []uint16, err error_p.Error) {
	d.Offset(1)                  // due to have [ at start
	slice = make([]uint16, 0, 8) // TODO::: Is cap efficient enough?

	var num integer.U16
	for !d.CheckToken(']') {
		num, err = d.Decode_Integer_U16()
		if err != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, uint16(num))
		d.Offset(1)
	}
	return
}

// Decode_Integer_U32SliceAsNumber convert uint32 number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderMinified[BUF]) Decode_Integer_U32SliceAsNumber() (slice []uint32, err error_p.Error) {
	d.Offset(1)                  // due to have [ at start
	slice = make([]uint32, 0, 8) // TODO::: Is cap efficient enough?

	var num integer.U32
	for !d.CheckToken(']') {
		num, err = d.Decode_Integer_U32()
		if err != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, uint32(num))
		d.Offset(1)
	}
	return
}

// Decode_Integer_U64SliceAsNumber convert uint64 number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderMinified[BUF]) Decode_Integer_U64SliceAsNumber() (slice []uint64, err error_p.Error) {
	d.Offset(1)                  // due to have [ at start
	slice = make([]uint64, 0, 8) // TODO::: Is cap efficient enough?

	var num integer.U64
	for !d.CheckToken(']') {
		num, err = d.Decode_Integer_U64()
		if err != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, uint64(num))
		d.Offset(1)
	}
	return
}

/*
	Slice as Base64
*/

// DecodeByteSliceAsBase64 convert base64 string to []byte
func (d *DecoderMinified[BUF]) DecodeByteSliceAsBase64() (slice []byte, err error_p.Error) {
	d.Offset(1) // due to have " at start

	// Coma, Colon, bracket, ... location
	var loc, _ = d.buf.Index('"')
	slice = make([]byte, base64.RawStdEncoding.DecodedLen(len(d.buf[:loc])))
	var n int
	var goErr error
	n, goErr = base64.RawStdEncoding.Decode(slice, d.buf[:loc])
	if goErr != nil {
		// err = &errs.EncodedSliceCorrupted
		return
	}
	slice = slice[:n]

	d.Offset(loc + 1)
	return
}

// Decode32ByteArraySliceAsBase64 decode [32]byte base64 string slice. pass buf start from after [ and receive from after ]
func (d *DecoderMinified[BUF]) Decode32ByteArraySliceAsBase64() (slice [][32]byte, err error_p.Error) {
	d.Offset(1) // due to have [ at start

	const base64Len = 43 // base64.RawStdEncoding.EncodedLen(len(32))	>>	(32*8 + 5) / 6
	slice = make([][32]byte, 0, 8)

	var goErr error
	var array [32]byte
	for d.buf[1] != ']' {
		d.Offset(2) // due to have `["` || `",`
		_, goErr = base64.RawStdEncoding.Decode(array[:], d.buf[:base64Len])
		if goErr != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, array)
		d.buf = d.buf[base64Len:]
	}

	d.Offset(2) // due to have	`"]`
	return
}
