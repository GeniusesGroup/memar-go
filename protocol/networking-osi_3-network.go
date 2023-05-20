/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Network (OSI Layer 2: Network)
**********************************************************************************
*/

// NetworkNetwork_Multiplexer indicate a transport packet multiplexer.
// OS and APP part must implement in dedicate structures.
type NetworkNetwork_Multiplexer interface {
	HeaderID() (headerID NetworkLink_NextHeaderID)

	// Receiver must release packet slice and don't use it after return.
	// Almost in most cases dev must copy packet payload to a stream.
	Receive(conn NetworkLink_Connection, packet []byte)

	ObjectLifeCycle
}

// Connection or App2AppConnection indicate how connection create and save in time series data.
// Each user pair with delegate-user has a chain that primary key is UserID+DelegateUserID
type Connection interface {
	/* Peer data */
	DomainName() string // if exist
	UserID() UserID
	DelegateUserID() UserID // Persons can delegate to things(as a user type)

	Close() (err Error)  // Just once
	Revoke() (err Error) // Just once

	NetworkAddress // string form of address (for example, "ipv4://192.0.2.1", "ipv6://[2001:db8::1]")
	NetworkMTU
	Network_Status
	OperationImportance
	ConnectionLowLevelAPIs
	Streams
	ConnectionMetrics

	/* Security data */
	// Cipher() Cipher
}

// ConnectionLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
// It will use in chunks managing packages e.g. sRPC, QUIC, TCP, UDP, ... or Application layer protocols e.g. HTTP, ...
type ConnectionLowLevelAPIs interface {
	NewPacket(payloadLen int) (packet []byte, payload []byte, err Error)

	// Send transmitting in non blocking mode and just queue frames for congestion situations.
	// Return error not related to packet situation, just about any hardware error.
	// A situation might be occur that the port available when a packet queued,
	// but when the time to send is come, the port broken and sender don't know about this.
	// Due to speed matters in link layer, and it is very rare situation, it is better to ignore suddenly port unavailability.
	// After return, caller can't reuse payload array anymore.
	Send(packet []byte) (err Error) // to send data use Send() that exist in stream of each connection
}
