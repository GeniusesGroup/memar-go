/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"bytes"

	"libgo/convert"
	"libgo/protocol"
)

// DecoderUnsafe store data to decode data by each method!
type DecoderUnsafe struct {
	Decoder
}

// DecodeString return string. pass d.buf start from after " and receive from from after "
func (d *DecoderUnsafe) DecodeString() (s string, err protocol.Error) {
	if d.CheckNullValue() {
		return
	}

	var loc = bytes.IndexByte(d.buf, '"')
	d.buf = d.buf[loc+1:] // remove any byte before first " due to don't need them

	loc = bytes.IndexByte(d.buf, '"')
	if loc < 0 {
		err = &ErrEncodedStringCorrupted
		return
	}

	var slice []byte = d.buf[:loc]

	d.buf = d.buf[loc+1:]
	s = convert.UnsafeByteSliceToString(slice)
	return
}
