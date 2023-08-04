/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/protocol"
)

type PacketListener struct {
}

//memar:impl memar/protocol.ObjectLifeCycle
func (pl *PacketListener) Init() (err protocol.Error) {
	// TODO:::
	return
}
func (pl *PacketListener) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (pl *PacketListener) Deinit() (err protocol.Error) {
	// TODO:::
	return
}

func (pl *PacketListener) NetworkPacketHandler(np protocol.Network_Packet) {
	var err = HandleFrames(np)
	if err != nil {
		// TODO:::
	}
}
