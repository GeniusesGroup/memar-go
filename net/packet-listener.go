/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	error_p "memar/process/error/protocol"
	net_p "memar/net/protocol"
)

type PacketListener struct {
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (pl *PacketListener) Init() (err error_p.Error) {
	// TODO:::
	return
}
func (pl *PacketListener) Reinit() (err error_p.Error) {
	// TODO:::
	return
}
func (pl *PacketListener) Deinit() (err error_p.Error) {
	// TODO:::
	return
}

func (pl *PacketListener) NetworkPacketHandler(np net_p.Packet) {
	var err = HandleFrames(np)
	if err != nil {
		// TODO:::
	}
}
