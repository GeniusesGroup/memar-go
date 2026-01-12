/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/codec/data_exchange/syllab"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

/*
registerStreamSignature

	type DataSignatureFrame struct {
		Length    [2]byte // including the header fields
		StreamID  [4]byte // uint32
		Signature []byte  // Checksum, MAC, Tag, ...
	}
*/
type DataSignatureFrame []byte

func (self DataSignatureFrame) ID() int64 { return syllab.GetInt64(self, 0) }

//memar:impl memar/protocol.Network_Frame
func (self DataSignatureFrame) NextFrame() []byte { return self[8:] }

func (self DataSignatureFrame) Do(sk net_p.Socket) (err error_p.Error) {
	return
}
