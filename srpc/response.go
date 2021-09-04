/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"bytes"
	"io"

	"../protocol"
)

// Response represent response protocol structure!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/sRPC.md
// type Response struct {
// 	  ErrorID uint64 // 0 means no error!
// 	  Payload []byte
// }
// Due to improve performance we use simple byte slice against above structure!
type Response []byte

// NewResponse make and return the new response!
func NewResponse(payloadLength uint64) (res Response) {
	res = make([]byte, MinLength+payloadLength)
	return
}

// Check will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
// Anyway expectedMinLen can't be under MinLength!
func (r Response) Check(expectedMinLen uint64) protocol.Error {
	if uint64(len(r)) < expectedMinLen {
		return ErrPacketTooShort
	}
	return nil
}

// ErrorID returns error ID of the response.
func (r Response) ErrorID() uint64 {
	return uint64(r[0]) | uint64(r[1])<<8 | uint64(r[2])<<16 | uint64(r[3])<<24 | uint64(r[4])<<32 | uint64(r[5])<<40 | uint64(r[6])<<48 | uint64(r[7])<<56
}

// Error returns error of the response.
func (r Response) Error() (err protocol.Error) {
	var errID = r.ErrorID()
	err = protocol.App.GetErrorByID(errID)
	return
}

// SetErrorID encodes error ID to the response.
func (r Response) SetErrorID(id uint64) {
	r[0] = byte(id)
	r[1] = byte(id >> 8)
	r[2] = byte(id >> 16)
	r[3] = byte(id >> 24)
	r[4] = byte(id >> 32)
	r[5] = byte(id >> 40)
	r[6] = byte(id >> 48)
	r[7] = byte(id >> 56)
}

// SetError encodes error ID to the response.
func (r Response) SetError(err protocol.Error) {
	var errID = err.URN().ID()
	r.SetErrorID(errID)
}

// Payload return payload of the response
func (r Response) Payload() []byte {
	return r[MinLength:]
}

/*
********** protocol.Codec interface **********
 */

func (r Response) MediaType() string    { return "application/srpc" }
func (r Response) CompressType() string { return "" }

func (r Response) Decode(buf protocol.Buffer) (err protocol.Error) {
	// TODO:::
	// use simple binding as Response(buf.Get())
	return
}

// Encode write compressed data to given buf.
func (r Response) Encode(buf protocol.Buffer) {
	buf.Set(r)
}

// Marshal ...
func (r Response) Marshal() (data []byte) {
	return r
}

// Marshal ...
func (r Response) MarshalTo(data []byte) []byte {
	data = append(data, r...)
	return data
}

// UnMarshal ...
func (r Response) UnMarshal(data []byte) (err protocol.Error) {
	return
}

// ReadFrom decodes r Raw data by read from given io.Reader!
func (r Response) ReadFrom(reader io.Reader) (n int64, err error) {
	var buf bytes.Buffer
	n, err = io.Copy(&buf, reader)
	r = buf.Bytes()
	return
}

// WriteTo enecodes r Raw data and write it to given io.Writer!
func (r Response) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	writeLen, err = w.Write(r)
	totalWrite = int64(writeLen)
	return
}

// Len return the len of the request
func (r Response) Len() int {
	return len(r)
}
