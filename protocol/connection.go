/* For license and copyright information please see LEGAL file in repository */

package protocol

type Connections interface {
	GetConnectionByID(connID [32]byte) (conn NetworkTransportConnection)
	GetConnectionByUserIDThingID(userID, thingID [32]byte) (conn NetworkTransportConnection)
	SaveConnection(conn NetworkTransportConnection)
}

// ConnectionMetrics
type ConnectionMetrics interface {
	ServiceCalled()
	ServiceCallFail()

	LastUsage() Time                     // Last use of the connection
	MaxBandwidth() uint64                // Peer must respect this, otherwise connection will terminate and GP go to black list!
	BytesSent() uint64                   // Counts the bytes of packets sent.
	PacketsSent() uint64                 // Counts sent packets.
	BytesReceived() uint64               // Counts the bytes of packets receive.
	PacketsReceived() uint64             // Counts received packets.
	FailedPacketsReceived() uint64       // Counts failed packets receive for firewalling server from some attack types!
	NotRequestedPacketsReceived() uint64 // Counts not requested packets received for firewalling server from some attack types!
	ServiceCallCount() uint64            // Count successful request.
	FailedServiceCallCount() uint64      // Count failed services call e.g. data validation failed, ...
}

// ConnectionWeight indicate connection and stream state
type ConnectionState uint8

// Standrad Connection States
const (
	ConnectionStateUnset  ConnectionState = iota // State not set yet!
	ConnectionStateNew                           // means connection not saved yet to storage!
	ConnectionStateLoaded                        // means connection load from storage

	ConnectionStateOpening // connection||stream plan to open and not ready to accept stream!
	ConnectionStateOpen    // connection||stream is open and ready to use
	ConnectionStateClosing // connection||stream plan to close and not accept new stream
	ConnectionStateClosed  // connection||stream had been closed

	ConnectionStateNotResponse // peer not response to recently send request!
	ConnectionStateRateLimited // connection||stream limited due to higher usage than permitted!

	ConnectionStateBrokenPacket
	ConnectionStateReceivedCompletely
	ConnectionStateSentCompletely
	ConnectionStateEncrypted
	ConnectionStateDecrypted
	ConnectionStateReady
	ConnectionStateIdle

	ConnectionStateBlocked
	ConnectionStateBlockedByPeer
)

// ConnectionWeight indicate connection and stream weight
type ConnectionWeight uint8

// Standrad Connection Weights
const (
	ConnectionWeightUnset ConnectionWeight = iota

	ConnectionWeightNormal
	ConnectionWeightTimeSensitive // If true must call related service in each received packet. VoIP, IPTV, Sensors data, ...
	// TODO::: 16 queue for priority weight of the connections exist.
)
