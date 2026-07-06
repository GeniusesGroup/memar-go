/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"bytes"
	"encoding/base64"
	"strconv"

	buffer_p "memar/computer/buffer/protocol"
	errs "memar/codec/data_exchange/json/errors"
	"memar/computer/buffer/byteslice/convert"
	error_p "memar/process/error/protocol"
	"memar/math/integer"
)

// Decoder store data to decode data by each method.
type Decoder[BUF buffer_p.Buffer] struct {
	// TODO::: json is string base why not wrap in ASCII to provide more helpful methods??
	buf      BUF
	token    byte
	lastItem []byte
	options  DecodeOptions
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Decoder[BUF]) Init(buf BUF, opts DecodeOptions) (err error_p.Error) {
	self.buf = buf
	self.options = opts
	return
}
func (self *Decoder[BUF]) Reinit() (err error_p.Error) { return }
func (self *Decoder[BUF]) Deinit() (err error_p.Error) { return }

func (self *Decoder[BUF]) Buf() BUF { return self.buf }

func (self *Decoder[BUF]) LastItem() string {
	return self.LastItem()
}

// Offset make self.buf to start of given offset
func (self *Decoder[BUF]) Offset(o int) {
	self.buf = self.buf[o:]
}

// FindEndToken find next end json token
func (self *Decoder[BUF]) FindEndToken() {
	for i, c := range self.buf {
		switch c {
		case ',':
			self.token = ','
			self.lastItem = self.buf[:i]
			self.buf = self.buf[i:]
			return
		case ']':
			self.token = ']'
			self.lastItem = self.buf[:i]
			self.buf = self.buf[i:]
			return
		case '}':
			self.token = '}'
			self.lastItem = self.buf[:i]
			self.buf = self.buf[i:]
			return
		}
	}
}

// TrimToDigit remove any data and set self.buf to first character as number
func (self *Decoder[BUF]) TrimToDigit() {
	for i, c := range self.buf {
		if '0' <= c && c <= '9' {
			self.buf = self.buf[i:]
			return
		}
	}
}

// TrimSpaces remove any spaces from self.buf
func (self *Decoder[BUF]) TrimSpaces() {
	for i, c := range self.buf {
		if c != ' ' {
			self.buf = self.buf[i:]
			return
		}
	}
}

// TrimToStringStart remove any data from self.buf to first character after quotation mark as '"'
func (self *Decoder[BUF]) TrimToStringStart() {
	for i, c := range self.buf {
		if c == '"' {
			self.buf = self.buf[i+1:] // remove any byte before " due to don't need them
			return
		}
	}
}

// CheckNullValue check if null exist as value. pass self.buf start from after : and receive from from after , if null exist
func (self *Decoder[BUF]) CheckNullValue() (null bool) {
	for i, c := range self.buf {
		switch c {
		case 'n':
			if bytes.Equal(self.buf[i:i+4], []byte("null")) {
				null = true
			}
		case '"':
			return false
		case ',':
			return
		}
	}
	return
}

// ResetToken set self.token to nil
func (self *Decoder[BUF]) ResetToken() {
	self.token = 0
}

// CheckToken set self.token to nil
func (self *Decoder[BUF]) CheckToken(t byte) bool {
	if self.token == t {
		self.ResetToken()
		return true
	}
	return false
}

// DecodeKey return key very safe for each decode iteration. pass self.buf start from any where and receive from after :
func (self *Decoder[BUF]) DecodeKey() string {
	self.TrimToStringStart()
	var loc, _ = self.buf.Index('"')
	if loc < 0 {
		return ""
	}

	var key []byte = self.buf[:loc]

	self.buf = self.buf[loc+1:] // remove any byte before last " due to don't need them
	loc, _ = self.buf.Index(':')
	self.buf = self.buf[loc+1:]
	return convert.UnsafeByteSliceToString(key)
}

// NotFoundKey call in default switch of each decode iteration
func (self *Decoder[BUF]) NotFoundKey() (err error_p.Error) {
	self.FindEndToken()
	return
}

// NotFoundKeyStrict call in default switch of each decode iteration in strict mode.
func (self *Decoder[BUF]) NotFoundKeyStrict() error_p.Error {
	return &errs.EncodedIncludeNotDefinedKey
}

func (self *Decoder[BUF]) End() bool {
	if len(self.buf) < 3 || self.token == '}' {
		return true
	}
	return false
}

// DecodeBool convert string base boolean to bool. pass self.buf start from after : and receive from after ,
func (self *Decoder[BUF]) DecodeBool() (b bool, err error_p.Error) {
	self.TrimSpaces()
	if self.buf[0] == 't' {
		b = true
		self.Offset(5) // true,
	} else {
		// b = false
		self.Offset(6) // false,
	}
	return
}

// Decode_Integer_U8 convert 8bit integer number string to number. pass self.buf start from number and receive from after ,
func (self *Decoder[BUF]) Decode_Integer_U8() (ui integer.U8, err error_p.Error) {
	self.TrimToDigit()
	self.FindEndToken()
	err = ui.FromString_Base10(self.LastItem())
	if err != nil {
		// err = &errs.EncodedIntegerCorrupted
		return
	}
	return
}

// Decode_Integer_U16 convert 16bit integer number string to number. pass self.buf start from number and receive from after ,
func (self *Decoder[BUF]) Decode_Integer_U16() (ui integer.U16, err error_p.Error) {
	self.TrimToDigit()
	self.FindEndToken()
	err = ui.FromString_Base10(self.LastItem())
	if err != nil {
		// err = &errs.EncodedIntegerCorrupted
		return
	}
	return
}

// Decode_Integer_U32 convert 32bit integer number string to number. pass self.buf start from number and receive from after ,
func (self *Decoder[BUF]) Decode_Integer_U32() (ui integer.U32, err error_p.Error) {
	self.TrimToDigit()
	self.FindEndToken()
	err = ui.FromString_Base10(self.LastItem())
	if err != nil {
		// err = &errs.EncodedIntegerCorrupted
		return
	}
	return
}

// Decode_Integer_U64 convert 64bit integer number string to number. pass self.buf start from after : and receive from after ,
func (self *Decoder[BUF]) Decode_Integer_U64() (ui integer.U64, err error_p.Error) {
	self.TrimToDigit()
	self.FindEndToken()
	err = ui.FromString_Base10(self.LastItem())
	if err != nil {
		// err = &errs.EncodedIntegerCorrupted
		return
	}
	return
}

// DecodeInt64 convert 64bit number string to number. pass self.buf start from number and receive from after ,
func (self *Decoder[BUF]) DecodeInt64() (i int64, err error_p.Error) {
	self.TrimToDigit()
	self.FindEndToken()
	var goErr error
	i, goErr = strconv.ParseInt(self.LastItem(), 10, 64)
	if goErr != nil {
		return 0, &errs.EncodedIntegerCorrupted
	}
	return
}

// DecodeFloat64AsNumber convert float64 number string to float64 number. pass self.buf start from after : and receive from ,
func (self *Decoder[BUF]) DecodeFloat64AsNumber() (f float64, err error_p.Error) {
	self.TrimToDigit()
	self.FindEndToken()
	var goErr error
	f, goErr = strconv.ParseFloat(self.LastItem(), 64)
	if goErr != nil {
		return 0, &errs.EncodedIntegerCorrupted
	}
	return
}

// DecodeString return string. pass self.buf start from after : and receive from from after "
func (self *Decoder[BUF]) DecodeString() (s string, err error_p.Error) {
	if self.CheckNullValue() {
		return
	}

	self.TrimToStringStart()
	var loc, _ = self.buf.Index('"')
	if loc < 0 {
		// err = &errs.EncodedStringCorrupted
		return
	}

	var slice []byte = self.buf[:loc]

	self.Offset(loc + 1)
	s = string(slice)
	return
}

/*
	Array part
*/

// DecodeByteArrayAsBase64 convert base64 string to [n]byte
func (self *Decoder[BUF]) DecodeByteArrayAsBase64(array []byte) (err error_p.Error) {
	if self.CheckNullValue() {
		return
	}

	self.TrimToStringStart()
	var loc, _ = self.buf.Index('"')
	if loc < 0 {
		// err = &errs.EncodedArrayCorrupted
		return
	}

	var goErr error
	_, goErr = base64.RawStdEncoding.Decode(array, self.buf[:loc])
	if goErr != nil {
		return &errs.EncodedArrayCorrupted
	}

	self.FindEndToken()
	return
}

// DecodeByteArrayAsNumber convert number array to [n]byte
func (self *Decoder[BUF]) DecodeByteArrayAsNumber(array []byte) (err error_p.Error) {
	if self.CheckNullValue() {
		return
	}

	var value uint8
	for i := 0; i < len(array); i++ {
		value, err = self.Decode_Integer_U8()
		if err != nil {
			// err = &errs.EncodedArrayCorrupted
			return
		}
		array[i] = value
		self.FindEndToken()
	}
	if self.token != ']' {
		// err = &errs.EncodedArrayCorrupted
	}
	return
}

/*
	Slice as Number
*/

// DecodeByteSliceAsNumber convert number string slice to []byte. pass buf start from after [ and receive from after ]
func (self *Decoder[BUF]) DecodeByteSliceAsNumber() (slice []byte, err error_p.Error) {
	slice = make([]byte, 0, 8) // TODO::: Is cap efficient enough?

	var num uint8
	for !self.CheckToken(']') {
		num, err = self.Decode_Integer_U8()
		if err != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, num)

		self.FindEndToken()
	}
	return
}

// Decode_Integer_U16SliceAsNumber convert uint16 number string slice to []byte. pass buf start from after [ and receive from after ]
func (self *Decoder[BUF]) Decode_Integer_U16SliceAsNumber() (slice []uint16, err error_p.Error) {
	slice = make([]uint16, 0, 8) // TODO::: Is cap efficient enough?

	var num integer.U16
	for !self.CheckToken(']') {
		num, err = self.Decode_Integer_U16()
		if err != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, uint16(num))

		self.FindEndToken()
	}
	return
}

// Decode_Integer_U32SliceAsNumber convert uint32 number string slice to []byte. pass buf start from after [ and receive from after ]
func (self *Decoder[BUF]) Decode_Integer_U32SliceAsNumber() (slice []uint32, err error_p.Error) {
	slice = make([]uint32, 0, 8) // TODO::: Is cap efficient enough?

	var num uint32
	for !self.CheckToken(']') {
		num, err = self.Decode_Integer_U32()
		if err != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, num)

		self.FindEndToken()
	}
	return
}

// Decode_Integer_U64SliceAsNumber convert uint64 number string slice to []byte. pass buf start from after [ and receive from after ]
func (self *Decoder[BUF]) Decode_Integer_U64SliceAsNumber() (slice []uint64, err error_p.Error) {
	slice = make([]uint64, 0, 8) // TODO::: Is cap efficient enough?

	var num uint64
	for !self.CheckToken(']') {
		num, err = self.Decode_Integer_U64()
		if err != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, num)

		self.FindEndToken()
	}
	return
}

/*
	Slice as Base64
*/

// DecodeByteSliceAsBase64 convert base64 string to []byte
func (self *Decoder[BUF]) DecodeByteSliceAsBase64() (slice []byte, err error_p.Error) {
	self.TrimToStringStart()
	var loc, _ = self.buf.Index('"')
	if loc < 0 {
		// err = &errs.EncodedSliceCorrupted
		return
	}

	slice = make([]byte, base64.RawStdEncoding.DecodedLen(len(self.buf[:loc])))
	var n int
	var goErr error
	n, goErr = base64.RawStdEncoding.Decode(slice, self.buf[:loc])
	if goErr != nil {
		return slice, &errs.EncodedSliceCorrupted
	}
	slice = slice[:n]

	self.FindEndToken()
	return
}

// Decode32ByteArraySliceAsBase64 decode [32]byte base64 string slice. pass buf start from after [ and receive from after ]
func (self *Decoder[BUF]) Decode32ByteArraySliceAsBase64() (slice [][32]byte, err error_p.Error) {
	const base64Len = 43 // base64.RawStdEncoding.EncodedLen(len(32))	>>	(32*8 + 5) / 6
	slice = make([][32]byte, 0, 8)

	var openBracketLoc, _ = self.buf.Index('[')
	if openBracketLoc < 0 {
		// err = &errs.EncodedSliceCorrupted
		return
	}
	self.buf = self.buf[openBracketLoc+1:]
	var goErr error
	var array [32]byte
	for !self.CheckToken(']') {
		self.TrimToStringStart()
		_, goErr = base64.RawStdEncoding.Decode(array[:], self.buf[:base64Len])
		if goErr != nil {
			// err = &errs.EncodedSliceCorrupted
			return
		}
		slice = append(slice, array)
		self.buf = self.buf[base64Len:]
		self.FindEndToken()
	}
	return
}
