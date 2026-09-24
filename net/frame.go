/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	net_p "memar/net/protocol"
)

/*
	type frame struct {
		Type     byte
		Payload []byte
	}
*/
type frame []byte

//memar:impl memar/net/protocol.Field_FrameType
func (f frame) FrameType() net_p.FrameType { return net_p.FrameType(f[0]) }

func (f frame) Payload() []byte { return f[1:] }
