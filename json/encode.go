/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"encoding/base64"
	"strconv"

	"libgo/convert"
	"libgo/protocol"
)

// Encoder store data to encode given data by each method!
type Encoder struct {
	buf []byte

	Options struct {
		// quoted causes primitive fields to be encoded inside JSON strings.
		Quoted bool
		// escapeHTML causes '<', '>', and '&' to be escaped in JSON strings.
		EscapeHTML bool
	}
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (e *Encoder) Init(buf []byte) (err protocol.Error) {
	e.buf = buf
	return
}
func (e *Encoder) Reinit() (err protocol.Error) { return }
func (e *Encoder) Deinit() (err protocol.Error) { return }

func (e *Encoder) Buf() []byte { return e.buf }

// AddTrailingComma add last value in Buf as trailing comma
func (e *Encoder) AddTrailingComma() {
	e.buf = append(e.buf, ',')
}

// RemoveTrailingComma remove last value in Buf as trailing comma
func (e *Encoder) RemoveTrailingComma() {
	var lastItemIndex = len(e.buf) - 1
	if e.buf[lastItemIndex] == ',' {
		e.buf = e.buf[:lastItemIndex]
	}
}

// EncodeByte append given byte to Buf
func (e *Encoder) EncodeByte(b byte) {
	e.buf = append(e.buf, b)
}

// EncodeBoolean append given bool as true or false string
func (e *Encoder) EncodeBoolean(b bool) {
	if b {
		e.buf = append(e.buf, "true"...)
	} else {
		e.buf = append(e.buf, "false"...)
	}
}

// EncodeUInt8 append given uint8 number as number string
// TODO::: not efficient enough code
func (e *Encoder) EncodeUInt8(ui uint8) {
	e.buf = strconv.AppendUint(e.buf, uint64(ui), 10)
}

// EncodeUInt16 append given uint16 number as number string
// TODO::: not efficient enough code
func (e *Encoder) EncodeUInt16(ui uint16) {
	e.buf = strconv.AppendUint(e.buf, uint64(ui), 10)
}

// EncodeUInt32 append given uint32 number as number string
func (e *Encoder) EncodeUInt32(ui uint32) {
	e.buf = strconv.AppendUint(e.buf, uint64(ui), 10)
}

// EncodeUInt64 append given uint64 number as number string
func (e *Encoder) EncodeUInt64(ui uint64) {
	e.buf = strconv.AppendUint(e.buf, ui, 10)
}

// EncodeInt64 append given int64 number as number string
func (e *Encoder) EncodeInt64(i int64) {
	e.buf = strconv.AppendInt(e.buf, i, 10)
}

// EncodeFloat32 append given float32 number as number string
func (e *Encoder) EncodeFloat32(f float32) {
	e.buf = strconv.AppendFloat(e.buf, float64(f), 'g', -1, 32)
}

// EncodeFloat64 append given float64 number as number string
func (e *Encoder) EncodeFloat64(f float64) {
	e.buf = strconv.AppendFloat(e.buf, f, 'g', -1, 64)
}

// EncodeString append given string as given format
func (e *Encoder) EncodeString(s string) {
	e.buf = append(e.buf, s...)
}

// EncodeKey append given string key as given format with ""
func (e *Encoder) EncodeKey(s string) {
	e.buf = append(e.buf, '"')
	e.buf = append(e.buf, s...)
	e.buf = append(e.buf, `":`...)
}

// EncodeStringValue append given string as given format
func (e *Encoder) EncodeStringValue(s string) {
	e.buf = append(e.buf, '"')
	e.buf = append(e.buf, s...)
	e.buf = append(e.buf, `",`...)
}

/*
	Slice as Number

encode nil slice as empty array ("key":[],) not null("key":null,) value.
*/

// EncodeByteSliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeByteSliceAsNumber(slice []byte) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.buf = strconv.AppendUint(e.buf, uint64(slice[i]), 10)
		e.buf = append(e.buf, ',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt16SliceAsNumber append given uint16 slice as number string
func (e *Encoder) EncodeUInt16SliceAsNumber(slice []uint16) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.buf = strconv.AppendUint(e.buf, uint64(slice[i]), 10)
		e.buf = append(e.buf, ',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt32SliceAsNumber append given uint32 slice as number string
func (e *Encoder) EncodeUInt32SliceAsNumber(slice []uint32) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.buf = strconv.AppendUint(e.buf, uint64(slice[i]), 10)
		e.buf = append(e.buf, ',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt64SliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeUInt64SliceAsNumber(slice []uint64) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.buf = strconv.AppendUint(e.buf, slice[i], 10)
		e.EncodeByte(',')
	}
	e.RemoveTrailingComma()
}

/*
	Slice as Base64

encode nil slice as empty string ("key":"",) not null("key":null,) value.
*/

// EncodeByteSliceAsBase64 use to append []byte as base64 string
func (e *Encoder) EncodeByteSliceAsBase64(slice []byte) {
	var base64Len int = base64.RawStdEncoding.EncodedLen(len(slice))
	var ln = len(e.buf)
	e.buf = e.buf[:ln+base64Len]
	base64.RawStdEncoding.Encode(e.buf[ln:], slice)
}

// EncodeUInt16SliceAsBase64 use to append []byte as base64 string
func (e *Encoder) EncodeUInt16SliceAsBase64(slice []uint16) {
	var base64Len int = base64.RawStdEncoding.EncodedLen(len(slice) * 2)
	var ln = len(e.buf)
	e.buf = e.buf[:ln+base64Len]
	base64.RawStdEncoding.Encode(e.buf[ln:], convert.UnsafeUInt16SliceToByteSlice(slice))
}

// EncodeUInt32SliceAsBase64 use to append []byte as base64 string
func (e *Encoder) EncodeUInt32SliceAsBase64(slice []uint32) {
	var base64Len int = base64.RawStdEncoding.EncodedLen(len(slice) * 4)
	var ln = len(e.buf)
	e.buf = e.buf[:ln+base64Len]
	base64.RawStdEncoding.Encode(e.buf[ln:], convert.UnsafeUInt32SliceToByteSlice(slice))
}

// Encode32ByteArraySliceAsBase64 use to append [][32]byte as base64 string
func (e *Encoder) Encode32ByteArraySliceAsBase64(slice [][32]byte) {
	const base64Len = 43 // base64.RawStdEncoding.EncodedLen(len(32))	>>	(32*8 + 5) / 6
	for _, s := range slice {
		e.buf = append(e.buf, '"')
		var ln = len(e.buf)
		e.buf = e.buf[:ln+base64Len]
		base64.RawStdEncoding.Encode(e.buf[ln:], s[:])
		e.buf = append(e.buf, `",`...)
	}
	e.RemoveTrailingComma()
}
