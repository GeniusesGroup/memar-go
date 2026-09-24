/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	error_p "memar/process/error/protocol"
	net_p "memar/net/protocol"
)

type PacketTarget struct {
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (pt *PacketTarget) Init() (err error_p.Error) {
	// TODO:::
	return
}
func (pt *PacketTarget) Reinit() (err error_p.Error) {
	// TODO:::
	return
}
func (pt *PacketTarget) Deinit() (err error_p.Error) {
	// TODO:::
	return
}

//memar:impl memar/net/protocol.PacketTarget
func (pt *PacketTarget) AddPacketListener(fID net_p.FrameType, callback net_p.PacketListener) (err error_p.Error) {
	// TODO:::
	return
}

func (pt *PacketTarget) RemovePacketListener(fID net_p.FrameType, callback net_p.PacketListener) (err error_p.Error) {
	// TODO:::
	return
}
