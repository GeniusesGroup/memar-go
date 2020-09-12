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

// EncodeString append given string as given format
func (e *Encoder) EncodeString(s string) {
	e.Buf = append(e.Buf, s...)
}

// EncodeSliceAsNumber append given byte slice as number string
func (e *Encoder) EncodeSliceAsNumber(slice []byte) {
	var ln = len(slice)
	for i := 0; i < ln; i++ {
		e.Buf = strconv.AppendUint(e.Buf, uint64(slice[i]), 10)
		e.Buf = append(e.Buf, ',')
	}
	e.Buf = e.Buf[:len(e.Buf)-1] // remove trailing comma
}

// EncodeSliceAsBase64 use to append []byte as base64 string
func (e *Encoder) EncodeSliceAsBase64(slice []byte, base64Len int){
	var ln = len(e.Buf)
	e.Buf = e.Buf[:ln+base64Len]
	base64.StdEncoding.Encode(e.Buf[ln:], slice)
}
