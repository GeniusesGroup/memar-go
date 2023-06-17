/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// ConnectionMetrics
type ConnectionMetrics interface {
	LastUsage() Time                     // Last use of the connection
	MaxBandwidth() uint64                // Byte/Second and Connection can limit to a fixed number
	BytesSent() uint64                   // Counts the bytes of packets sent
	PacketsSent() uint64                 // Counts sent packets
	BytesReceived() uint64               // Counts the bytes of packets receive
	PacketsReceived() uint64             // Counts received packets
	LostPackets() uint64                 // Counts any lost packet that peer request to resend it
	LostBytes() uint64                   //
	ResendPackets() uint64               // Counts any duplicate packet that not request to resend it, use to prevent attacks
	ResendBytes() uint64                 //
	FailedPacketsSent() uint64           //
	FailedPacketsReceived() uint64       // Counts failed packets receive for firewalling server from some attack types
	NotRequestedPacketsReceived() uint64 // Counts not requested packets received for firewalling server from some attack types
	SucceedStreamCount() uint64          // Count successful request
	FailedStreamCount() uint64           // Count failed services call e.g. data validation failed, ...

	StreamSucceed()
	StreamFailed()
	PacketReceived(packetLength uint64)
	DuplicatePacketReceived(packetLength uint64)
	PacketSent(packetLength uint64)
	PacketResend(packetLength uint64)

	// Rate() uint64 // Byte/Second
}

// ConnectionMetrics is an atomic counters.
type ConnectionsMetrics interface {
	LastUsage() Time            // Last use of the connection
	OpenCount() int64           // number of opened and pending open connections
	GuestCount() int64          // number of opened connection for guest users
	InUseCount() int64          // The number of connections currently in use.
	IdleCount() int64           // The number of idle connections.
	WaitCount() int64           // Total number of connections waited for.
	ClosedCount() int64         // Total number of closed connections.
	IdleClosedCount() int64     // Total number of idle state connections closed due to timeout.
	WaitClosedCount() int64     // Total number of wait state connections closed due to timeout.
	LifetimeClosedCount() int64 // Total number of connections closed due to max connection lifetime limit.

	// Rate() uint64 // Byte/Second
}
