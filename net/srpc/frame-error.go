/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/protocol"
	"memar/syllab"
)

/*
	type ErrorFrame struct {
		StreamID uint64
		ErrorID  uint64
	}
*/
type ErrorFrame []byte

func (f ErrorFrame) StreamID() uint64 { return syllab.GetUInt64(f, 0) }
func (f ErrorFrame) ErrorID() uint64  { return syllab.GetUInt64(f, 8) }

//memar:impl memar/protocol.Network_Frame
func (f ErrorFrame) NextFrame() []byte { return f[16:] }

func (f ErrorFrame) Do(sk protocol.Socket) (err protocol.Error) {
	var al = sk.ApplicationLayer()
	if al == nil {
		// conn.StreamFailed()
		// Send response or just ignore stream
		// TODO::: DDOS!!??
		return
	}
	var peerErrorID uint64 = f.ErrorID()
	var peerError = protocol.App.GetErrorByID(protocol.ID(peerErrorID))
	al.SetError(peerError)
	return
}
