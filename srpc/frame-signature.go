/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../protocol"
	"../syllab"
)

/*
type signatureFrame struct {
	Length    [2]byte // including the header fields
	StreamID  [4]byte // uint32
	Signature []byte
}
*/
type signatureFrame []byte

func (f signatureFrame) ID() int64         { return syllab.GetInt64(f, 0) }
func (f signatureFrame) NextFrame() []byte { return f[8:] }

func registerStreamSignature(conn protocol.Connection, frame signatureFrame) (err protocol.Error) {
	return
}
