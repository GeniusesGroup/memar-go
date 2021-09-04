/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"bytes"
	"io"

	"../protocol"
)

// Request represent request protocol structure!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/sRPC.md
// type Request struct {
// 	  ServiceID uint64
// 	  Payload   []byte
// }
// Due to improve performance we use simple byte slice against above structure!
type Request []byte

// NewRequest make and return the new request!
func NewRequest(payloadLength uint64) (req Request) {
	req = make([]byte, MinLength+payloadLength)
	return
}

// Check check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
// Anyway expectedMinLen can't be under MinLength!
func (r Request) Check(expectedMinLen uint64) protocol.Error {
	if uint64(len(r)) < expectedMinLen {
		return ErrPacketTooShort
	}
	return nil
}

// ServiceID returns service ID of the request.
func (r Request) ServiceID() uint64 {
	return uint64(r[0]) | uint64(r[1])<<8 | uint64(r[2])<<16 | uint64(r[3])<<24 | uint64(r[4])<<32 | uint64(r[5])<<40 | uint64(r[6])<<48 | uint64(r[7])<<56
}

// SetServiceID encodes service ID to the request.
func (r Request) SetServiceID(id uint64) {
	r[0] = byte(id)
	r[1] = byte(id >> 8)
	r[2] = byte(id >> 16)
	r[3] = byte(id >> 24)
	r[4] = byte(id >> 32)
	r[5] = byte(id >> 40)
	r[6] = byte(id >> 48)
	r[7] = byte(id >> 56)
}

// Payload return payload of the request
func (r Request) Payload() []byte {
	return r[MinLength:]
}

/*
********** protocol.Codec interface **********
 */

func (r Request) MediaType() string    { return "application/srpc" }
func (r Request) CompressType() string { return "" }

func (r Request) Decode(buf protocol.Buffer) (err protocol.Error) {
	// TODO:::
	// use simple binding as Request(buf.Get())
	return
}

// Encode write compressed data to given buf.
func (r Request) Encode(buf protocol.Buffer) {
	buf.Set(r)
}

// Marshal ...
func (r Request) Marshal() (data []byte) {
	return r
}

// Marshal ...
func (r Request) MarshalTo(data []byte) []byte {
	data = append(data, r...)
	return data
}

// UnMarshal ...
func (r Request) UnMarshal(data []byte) (err protocol.Error) {
	return
}

// ReadFrom decodes r Raw data by read from given io.Reader!
func (r Request) ReadFrom(reader io.Reader) (n int64, err error) {
	var buf bytes.Buffer
	n, err = io.Copy(&buf, reader)
	r = buf.Bytes()
	return
}

// WriteTo enecodes r Raw data and write it to given io.Writer!
func (r Request) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	writeLen, err = w.Write(r)
	totalWrite = int64(writeLen)
	return
}

// Len return the len of the request
func (r Request) Len() int {
	return len(r)
}
