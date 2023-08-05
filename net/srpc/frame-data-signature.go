/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/protocol"
	"memar/syllab"
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

func (f DataSignatureFrame) ID() int64 { return syllab.GetInt64(f, 0) }

//memar:impl memar/protocol.Network_Frame
func (f DataSignatureFrame) NextFrame() []byte { return f[8:] }

func (f DataSignatureFrame) Do(sk protocol.Socket) (err protocol.Error) {
	return
}
