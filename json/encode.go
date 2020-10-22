/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"encoding/base64"
	"strconv"
)

// Encoder store data to encode given data by each method!
type Encoder struct {
	Buf []byte
}

// RemoveTrailingComma remove last value in Buf as trailing comma
func (e *Encoder) RemoveTrailingComma() {
	e.Buf = e.Buf[:len(e.Buf)-1]
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

// EncodeInt64 append given int64 number as number string
func (e *Encoder) EncodeInt64(i int64) {
	e.Buf = strconv.AppendInt(e.Buf, i, 10)
}

// EncodeUInt64 append given uint64 number as number string
func (e *Encoder) EncodeUInt64(ui uint64) {
	e.Buf = strconv.AppendUint(e.Buf, ui, 10)
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

// EncodeByteSliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeByteSliceAsNumber(slice []byte) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Buf = strconv.AppendUint(e.Buf, uint64(slice[i]), 10)
		e.Buf = append(e.Buf, ',')
	}
	e.RemoveTrailingComma()
}

// EncodeByteSliceAsBase64 use to append []byte as base64 string
func (e *Encoder) EncodeByteSliceAsBase64(slice []byte) {
	var base64Len int = base64.StdEncoding.EncodedLen(len(slice))
	var ln = len(e.Buf)
	e.Buf = e.Buf[:ln+base64Len]
	base64.StdEncoding.Encode(e.Buf[ln:], slice)
}

// EncodeUInt64SliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeUInt64SliceAsNumber(slice []uint64) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Buf = strconv.AppendUint(e.Buf, slice[i], 10)
		e.Buf = append(e.Buf, ',')
	}
	e.RemoveTrailingComma()
}
