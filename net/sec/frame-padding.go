/* For license and copyright information please see the LEGAL file in the code repository */

package sec

import (
	"memar/binary"
)

/*
	type PaddingFrame struct {
		Length  [2]byte // including the header fields
		Padding []byte
	}
*/
type PaddingFrame []byte

func (f PaddingFrame) Length() uint16  { return binary.BigEndian(f[0:]).Uint16() }
func (f PaddingFrame) Payload() []byte { return f[2:f.Length()] }

//memar:impl memar/protocol.Network_Frame
func (f PaddingFrame) NextFrame() []byte { return f[f.Length():] }
