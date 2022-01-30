/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"

	"../convert"
	"../protocol"
)

// DecoderUnsafe store data to decode data by each method!
type DecoderUnsafe struct {
	Decoder
}

// DecodeString return string. pass d.Buf start from after " and receive from from after "
func (d *DecoderUnsafe) DecodeString() (s string, err protocol.Error) {
	if d.CheckNullValue() {
		return
	}

	var loc = bytes.IndexByte(d.Buf, '"')
	d.Buf = d.Buf[loc+1:] // remove any byte before first " due to don't need them

	loc = bytes.IndexByte(d.Buf, '"')
	if loc < 0 {
		err = ErrEncodedStringCorrupted
		return
	}

	var slice []byte = d.Buf[:loc]

	d.Buf = d.Buf[loc+1:]
	s = convert.UnsafeByteSliceToString(slice)
	return
}
