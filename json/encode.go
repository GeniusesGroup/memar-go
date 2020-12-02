/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"encoding/base64"
	"strconv"

	"../convert"
)

// Encoder store data to encode given data by each method!
type Encoder struct {
	Buf []byte
}

// RemoveTrailingComma remove last value in Buf as trailing comma
func (e *Encoder) RemoveTrailingComma() {
	if e.Buf[len(e.Buf)-1] == ',' {
		e.Buf = e.Buf[:len(e.Buf)-1]
	}
}

// EncodeByte append given byte to Buf
func (e *Encoder) EncodeByte(b byte) {
	e.Buf = append(e.Buf, b)
}

// EncodeBoolean append given bool as true or false string
func (e *Encoder) EncodeBoolean(b bool) {
	if b {
		e.Buf = append(e.Buf, "true"...)
	} else {
		e.Buf = append(e.Buf, "false"...)
	}
}

// EncodeUInt8 append given uint8 number as number string
// TODO::: not efficient enough code
func (e *Encoder) EncodeUInt8(ui uint8) {
	e.Buf = strconv.AppendUint(e.Buf, uint64(ui), 10)
}

// EncodeUInt16 append given uint16 number as number string
// TODO::: not efficient enough code
func (e *Encoder) EncodeUInt16(ui uint16) {
	e.Buf = strconv.AppendUint(e.Buf, uint64(ui), 10)
}

// EncodeUInt32 append given uint32 number as number string
func (e *Encoder) EncodeUInt32(ui uint32) {
	e.Buf = strconv.AppendUint(e.Buf, uint64(ui), 10)
}

// EncodeUInt64 append given uint64 number as number string
func (e *Encoder) EncodeUInt64(ui uint64) {
	e.Buf = strconv.AppendUint(e.Buf, ui, 10)
}

// EncodeInt64 append given int64 number as number string
func (e *Encoder) EncodeInt64(i int64) {
	e.Buf = strconv.AppendInt(e.Buf, i, 10)
}

// EncodeFloat32 append given float32 number as number string
func (e *Encoder) EncodeFloat32(f float32) {
	e.Buf = strconv.AppendFloat(e.Buf, float64(f), 'g', -1, 32)
}

// EncodeFloat64 append given float64 number as number string
func (e *Encoder) EncodeFloat64(f float64) {
	e.Buf = strconv.AppendFloat(e.Buf, f, 'g', -1, 64)
}

// EncodeString append given string as given format
func (e *Encoder) EncodeString(s string) {
	e.Buf = append(e.Buf, s...)
}

// EncodeKey append given string key as given format with ""
func (e *Encoder) EncodeKey(s string) {
	e.Buf = append(e.Buf, '"')
	e.Buf = append(e.Buf, s...)
	e.Buf = append(e.Buf, `":`...)
}

// EncodeStringValue append given string as given format
func (e *Encoder) EncodeStringValue(s string) {
	e.Buf = append(e.Buf, '"')
	e.Buf = append(e.Buf, s...)
	e.Buf = append(e.Buf, `",`...)
}

/*
	Slice as Number
*/

// EncodeByteSliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeByteSliceAsNumber(slice []byte) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Buf = strconv.AppendUint(e.Buf, uint64(slice[i]), 10)
		e.Buf = append(e.Buf, ',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt16SliceAsNumber append given uint16 slice as number string
func (e *Encoder) EncodeUInt16SliceAsNumber(slice []uint16) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Buf = strconv.AppendUint(e.Buf, uint64(slice[i]), 10)
		e.Buf = append(e.Buf, ',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt32SliceAsNumber append given uint32 slice as number string
func (e *Encoder) EncodeUInt32SliceAsNumber(slice []uint32) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Buf = strconv.AppendUint(e.Buf, uint64(slice[i]), 10)
		e.Buf = append(e.Buf, ',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt64SliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeUInt64SliceAsNumber(slice []uint64) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Buf = strconv.AppendUint(e.Buf, slice[i], 10)
		e.EncodeByte(',')
	}
	e.RemoveTrailingComma()
}

/*
	Slice as Base64
*/

// EncodeByteSliceAsBase64 use to append []byte as base64 string
func (e *Encoder) EncodeByteSliceAsBase64(slice []byte) {
	var base64Len int = base64.RawStdEncoding.EncodedLen(len(slice))
	var ln = len(e.Buf)
	e.Buf = e.Buf[:ln+base64Len]
	base64.RawStdEncoding.Encode(e.Buf[ln:], slice)
}

// EncodeUInt16SliceAsBase64 use to append []byte as base64 string
func (e *Encoder) EncodeUInt16SliceAsBase64(slice []uint16) {
	var base64Len int = base64.RawStdEncoding.EncodedLen(len(slice) * 2)
	var ln = len(e.Buf)
	e.Buf = e.Buf[:ln+base64Len]
	base64.RawStdEncoding.Encode(e.Buf[ln:], convert.UnsafeUInt16SliceToByteSlice(slice))
}

// EncodeUInt32SliceAsBase64 use to append []byte as base64 string
func (e *Encoder) EncodeUInt32SliceAsBase64(slice []uint32) {
	var base64Len int = base64.RawStdEncoding.EncodedLen(len(slice) * 4)
	var ln = len(e.Buf)
	e.Buf = e.Buf[:ln+base64Len]
	base64.RawStdEncoding.Encode(e.Buf[ln:], convert.UnsafeUInt32SliceToByteSlice(slice))
}

// Encode32ByteArraySliceAsBase64 use to append [][32]byte as base64 string
func (e *Encoder) Encode32ByteArraySliceAsBase64(slice [][32]byte) {
	const base64Len = 43 // base64.RawStdEncoding.EncodedLen(len(32))	>>	(32*8 + 5) / 6
	for _, s := range slice {
		e.Buf = append(e.Buf, '"')
		var ln = len(e.Buf)
		e.Buf = e.Buf[:ln+base64Len]
		base64.RawStdEncoding.Encode(e.Buf[ln:], s[:])
		e.Buf = append(e.Buf, `",`...)
	}
	e.RemoveTrailingComma()
}
