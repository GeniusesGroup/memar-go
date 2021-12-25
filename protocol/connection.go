/* For license and copyright information please see LEGAL file in repository */

package protocol

type Connections interface {
	GetConnectionByPeerAddr(addr [16]byte) (conn Connection, err Error)
	// A connection can use just by single app node, so user can't use same connection to connect other node before close connection on usage node.
	GetConnectionByUserIDDelegateUserID(userID, delegateUserID [16]byte) (conn Connection, err Error)
	GetConnectionsByUserID(userID [16]byte) (conns []Connection, err Error)
	GetConnectionByDomain(domain string) (conn Connection, err Error)

	RegisterConnection(conn Connection, err Error)
	CloseConnection(conn Connection, err Error)
	RevokeConnection(conn Connection, err Error)
}

// Connection or App2AppConnection indicate how connection create and save in time series data.
// Each user pair with delegate-user has a chain that primary key is UserID+DelegateUserID
type Connection interface {
	/* Connection data */
	MTU() int
	Status() ConnectionState     // return last connection state
	State() chan ConnectionState // return state channel to listen to new connection state. for more than one listener use channel hub(repeater)
	Weight() ConnectionWeight

	/* Peer data */
	Addr() [16]byte
	AddrType() NetworkLinkNextHeaderID
	DomainName() string // if exist
	UserID() UserID
	DelegateUserID() UserID // Persons can delegate to things(as a user type) in old devices that need to use protocols like HTTP, ...

	/* Security data */
	Cipher() Cipher

	Streams
	ConnectionMetrics
}

// ConnectionMetrics
type ConnectionMetrics interface {
	LastUsage() Time                     // Last use of the connection
	MaxBandwidth() uint64                // Byte/Second and Connection can limit to a fixed number
	BytesSent() uint64                   // Counts the bytes of packets sent.
	PacketsSent() uint64                 // Counts sent packets.
	BytesReceived() uint64               // Counts the bytes of packets receive.
	PacketsReceived() uint64             // Counts received packets.
	FailedPacketsReceived() uint64       // Counts failed packets receive for firewalling server from some attack types!
	NotRequestedPacketsReceived() uint64 // Counts not requested packets received for firewalling server from some attack types!
	SucceedStreamCount() uint64          // Count successful request.
	FailedStreamCount() uint64           // Count failed services call e.g. data validation failed, ...

	StreamSucceed()
	StreamFailed()
	PacketReceived(packetLength uint64)
	PacketSent(packetLength uint64)
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
	ConnectionStateSending
	ConnectionStateReceiving
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
