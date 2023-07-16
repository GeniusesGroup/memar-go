/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Receiver must release Network_Packet slice and don't use it after return.
// So almost in most cases dev must copy Network_Packet payload to the socket.
type Network_Packet []byte

type PacketTarget interface {
	// TODO::: just accept Network_FrameID? other conditions? some thing like Regex??
	AddPacketListener(fID Network_FrameID, callback Network_PacketListener) (err Error)
	RemovePacketListener(fID Network_FrameID, callback Network_PacketListener) (err Error)
}

// Network_PacketListener
// It isn't just for hardware network packet, It is use for any requirements e.g. IPC, D-Bus, ...
type Network_PacketListener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	NetworkPacketHandler(np Network_Packet)
}
