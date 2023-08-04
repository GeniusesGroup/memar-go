/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/protocol"
)

type PacketTarget struct {
}

//memar:impl memar/protocol.ObjectLifeCycle
func (pt *PacketTarget) Init() (err protocol.Error) {
	// TODO:::
	return
}
func (pt *PacketTarget) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (pt *PacketTarget) Deinit() (err protocol.Error) {
	// TODO:::
	return
}

//memar:impl memar/protocol.PacketTarget
func (pt *PacketTarget) AddPacketListener(fID protocol.Network_FrameType, callback protocol.Network_PacketListener) (err protocol.Error) {
	// TODO:::
	return
}

func (pt *PacketTarget) RemovePacketListener(fID protocol.Network_FrameType, callback protocol.Network_PacketListener) (err protocol.Error) {
	// TODO:::
	return
}
