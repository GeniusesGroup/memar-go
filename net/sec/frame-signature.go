/* For license and copyright information please see the LEGAL file in the code repository */

package sec

import (
	"memar/codec/syllab"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
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

func (f SignatureFrame) Do(sk net_p.Socket) (err error_p.Error) {
	return
}

func checkSignature() {}
