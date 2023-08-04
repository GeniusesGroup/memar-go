/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/protocol"
)

/*
	type frame struct {
		Type     byte
		Payload []byte
	}
*/
type frame []byte

//memar:impl memar/protocol.Network_Framer
func (f frame) FrameType() protocol.Network_FrameType { return protocol.Network_FrameType(f[0]) }

func (f frame) Payload() []byte { return f[1:] }
