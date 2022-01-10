/* For license and copyright information please see LEGAL file in repository */

package connection

import (
	"sync/atomic"
	"time"

	"../protocol"
)

// Metric store the connection metric data and impelement protocol.ConnectionMetrics
type Metric struct {
	lastUsage                   int64  // Last use of this connection
	maxBandwidth                uint64 // Byte/Second and Connection can limit to a fixed number
	bytesSent                   uint64 // Counts the bytes of packets sent.
	packetsSent                 uint64 // Counts sent packets.
	bytesReceived               uint64 // Counts the bytes of packets receive.
	packetsReceived             uint64 // Counts received packets.
	failedPacketsReceived       uint64 // Counts failed packets receive for firewalling server from some attack types!
	notRequestedPacketsReceived uint64 // Counts not requested packets received for firewalling server from some attack types!
	succeedStreamCount          uint64 // Count successful request.
	failedStreamCount           uint64 // Count failed services call e.g. data validation failed, ...
}

func (m *Metric) LastUsage() protocol.TimeUnixMilli   { return protocol.TimeUnixMilli(m.lastUsage) }
func (m *Metric) MaxBandwidth() uint64                { return m.maxBandwidth }
func (m *Metric) BytesSent() uint64                   { return m.bytesSent }
func (m *Metric) PacketsSent() uint64                 { return m.packetsSent }
func (m *Metric) BytesReceived() uint64               { return m.bytesReceived }
func (m *Metric) PacketsReceived() uint64             { return m.packetsReceived }
func (m *Metric) FailedPacketsReceived() uint64       { return m.failedPacketsReceived }
func (m *Metric) NotRequestedPacketsReceived() uint64 { return m.notRequestedPacketsReceived }
func (m *Metric) SucceedStreamCount() uint64          { return m.succeedStreamCount }
func (m *Metric) FailedStreamCount() uint64           { return m.failedStreamCount }

// StreamSucceed store successfull service call occur on this connection
func (m *Metric) StreamSucceed() {
	atomic.StoreInt64(&m.lastUsage, time.Now().Unix())
	atomic.AddUint64(&m.succeedStreamCount, 1)
	// TODO::: Any other job??
}

// StreamFailed store failed service request occur on this connection!
// Base on the connection it can other action to prevent any attack! e.g. tel router to block
func (m *Metric) StreamFailed() {
	atomic.StoreInt64(&m.lastUsage, time.Now().Unix())
	atomic.AddUint64(&m.failedStreamCount, 1)
	// TODO::: Any other job??
}

func (m *Metric) PacketReceived(packetLength uint64) {
	atomic.StoreInt64(&m.lastUsage, time.Now().Unix())
	atomic.AddUint64(&m.packetsReceived, 1)
	atomic.AddUint64(&m.bytesReceived, packetLength)
}

func (m *Metric) PacketSent(packetLength uint64) {
	atomic.StoreInt64(&m.lastUsage, time.Now().Unix())
	atomic.AddUint64(&m.packetsSent, 1)
	atomic.AddUint64(&m.bytesSent, packetLength)
}
