/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/codec/data_exchange/syllab"
	datatype_p "memar/computer/datatype/protocol"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
	"memar/process/errors"
)

/*
	type ErrorFrame struct {
		StreamID uint64
		ErrorID  uint64
	}
*/
type ErrorFrame []byte

func (self ErrorFrame) StreamID() uint64 { return syllab.GetUInt64(self, 0) }
func (self ErrorFrame) ErrorID() uint64  { return syllab.GetUInt64(self, 8) }

//memar:impl memar/protocol.Network_Frame
func (self ErrorFrame) NextFrame() []byte { return self[16:] }

func (self ErrorFrame) Do(sk net_p.Socket) (err error_p.Error) {
	var al = sk.OSI_ApplicationLayer()
	if al == nil {
		// conn.StreamFailed()
		// Send response or just ignore stream
		// TODO::: DDOS!!??
		return
	}
	var peerErrorID uint64 = self.ErrorID()
	var peerError = errors.GetByID(datatype_p.ID(peerErrorID))
	al.SetError(peerError)
	return
}
