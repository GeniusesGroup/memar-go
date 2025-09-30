/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"bytes"

	"libgo/convert"
	"libgo/protocol"
)

// DecoderUnsafeMinified store data to decode data by each method.
type DecoderUnsafeMinified struct {
	DecoderMinified
}

// DecodeString return string. pass d.buf start from after " and receive from from after "
func (d *DecoderUnsafeMinified) DecodeString() (s string, err protocol.Error) {
	d.Offset(1) // due to have " at start

	var loc = bytes.IndexByte(d.buf, '"')
	if loc < 0 {
		err = &ErrEncodedStringCorrupted
		return
	}

	var slice []byte = d.buf[:loc]
	d.Offset(loc + 1)
	s = convert.UnsafeByteSliceToString(slice)
	return
}
