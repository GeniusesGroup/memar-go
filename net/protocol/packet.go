/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	buffer_p "memar/buffer/protocol"
	error_p "memar/error/protocol"
)

// Receiver must release Packet and don't use it after return.
// So almost in most cases dev must copy Packet payload to the socket.
type Packet = buffer_p.Buffer

type PacketTarget interface {
	// TODO::: just accept FrameType? other conditions? some thing like Regex??
	AddPacketListener(fID FrameType, callback PacketListener) (err error_p.Error)
	RemovePacketListener(fID FrameType, callback PacketListener) (err error_p.Error)
}

// PacketListener
// It isn't just for hardware network packet, It is use for any requirements e.g. IPC, D-Bus, ...
type PacketListener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	NetworkPacketHandler(np Packet)
}
