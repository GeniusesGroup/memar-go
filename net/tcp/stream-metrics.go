/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/protocol"
)

type StreamMetrics struct {
	TotalPacket     uint32 // Expected packets count that must received!
	PacketReceived  uint32 // Count of packets received!
	LastPacketID    uint32 // Last send or received Packet use to know order of packets!
	PacketDropCount uint8  // Count drop packets to prevent some attacks type!
}

//memar:impl memar/protocol.ObjectLifeCycle
func (sm *StreamMetrics) Init(timeout protocol.Duration) (err protocol.Error) {
	return
}
func (sm *StreamMetrics) Reinit() (err protocol.Error) {
	// TODO:::
	return
}
func (sm *StreamMetrics) Deinit() (err protocol.Error) {
	return
}
