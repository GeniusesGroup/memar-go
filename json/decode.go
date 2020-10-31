/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"encoding/base64"
	"strconv"

	"../convert"
	er "../error"
)

// Decoder store data to decode data by each method!
type Decoder struct {
	Buf      []byte
	Token    byte
	LastItem []byte
}

// Offset make d.Buf to start of given offset
func (d *Decoder) Offset(o int) {
	d.Buf = d.Buf[o:]
}

// FindEndToken find next end json token
func (d *Decoder) FindEndToken() {
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
func (d *Decoder) ResetToken() {
	d.Token = 0
}

// CheckToken set d.Token to nil
func (d *Decoder) CheckToken(t byte) bool {
	if d.Token == t {
		d.ResetToken()
		return true
	}
	return false
}

// DecodeKey return key very safe for each decode iteration. pass d.Buf start from any where and receive from after :
func (d *Decoder) DecodeKey() string {
	var loc = bytes.IndexByte(d.Buf, '"')
	d.Buf = d.Buf[loc+1:] // remove any byte before first coma due to don't need them
	loc = bytes.IndexByte(d.Buf, '"')

	var slice []byte
	slice = slice[:loc]

	d.Buf = d.Buf[loc+1:] // remove any byte before last coma due to don't need them
	loc = bytes.IndexByte(d.Buf, ':')
	d.Buf = d.Buf[loc+1:]
	return string(slice)
}

// DecodeUInt8 convert 8bit integer number string to number. pass d.Buf start from number and receive from after ,
func (d *Decoder) DecodeUInt8() (ui uint8, err *er.Error) {
	d.FindEndToken()
	ui, err = convert.Base10StringToUint8(convert.UnsafeByteSliceToString(d.LastItem))
	if err != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeUInt64 convert 64bit integer number string to number. pass d.Buf start from after : and receive from after ,
func (d *Decoder) DecodeUInt64() (ui uint64, err *er.Error) {
	d.FindEndToken()
	var goErr error
	ui, goErr = strconv.ParseUint(convert.UnsafeByteSliceToString(d.LastItem), 10, 64)
	if goErr != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeInt64 convert 64bit number string to number. pass d.Buf start from number and receive from after ,
func (d *Decoder) DecodeInt64() (i int64, err *er.Error) {
	d.FindEndToken()
	var goErr error
	i, goErr = strconv.ParseInt(convert.UnsafeByteSliceToString(d.LastItem), 10, 64)
	if goErr != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeFloat64AsNumber convert float64 number string to float64 number. pass d.Buf start from after : and receive from ,
func (d *Decoder) DecodeFloat64AsNumber() (f float64, err *er.Error) {
	d.FindEndToken()
	var goErr error
	f, goErr = strconv.ParseFloat(convert.UnsafeByteSliceToString(d.LastItem), 64)
	if goErr != nil {
		return 0, ErrJSONEncodedStringCorrupted
	}
	return
}

// DecodeByteSliceAsNumber convert number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *Decoder) DecodeByteSliceAsNumber() (slice []byte, err *er.Error) {
	var loc int // Coma, Colon, bracket, ... location
	var num uint8

	loc = bytes.IndexByte(d.Buf, '[')
	d.Buf = d.Buf[loc+1:]
	slice = make([]byte, 0, 8) // TODO::: Is cap efficient enough?

	for !d.CheckToken(']') {
		num, err = d.DecodeUInt8()
		if err != nil {
			return
		}
		slice = append(slice, num)

		d.FindEndToken()
	}
	return
}

// DecodeByteSliceAsBase64 convert base64 string to []byte
func (d *Decoder) DecodeByteSliceAsBase64() (slice []byte, err *er.Error) {
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

	d.Buf = d.Buf[loc+1:]
	return
}

// DecodeByteArrayAsBase64 convert base64 string to [n]byte
func (d *Decoder) DecodeByteArrayAsBase64(array []byte) (err *er.Error) {
	var loc int // Coma, Colon, bracket, ... location
	loc = bytes.IndexByte(d.Buf, '"')
	var goErr error
	_, goErr = base64.RawStdEncoding.Decode(array, d.Buf[:loc])
	if goErr != nil {
		return ErrJSONEncodedStringCorrupted
	}

	d.Buf = d.Buf[loc+1:]
	return
}

// DecodeString return string. pass d.Buf start from after " and receive from from after "
func (d *Decoder) DecodeString() (s string) {
	var loc int // Coma, Colon, bracket, ... location

	loc = bytes.IndexByte(d.Buf, '"')
	if loc < 0 {
		// Reach last item of d.Buf!
		loc = len(d.Buf) - 1
	}

	var slice []byte
	slice = slice[:loc]

	d.Buf = d.Buf[loc+1:]
	return convert.UnsafeByteSliceToString(slice)
}
