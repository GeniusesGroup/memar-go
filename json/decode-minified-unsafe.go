/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"encoding/base64"
	"strconv"

	"../convert"
	er "../error"
)

// DecoderUnsafeMinifed store data to decode data by each method!
type DecoderUnsafeMinifed struct {
	Buf      []byte
	Token    byte
	LastItem []byte
	Found    bool
}

// Offset make d.Buf to start of given offset
func (d *DecoderUnsafeMinifed) Offset(o int) {
	d.Buf = d.Buf[o:]
}

// SetFounded indicate that iteration is legal
func (d *DecoderUnsafeMinifed) SetFounded() {
	d.ResetToken()
	d.Found = true
}

// IterationCheck call in end of each decode iteration.
func (d *DecoderUnsafeMinifed) IterationCheck() *er.Error {
	if d.Found {
		d.Found = false
		return nil
	}
	return ErrJSONEncodedStringCorrupted
}

// FindEndToken find next end json token
func (d *DecoderUnsafeMinifed) FindEndToken() {
	for i, c := range d.Buf {
		switch c {
		case ',':
			d.Token = ','
			d.LastItem = d.Buf[:i]
			d.Buf = d.Buf[i:]
			return
		case ']':
			d.Token = ']'
			d.LastItem = d.Buf[:i]
			d.Buf = d.Buf[i:]
			return
		case '}':
			d.Token = '}'
			d.LastItem = d.Buf[:i]
			d.Buf = d.Buf[i:]
			return
		}
	}
}

// ResetToken set d.Token to nil
func (d *DecoderUnsafeMinifed) ResetToken() {
	d.Token = 0
}

// CheckToken set d.Token to nil
func (d *DecoderUnsafeMinifed) CheckToken(t byte) bool {
	if d.Token == t {
		d.ResetToken()
		return true
	}
	return false
}

// DecodeKey return json key. pass d.Buf start from after " and receive from after :
func (d *DecoderUnsafeMinifed) DecodeKey() string {
	var loc = bytes.IndexByte(d.Buf, '"')
	var slice []byte = d.Buf[:loc]
	d.Offset(loc + 2) // +2 due to have ':"' after key name end!
	return convert.UnsafeByteSliceToString(slice)
}

// DecodeBool convert 64bit integer number string to number. pass d.Buf start from after : and receive from after ,
func (d *DecoderUnsafeMinifed) DecodeBool() (b bool, err *er.Error) {
	if d.Buf[0] == 't' {
		b = true
		d.Offset(5) // true,
	} else {
		// b = false
		d.Offset(6) // false,
	}
	return
}

// DecodeUInt8 convert 8bit integer number string to number. pass d.Buf start from number and receive from after ,
func (d *DecoderUnsafeMinifed) DecodeUInt8() (ui uint8, err *er.Error) {
	d.FindEndToken()
	ui, err = convert.Base10StringToUint8(convert.UnsafeByteSliceToString(d.LastItem))
	if err != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeUInt16 convert 16bit integer number string to number. pass d.Buf start from number and receive from after ,
func (d *DecoderUnsafeMinifed) DecodeUInt16() (ui uint16, err *er.Error) {
	d.FindEndToken()
	ui, err = convert.Base10StringToUint16(convert.UnsafeByteSliceToString(d.LastItem))
	if err != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeUInt32 convert 32bit integer number string to number. pass d.Buf start from number and receive from after ,
func (d *DecoderUnsafeMinifed) DecodeUInt32() (ui uint32, err *er.Error) {
	d.FindEndToken()
	ui, err = convert.Base10StringToUint32(convert.UnsafeByteSliceToString(d.LastItem))
	if err != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeUInt64 convert 64bit integer number string to number. pass d.Buf start from number and receive from after ,
func (d *DecoderUnsafeMinifed) DecodeUInt64() (ui uint64, err *er.Error) {
	d.FindEndToken()
	var goErr error
	ui, goErr = strconv.ParseUint(convert.UnsafeByteSliceToString(d.LastItem), 10, 64)
	if goErr != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeInt64 convert 64bit number string to number. pass d.Buf start from number and receive from after ,
func (d *DecoderUnsafeMinifed) DecodeInt64() (i int64, err *er.Error) {
	d.FindEndToken()
	var goErr error
	i, goErr = strconv.ParseInt(convert.UnsafeByteSliceToString(d.LastItem), 10, 64)
	if goErr != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeFloat64AsNumber convert float64 number string to float64 number. pass d.Buf start from number and receive from ,
func (d *DecoderUnsafeMinifed) DecodeFloat64AsNumber() (f float64, err *er.Error) {
	d.FindEndToken()
	var goErr error
	f, goErr = strconv.ParseFloat(convert.UnsafeByteSliceToString(d.LastItem), 64)
	if goErr != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeString return string. pass d.Buf start from after " and receive from from after "
func (d *DecoderUnsafeMinifed) DecodeString() (s string) {
	var loc int // Coma, Colon, bracket, ... location
	loc = bytes.IndexByte(d.Buf, '"')
	if loc < 0 {
		// Reach last item of d.Buf!
		loc = len(d.Buf) - 1
	}

	var slice []byte = d.Buf[:loc]
	d.Offset(loc + 1)
	return convert.UnsafeByteSliceToString(slice)
}

/*
	Array part
*/

// DecodeByteArrayAsBase64 convert base64 string to [n]byte
func (d *DecoderUnsafeMinifed) DecodeByteArrayAsBase64(array []byte) (err *er.Error) {
	var loc int // Coma, Colon, bracket, ... location
	loc = bytes.IndexByte(d.Buf, '"')
	// var goErr error
	base64.RawStdEncoding.Decode(array, d.Buf[:loc])
	// log.Info(goErr)
	// if goErr != nil {
	// 	return ErrJSONEncodedStringCorrupted
	// }

	d.Offset(loc + 1)
	return
}

/*
	Slice as Number
*/

// DecodeByteSliceAsNumber convert number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderUnsafeMinifed) DecodeByteSliceAsNumber() (slice []byte, err *er.Error) {
	var num uint8
	slice = make([]byte, 0, 8) // TODO::: Is cap efficient enough?

	for !d.CheckToken(']') {
		num, err = d.DecodeUInt8()
		if err != nil {
			return
		}
		slice = append(slice, num)
		d.Offset(1)
	}
	return
}

// DecodeUInt16SliceAsNumber convert uint16 number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderUnsafeMinifed) DecodeUInt16SliceAsNumber() (slice []uint16, err *er.Error) {
	var num uint16
	slice = make([]uint16, 0, 8) // TODO::: Is cap efficient enough?

	for !d.CheckToken(']') {
		num, err = d.DecodeUInt16()
		if err != nil {
			return
		}
		slice = append(slice, num)
		d.Offset(1)
	}
	return
}

// DecodeUInt32SliceAsNumber convert uint32 number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderUnsafeMinifed) DecodeUInt32SliceAsNumber() (slice []uint32, err *er.Error) {
	var num uint32
	slice = make([]uint32, 0, 8) // TODO::: Is cap efficient enough?

	for !d.CheckToken(']') {
		num, err = d.DecodeUInt32()
		if err != nil {
			return
		}
		slice = append(slice, num)
		d.Offset(1)
	}
	return
}

// DecodeUInt64SliceAsNumber convert uint64 number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderUnsafeMinifed) DecodeUInt64SliceAsNumber() (slice []uint64, err *er.Error) {
	var num uint64
	slice = make([]uint64, 0, 8) // TODO::: Is cap efficient enough?

	for !d.CheckToken(']') {
		num, err = d.DecodeUInt64()
		if err != nil {
			return
		}
		slice = append(slice, num)
		d.Offset(1)
	}
	return
}

/*
	Slice as Base64
*/

// DecodeByteSliceAsBase64 convert base64 string to []byte
func (d *DecoderUnsafeMinifed) DecodeByteSliceAsBase64() (slice []byte, err *er.Error) {
	var loc int // Coma, Colon, bracket, ... location
	loc = bytes.IndexByte(d.Buf, '"')
	slice = make([]byte, base64.RawStdEncoding.DecodedLen(len(d.Buf[:loc])))
	var n int
	var goErr error
	n, goErr = base64.RawStdEncoding.Decode(slice, d.Buf[:loc])
	if goErr != nil {
		return slice, ErrJSONEncodedStringCorrupted
	}
	slice = slice[:n]

	d.Offset(loc + 1)
	return
}

// Decode32ByteArraySliceAsBase64 decode [32]byte base64 string slice. pass buf start from after [ and receive from after ]
func (d *DecoderUnsafeMinifed) Decode32ByteArraySliceAsBase64() (slice [][32]byte, err *er.Error) {
	const base64Len = 43 // base64.RawStdEncoding.EncodedLen(len(32))	>>	(32*8 + 5) / 6
	slice = make([][32]byte, 0, 8)

	var goErr error
	var array [32]byte
	for d.Buf[1] != ']' {
		d.Offset(2) // due to have `["` || `",`
		_, goErr = base64.RawStdEncoding.Decode(array[:], d.Buf[:base64Len])
		if goErr != nil {
			return slice, ErrJSONEncodedStringCorrupted
		}
		slice = append(slice, array)
		d.Buf = d.Buf[base64Len:]
	}

	d.Offset(2) // due to have	`"]`
	return
}
