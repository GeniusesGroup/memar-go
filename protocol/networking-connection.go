/* For license and copyright information please see LEGAL file in repository */

package protocol

type Connections interface {
	GuestConnectionCount() uint64

	GetConnectionByPeerAddr(addr [16]byte) (conn Connection, err Error)
	// A connection can use just by single app node, so user can't use same connection to connect other node before close connection on usage node.
	GetConnectionByUserIDDelegateUserID(userID, delegateUserID [16]byte) (conn Connection, err Error)
	GetConnectionsByUserID(userID [16]byte) (conns []Connection, err Error)
	GetConnectionByDomain(domain string) (conn Connection, err Error)

	// state=unregistered -> 'register' -> state=registered -> 'deregister' -> state=unregistered.
	RegisterConnection(conn Connection) (err Error)
	DeregisterConnection(conn Connection) (err Error)
}

// Connection or App2AppConnection indicate how connection create and save in time series data.
// Each user pair with delegate-user has a chain that primary key is UserID+DelegateUserID
type Connection interface {
	/* Connection data */
	MTU() int
	Status() ConnectionState     // return last connection state
	State() chan ConnectionState // return state channel to listen to new connection state. for more than one listener use channel hub(repeater)
	Weight() Weight

	/* Peer data */
	Addr() [16]byte
	AddrType() NetworkLinkNextHeaderID
	DomainName() string // if exist
	UserID() UserID
	DelegateUserID() UserID // Persons can delegate to things(as a user type)

	Close() (err Error)  // Just once
	Revoke() (err Error) // Just once

	Streams
	ConnectionMetrics
}

// ConnectionMetrics
type ConnectionMetrics interface {
	LastUsage() TimeUnixMilli            // Last use of the connection
	MaxBandwidth() uint64                // Byte/Second and Connection can limit to a fixed number
	BytesSent() uint64                   // Counts the bytes of packets sent
	PacketsSent() uint64                 // Counts sent packets
	BytesReceived() uint64               // Counts the bytes of packets receive
	PacketsReceived() uint64             // Counts received packets
	LostPackets() uint64                 //
	LostBytes() uint64                   //
	FailedPacketsReceived() uint64       // Counts failed packets receive for firewalling server from some attack types
	NotRequestedPacketsReceived() uint64 // Counts not requested packets received for firewalling server from some attack types
	SucceedStreamCount() uint64          // Count successful request
	FailedStreamCount() uint64           // Count failed services call e.g. data validation failed, ...

	StreamSucceed()
	StreamFailed()
	PacketReceived(packetLength uint64)
	PacketSent(packetLength uint64)
	PacketResend(packetLength uint64)
}

// ConnectionState indicate connection and stream state
type ConnectionState uint8

// Standrad Connection States
const (
	ConnectionState_Unset  ConnectionState = iota // State not set yet
	ConnectionState_New                           // means connection not saved yet to storage
	ConnectionState_Loaded                        // means connection load from storage
	ConnectionState_Unregistered

	ConnectionState_Opening // connection||stream plan to open and not ready to accept stream
	ConnectionState_Open    // connection||stream is open and ready to use
	ConnectionState_Closing // connection||stream plan to close and not accept new stream
	ConnectionState_Closed  // connection||stream had been closed

	ConnectionState_NotResponse // peer not response to recently send request
	ConnectionState_RateLimited // connection||stream limited due to higher usage than permitted
	ConnectionState_Timeout     // connection||stream timeout(DeadlineExceeded) and must close

	ConnectionState_BrokenPacket
	ConnectionState_NeedMoreData
	ConnectionState_Sending
	ConnectionState_Receiving
	ConnectionState_ReceivedCompletely
	ConnectionState_SentCompletely

	ConnectionState_Encrypted
	ConnectionState_Decrypted
	ConnectionState_Ready
	ConnectionState_Idle

	ConnectionState_Blocked
	ConnectionState_BlockedByPeer

	// Each package can declare it's own state with ConnectionState > 127 e.g. tcp: TCP_SYN_SENT, TCP_SYN_RECV, ...
)
