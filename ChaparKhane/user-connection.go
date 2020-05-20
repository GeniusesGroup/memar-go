/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import "../crypto"

// UserConnection can use by any type users
type UserConnection struct {
	/* Connection data */
	UserConnectionID [16]byte
	UserGPID         uint32 // Part of GP that lease by OS delegate of a user!
	Status           uint8  // States locate in this file.
	Weight           uint8  // 16 queue for priority weight of the connections exist.
	MaxBandwidth     uint64 // Peer must respect this, otherwise connection will terminate and User go to black list!

	/* Peer data */
	Path            []byte // Chapar switch spec
	ReversePath     []byte // Chapar switch spec
	AlternativePath [][]byte
	ThingID         [16]byte
	UserID          [16]byte // Can't change after first set. 0 for Guest!
	UserType        uint8    // 0:Guest, 1:Registered 2:Person, 3:Org, 4:App, ...

	/* Security data */
	PeerPublicKey [32]byte
	Cipher        crypto.Cipher // Selected cipher algorithms https://en.wikipedia.org/wiki/Cipher_suite

	/* Metrics data */
	BytesSent             uint64 // Counts the bytes of payload data sent.
	PacketsSent           uint64 // Counts packets sent.
	BytesReceived         uint64 // Counts the bytes of payload data Receive.
	PacketsReceived       uint64 // Counts packets Receive.
	FailedPacketsReceived uint64 // Counts failed packets receive for firewalling server from some attack types!
}

// UserConnection Status
const (
	// UserConnectionStateClosed indicate connection had been closed
	UserConnectionStateClosed uint8 = iota
	// UserConnectionStateOpen indicate connection is open and ready to use
	UserConnectionStateOpen
	// UserConnectionStateRateLimited indicate connection limited due to higher usage than permitted!
	UserConnectionStateRateLimited
	// UserConnectionStateOpening indicate connection plan to open and not ready to accept stream!
	UserConnectionStateOpening
	// UserConnectionStateClosing indicate connection plan to close and not accept new stream
	UserConnectionStateClosing
	// UserConnectionStateNotResponse indicate peer not response to recently send request!
	UserConnectionStateNotResponse
)
