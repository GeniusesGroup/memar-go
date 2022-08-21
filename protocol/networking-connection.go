/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Connection or App2AppConnection indicate how connection create and save in time series data.
// Each user pair with delegate-user has a chain that primary key is UserID+DelegateUserID
type Connection interface {
	/* Connection data */
	HeaderID() NetworkLink_NextHeaderID
	MTU() int

	/* Peer data */
	LocalAddr() []byte
	RemoteAddr() []byte
	DomainName() string // if exist
	UserID() UserID
	DelegateUserID() UserID // Persons can delegate to things(as a user type)

	Close() (err Error)  // Just once
	Revoke() (err Error) // Just once
	
	Network_Status
	OperationImportance
	ConnectionLowLevelAPIs
	Streams
	ConnectionMetrics
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
