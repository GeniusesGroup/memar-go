/* For license and copyright information please see the LEGAL file in the code repository */

package connection

import (
	"sync/atomic"

	"libgo/protocol"
	"libgo/time/unix"
)

// Metric store the connection metric data and implement protocol.ConnectionMetrics
type Metric struct {
	lastUsage                   unix.Atomic   // Last use of this connection
	maxBandwidth                atomic.Uint64 // Byte/Second and Connection can limit to a fixed number
	bytesSent                   atomic.Uint64 // Counts the bytes of packets sent.
	packetsSent                 atomic.Uint64 // Counts sent packets.
	bytesReceived               atomic.Uint64 // Counts the bytes of packets receive.
	packetsReceived             atomic.Uint64 // Counts received packets.
	lostPackets                 atomic.Uint64 // Counts any lost packet that peer request to resend it
	lostBytes                   atomic.Uint64 //
	resendPackets               atomic.Uint64 // Counts any duplicate packet that not request to resend it, use to prevent attacks
	resendBytes                 atomic.Uint64 //
	failedPacketsSent           atomic.Uint64 //
	failedPacketsReceived       atomic.Uint64 // Counts failed packets receive for fire-walling server from some attack types
	notRequestedPacketsReceived atomic.Uint64 // Counts not requested packets received for fire-walling server from some attack types
	succeedStreamCount          atomic.Uint64 // Count successful request.
	failedStreamCount           atomic.Uint64 // Count failed services call e.g. data validation failed, ...
}

//libgo:impl libgo/protocol.ConnectionMetrics
func (m *Metric) LastUsage() protocol.Time            { return &m.lastUsage }
func (m *Metric) MaxBandwidth() uint64                { return m.maxBandwidth.Load() }
func (m *Metric) BytesSent() uint64                   { return m.bytesSent.Load() }
func (m *Metric) PacketsSent() uint64                 { return m.packetsSent.Load() }
func (m *Metric) BytesReceived() uint64               { return m.bytesReceived.Load() }
func (m *Metric) PacketsReceived() uint64             { return m.packetsReceived.Load() }
func (m *Metric) LostPackets() uint64                 { return m.lostPackets.Load() }
func (m *Metric) LostBytes() uint64                   { return m.lostBytes.Load() }
func (m *Metric) ResendPackets() uint64               { return m.resendPackets.Load() }
func (m *Metric) ResendBytes() uint64                 { return m.resendBytes.Load() }
func (m *Metric) FailedPacketsSent() uint64           { return m.failedPacketsSent.Load() }
func (m *Metric) FailedPacketsReceived() uint64       { return m.failedPacketsReceived.Load() }
func (m *Metric) NotRequestedPacketsReceived() uint64 { return m.notRequestedPacketsReceived.Load() }
func (m *Metric) SucceedStreamCount() uint64          { return m.succeedStreamCount.Load() }
func (m *Metric) FailedStreamCount() uint64           { return m.failedStreamCount.Load() }

// StreamSucceed store successful service call occur on this connection
func (m *Metric) StreamSucceed() {
	m.lastUsage.Now()
	m.succeedStreamCount.Add(1)
	// TODO::: Any other job??
}

// StreamFailed store failed service request occur on this connection!
// Base on the connection it can other action to prevent any attack! e.g. tel router to block
func (m *Metric) StreamFailed() {
	m.lastUsage.Now()
	m.failedStreamCount.Add(1)
	// TODO::: Any other job??
}

func (m *Metric) PacketReceived(packetLength uint64) {
	m.lastUsage.Now()
	m.packetsReceived.Add(1)
	m.bytesReceived.Add(packetLength)
}

func (m *Metric) DuplicatePacketReceived(packetLength uint64) {
	m.lastUsage.Now()
	m.resendPackets.Add(1)
	m.resendBytes.Add(packetLength)
}

func (m *Metric) PacketSent(packetLength uint64) {
	m.lastUsage.Now()
	m.packetsSent.Add(1)
	m.bytesSent.Add(packetLength)
}

func (m *Metric) PacketResend(packetLength uint64) {
	m.PacketSent(packetLength)
	m.lostPackets.Add(1)
	m.lostBytes.Add(packetLength)
}
