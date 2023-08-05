/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/protocol"
	"memar/syllab"
)

/*
callService use to call a service without need to open any stream.
It can also use when service request data is smaller than network MTU.
Or use for time sensitive data like audio and video that streams shape in app layer

	type ServiceFrame struct {
		Length     uint16 // including the header fields
		ServiceID  uint64
		CompressID uint64
		Time       int64 // It is used to match the request and response and drop if TTL
		Payload    []byte
	}
*/
type ServiceFrame []byte

func (f ServiceFrame) Length() uint16     { return syllab.GetUInt16(f, 0) }
func (f ServiceFrame) ServiceID() uint64  { return syllab.GetUInt64(f, 2) }
func (f ServiceFrame) CompressID() uint64 { return syllab.GetUInt64(f, 10) }
func (f ServiceFrame) Time() int64        { return syllab.GetInt64(f, 18) }
func (f ServiceFrame) Payload() []byte    { return f[26:f.Length()] }

//memar:impl memar/protocol.Network_Frame
func (f ServiceFrame) NextFrame() []byte { return f[f.Length():] }

func (f ServiceFrame) Do(sk protocol.Socket) (err protocol.Error) {
	// var serviceID uint32 = ServiceFrame.ServiceID()
	// TODO:::
	return
}
