/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"

	"../convert"
	"../protocol"
)

// DecoderUnsafeMinifed store data to decode data by each method!
type DecoderUnsafeMinifed struct {
	DecoderMinifed
}

// DecodeString return string. pass d.Buf start from after " and receive from from after "
func (d *DecoderUnsafeMinifed) DecodeString() (s string, err protocol.Error) {
	d.Offset(1) // due to have " at start

	var loc = bytes.IndexByte(d.Buf, '"')
	if loc < 0 {
		err = ErrEncodedStringCorrupted
		return
	}

	var slice []byte = d.Buf[:loc]
	d.Offset(loc + 1)
	s = convert.UnsafeByteSliceToString(slice)
	return
}
