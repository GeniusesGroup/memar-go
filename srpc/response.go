/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../giti"
)

// Response is represent response protocol structure!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/sRPC.md
// type Response struct {
// 	  ErrorID uint64 // 0 means no error!
// 	  Payload []byte
// }
// Due to improve performance we use simple byte slice against above structure!
type Response []byte

// MakeNewResponse make new response!
func MakeNewResponse(errorID uint64, payloadLength int) (res Response) {
	res = make([]byte, MinLength+payloadLength)
	res.SetErrorID(errorID)
	return
}

// Check will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
// Anyway expectedMinLen can't be under MinLength!
func (r Response) Check(expectedMinLen int) giti.Error {
	if len(r) < expectedMinLen {
		return ErrPacketTooShort
	}
	return nil
}

// ErrorID returns error ID of the response.
func (r Response) ErrorID() uint64 {
	return uint64(r[0]) | uint64(r[1])<<8 | uint64(r[2])<<16 | uint64(r[3])<<24 | uint64(r[4])<<32 | uint64(r[5])<<40 | uint64(r[6])<<48 | uint64(r[7])<<56
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
func (r Response) SetError(err giti.Error) {
	var errID = err.ID()
	r.SetErrorID(errID)
}

// Payload return payload of the response
func (r Response) Payload() []byte {
	return r[MinLength:]
}

// Len return the len of the response
func (r Response) Len() int {
	return len(r)
}
