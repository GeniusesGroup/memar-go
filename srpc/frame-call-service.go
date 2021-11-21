/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../protocol"
	"../syllab"
)

/*
type serviceFrame struct {
	Length     [2]byte // including the header fields
	ServiceID  uint64
	CompressID uint64
	Time       int64 // It is used to match the request and response and drop if TTL
	Payload    []byte
}
*/
type serviceFrame []byte

func (f serviceFrame) Length() uint16     { return syllab.GetUInt16(f, 0) }
func (f serviceFrame) ServiceID() uint64  { return syllab.GetUInt64(f, 2) }
func (f serviceFrame) CompressID() uint64 { return syllab.GetUInt64(f, 10) }
func (f serviceFrame) Time() int64        { return syllab.GetInt64(f, 18) }
func (f serviceFrame) Payload() []byte    { return f[26:f.Length()] }
func (f serviceFrame) NextFrame() []byte  { return f[f.Length():] }

// callService use to call a service without need to open any stream.
// It can also use when service request data is smaller than network MTU.
// Or use for time sensitive data like audio and video that streams shape in app layer
func callService(conn protocol.Connection, frame serviceFrame) (err protocol.Error) {
	// var serviceID uint32 = serviceFrame.ServiceID()
	// TODO:::
	return
}
