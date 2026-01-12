/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/codec/data_exchange/syllab"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

/*
callService use to call a service without need to open any stream.
It can also use when service request data is smaller than network MTU.
Or use for time sensitive data like audio and video that streams shape in app layer

	type ServiceFrame struct {
		FrameLength uint16 // including the header fields
		ServiceID   uint64
		CompressID  uint64
		Time        int64 // It is used to match the request and response and drop iself TTL
		Payload     []byte
	}
*/
type ServiceFrame []byte

func (self ServiceFrame) FrameLength() uint16 { return syllab.GetUInt16(self, 0) }
func (self ServiceFrame) ServiceID() uint64   { return syllab.GetUInt64(self, 2) }
func (self ServiceFrame) CompressID() uint64  { return syllab.GetUInt64(self, 10) }
func (self ServiceFrame) Time() int64         { return syllab.GetInt64(self, 18) }
func (self ServiceFrame) Payload() []byte     { return self[26:self.FrameLength()] }

//memar:impl memar/protocol.Network_Frame
func (self ServiceFrame) NextFrame() []byte { return self[self.FrameLength():] }

func (self ServiceFrame) Do(sk net_p.Socket) (err error_p.Error) {
	// var serviceID uint32 = ServiceFrame.ServiceID()
	// TODO:::
	return
}
