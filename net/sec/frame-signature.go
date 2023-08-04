/* For license and copyright information please see the LEGAL file in the code repository */

package sec

import (
	"memar/protocol"
	"memar/syllab"
)

/*
registerStreamSignature

	type SignatureFrame struct {
		Length    [2]byte // including the header fields
		StreamID  [4]byte // uint32
		Signature []byte  // Checksum, MAC, Tag, ...
	}
*/
type SignatureFrame []byte

func (f SignatureFrame) ID() int64 { return syllab.GetInt64(f, 0) }

//memar:impl memar/protocol.Network_Frame
func (f SignatureFrame) NextFrame() []byte { return f[8:] }

func (f SignatureFrame) Do(sk protocol.Socket) (err protocol.Error) {
	return
}

func checkSignature() {}
