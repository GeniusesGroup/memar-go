/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"memar/computer/buffer/byteslice/convert"
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
)

// DecoderUnsafe store data to decode data by each method!
type DecoderUnsafe[BUF buffer_p.Buffer] struct {
	Decoder[BUF]
}

// DecodeString return string. pass d.buf start from after " and receive from from after "
func (d *DecoderUnsafe[BUF]) DecodeString() (s string, err error_p.Error) {
	if d.CheckNullValue() {
		return
	}

	var loc, _ = d.buf.Index('"')
	d.buf = d.buf[loc+1:] // remove any byte before first " due to don't need them

	loc, _ = d.buf.Index('"')
	if loc < 0 {
		// err = &json_errs.EncodedStringCorrupted
		return
	}

	var slice []byte = d.buf[:loc]

	d.buf = d.buf[loc+1:]
	s = convert.UnsafeByteSliceToString(slice)
	return
}
