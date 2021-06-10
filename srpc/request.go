/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../giti"
)

// Request is represent request protocol structure!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/sRPC.md
// type Request struct {
// 	  ServiceID uint64
// 	  Payload   []byte
// }
// Due to improve performance we use simple byte slice against above structure!
type Request []byte

// MakeNewRequest make new request!
func MakeNewRequest(serviceID uint64, payloadLength int) (req Request) {
	req = make([]byte, MinLength+payloadLength)
	req.SetServiceID(serviceID)
	return
}

// Check check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
// Anyway expectedMinLen can't be under MinLength!
func (r Request) Check(expectedMinLen int) giti.Error {
	if len(r) < expectedMinLen {
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

// Len return the len of the request
func (r Request) Len() int {
	return len(r)
}
