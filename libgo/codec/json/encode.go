/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"encoding/base64"

	buffer_p "memar/computer/buffer/protocol"
	"memar/computer/buffer/byteslice/convert"
	error_p "memar/process/error/protocol"
	"memar/math/boolean"
	"memar/math/float"
	"memar/math/integer"
)

// Encoder provide some functionality over the Buffer to encode easily some data types.
type Encoder struct {
	buf     buffer_p.Buffer
	options EncoderOptions
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (e *Encoder) Init(buf buffer_p.Buffer, opts EncoderOptions) (err error_p.Error) {
	e.buf = buf
	e.options = opts
	return
}
func (e *Encoder) Reinit(buf buffer_p.Buffer, opts EncoderOptions) (err error_p.Error) {
	err = e.Init(buf, opts)
	return
}
func (e *Encoder) Deinit() (err error_p.Error) { return }

func (e *Encoder) Buf() buffer_p.Buffer { return e.buf }

// AddTrailingComma add last value in Buf as trailing comma
func (e *Encoder) AddTrailingComma() {
	e.EncodeByte(',')
}

// RemoveTrailingComma remove last value in Buf as trailing comma
func (e *Encoder) RemoveTrailingComma() {
	var lastItem, _ = e.buf.Peek()
	if lastItem == ',' {
		e.buf.Pop()
	}
}

// EncodeByte append given byte to Buf
func (e *Encoder) EncodeByte(b byte) {
	e.buf.Push(b)
}

// EncodeString append given string as given format
func (e *Encoder) EncodeString(s string) {
	e.buf.Append(convert.UnsafeStringToByteSlice(s)...)
}

// EncodeBoolean append given bool as true or false string
func (e *Encoder) EncodeBoolean(b boolean.Boolean) {
	var str, _ = b.ToString()
	e.EncodeString(str)
}

// Below encode methods append given number as number string.
func (e *Encoder) Encode_Integer_U8(u8 integer.U8) {
	var str, _ = u8.ToString()
	e.EncodeString(str)
}
func (e *Encoder) Encode_Integer_U16(u16 integer.U16) {
	var str, _ = u16.ToString()
	e.EncodeString(str)
}
func (e *Encoder) Encode_Integer_U32(u32 integer.U32) {
	var str, _ = u32.ToString()
	e.EncodeString(str)
}
func (e *Encoder) Encode_Integer_U64(u64 integer.U64) {
	var str, _ = u64.ToString()
	e.EncodeString(str)
}
func (e *Encoder) Encode_Integer_S8(s8 integer.S8) {
	var str, _ = s8.ToString()
	e.EncodeString(str)
}
func (e *Encoder) Encode_Integer_S64(s64 integer.S64) {
	var str, _ = s64.ToString()
	e.EncodeString(str)
}
func (e *Encoder) EncodeFloat32(f float.F32) {
	var str, _ = f.ToString()
	e.EncodeString(str)
}
func (e *Encoder) EncodeFloat64(f float.F64) {
	var str, _ = f.ToString()
	e.EncodeString(str)
}

// EncodeKey append given string key as given format with ""
func (e *Encoder) EncodeKey(s string) {
	e.EncodeByte('"')
	e.EncodeString(s)
	e.EncodeString(`":`)
}

// EncodeStringValue append given string as given format
func (e *Encoder) EncodeStringValue(s string) {
	e.EncodeByte('"')
	e.EncodeString(s)
	e.EncodeString(`",`)
}

/*
	Slice as Number

encode nil slice as empty array ("key":[],) not null("key":null,) value.
*/

// EncodeByteSliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeByteSliceAsNumber(slice []byte) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Encode_Integer_U8(integer.U8(slice[i]))
		e.EncodeByte(',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt16SliceAsNumber append given uint16 slice as number string
func (e *Encoder) EncodeUInt16SliceAsNumber(slice []uint16) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Encode_Integer_U16(integer.U16(slice[i]))
		e.EncodeByte(',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt32SliceAsNumber append given uint32 slice as number string
func (e *Encoder) EncodeUInt32SliceAsNumber(slice []uint32) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Encode_Integer_U32(integer.U32(slice[i]))
		e.EncodeByte(',')
	}
	e.RemoveTrailingComma()
}

// EncodeUInt64SliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeUInt64SliceAsNumber(slice []uint64) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Encode_Integer_U64(integer.U64(slice[i]))
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
		e.EncodeByte('"')
		var ln = len(e.buf)
		e.buf = e.buf[:ln+base64Len]
		base64.RawStdEncoding.Encode(e.buf[ln:], s[:])
		e.EncodeString(`",`)
	}
	e.RemoveTrailingComma()
}
