/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	buffer_p "memar/computer/buffer/protocol"
	"memar/computer/buffer/byteslice/convert"
	error_p "memar/process/error/protocol"
)

// DecoderUnsafeMinified store data to decode data by each method.
type DecoderUnsafeMinified[BUF buffer_p.Buffer] struct {
	DecoderMinified[BUF]
}

// DecodeKey return json key. pass d.buf start from after {||, and receive from after :
func (d *DecoderUnsafeMinified[BUF]) DecodeKey() string {
	d.Offset(2)
	var loc, _ = d.buf.Index('"')
	var slice []byte = d.buf[:loc]
	d.Offset(loc + 2) // +2 due to have '":' after key name end.
	return convert.UnsafeByteSliceToString(slice)
}

// DecodeString return string. pass d.buf start from after " and receive from from after "
func (d *DecoderUnsafeMinified[BUF]) DecodeString() (s string, err error_p.Error) {
	d.Offset(1) // due to have " at start

	var loc = d.buf.Index('"')
	if loc < 0 {
		// err = &errs.EncodedStringCorrupted
		return
	}

	var slice []byte = d.buf[:loc]
	d.Offset(loc + 1)
	s = convert.UnsafeByteSliceToString(slice)
	return
}
