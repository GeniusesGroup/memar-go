/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"encoding/base64"
	"strconv"

	"../convert"
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
func (d *DecoderUnsafeMinifed) IterationCheck() error {
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
func (d *DecoderUnsafeMinifed) DecodeBool() (b bool, err error) {
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
func (d *DecoderUnsafeMinifed) DecodeUInt8() (ui uint8, err error) {
	d.FindEndToken()
	var num uint64
	num, err = strconv.ParseUint(convert.UnsafeByteSliceToString(d.LastItem), 10, 8)
	ui = uint8(num)
	return
}

// DecodeUInt64 convert 64bit integer number string to number. pass d.Buf start from number and receive from after ,
func (d *DecoderUnsafeMinifed) DecodeUInt64() (ui uint64, err error) {
	d.FindEndToken()
	ui, err = strconv.ParseUint(convert.UnsafeByteSliceToString(d.LastItem), 10, 64)
	return
}

// DecodeInt64 convert 64bit number string to number. pass d.Buf start from number and receive from after ,
func (d *DecoderUnsafeMinifed) DecodeInt64() (i int64, err error) {
	d.FindEndToken()
	i, err = strconv.ParseInt(convert.UnsafeByteSliceToString(d.LastItem), 10, 64)
	return
}

// DecodeFloat64AsNumber convert float64 number string to float64 number. pass d.Buf start from number and receive from ,
func (d *DecoderUnsafeMinifed) DecodeFloat64AsNumber() (f float64, err error) {
	d.FindEndToken()
	f, err = strconv.ParseFloat(convert.UnsafeByteSliceToString(d.LastItem), 64)
	return
}

// DecodeSliceAsNumber convert number string slice to []byte. pass buf start from after [ and receive from after ]
func (d *DecoderUnsafeMinifed) DecodeSliceAsNumber() (slice []byte, err error) {
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

// DecodeSliceAsBase64 convert base64 string to []byte
func (d *DecoderUnsafeMinifed) DecodeSliceAsBase64() (slice []byte, err error) {
	var loc int // Coma, Colon, bracket, ... location
	loc = bytes.IndexByte(d.Buf, '"')
	slice = make([]byte, base64.StdEncoding.DecodedLen(len(d.Buf[:loc])))
	var n int
	n, err = base64.StdEncoding.Decode(slice, d.Buf[:loc])
	if err != nil {
		return
	}
	slice = slice[:n]

	d.Offset(loc + 1)
	return
}

// DecodeArrayAsBase64 convert base64 string to [n]byte
func (d *DecoderUnsafeMinifed) DecodeArrayAsBase64(array []byte) (err error) {
	var loc int // Coma, Colon, bracket, ... location
	loc = bytes.IndexByte(d.Buf, '"')
	_, err = base64.StdEncoding.Decode(array, d.Buf[:loc])
	if err != nil {
		return
	}

	d.Offset(loc + 1)
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
