/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import "../crypto"

// routerConnection store router to router connection
type routerConnection struct {
	/* Connection data */
	XPID         uint32
	RouterID     uint32 // Part of GP that lease by routers!
	Status       uint8  // States locate in this file.
	Weight       uint8  // 16 queue for priority weight of the connections exist.
	MaxBandwidth uint64 // Peer must respect this, otherwise connection will terminate and User go to black list!

	/* Peer data */
	Path            []byte // Chapar switch spec
	ReversePath     []byte // Chapar switch spec
	AlternativePath [][]byte
	ThingID         [16]byte

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

// routerConnection Status
const (
	// routerConnectionStateClosed indicate connection had been closed
	routerConnectionStateClosed uint8 = iota
	// routerConnectionStateOpen indicate connection is open and ready to use
	routerConnectionStateOpen
	// routerConnectionStateRateLimited indicate connection limited due to higher usage than permitted!
	routerConnectionStateRateLimited
	// routerConnectionStateOpening indicate connection plan to open and not ready to accept stream!
	routerConnectionStateOpening
	// routerConnectionStateClosing indicate connection plan to close and not accept new stream
	routerConnectionStateClosing
	// routerConnectionStateNotResponse indicate peer not response to recently send request!
	routerConnectionStateNotResponse
)
