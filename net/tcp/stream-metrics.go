/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	error_p "memar/error/protocol"
	"memar/time/duration"
)

type StreamMetrics struct {
	TotalPacket     uint32 // Expected packets count that must received!
	PacketReceived  uint32 // Count of packets received!
	LastPacketID    uint32 // Last send or received Packet use to know order of packets!
	PacketDropCount uint8  // Count drop packets to prevent some attacks type!
}

// memar/computer/language/object/protocol.LifeCycle
func (sm *StreamMetrics) Init(timeout duration.NanoSecond) (err error_p.Error) {
	return
}
func (sm *StreamMetrics) Reinit() (err error_p.Error) {
	// TODO:::
	return
}
func (sm *StreamMetrics) Deinit() (err error_p.Error) {
	return
}
