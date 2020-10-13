/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net"

	"../authorization"
	"../crypto"
)

// Connection can use by any type users itself or delegate to other users to act as the owner!
// Each user in each device need unique connection to another party.
type Connection struct {
	/* Connection data */
	Server     *Server
	ID         [16]byte
	State      state
	Weight     weight
	StreamPool StreamPool

	/* Peer data */
	DomainID       [16]byte // Usually use for server to server connections that peer has domainID!
	SocietyID      uint32
	RouterID       uint32
	GPAddr         [14]byte   // Without protocol part
	IPAddr         net.IPAddr // TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
	ThingID        [16]byte
	UserID         [16]byte // Can't change after first set. initial is 0 as Guest!
	UserType       uint8    // 0:Guest, 1:Registered 2:Person, 3:Org, 4:App, ...
	DelegateUserID [16]byte // Can't change after first set

	/* Security data */
	PeerPublicKey [32]byte
	Cipher        crypto.Cipher // Selected cipher algorithms https://en.wikipedia.org/wiki/Cipher_suite
	AccessControl authorization.AccessControl

	/* Metrics data */
	PacketPayloadSize     uint16 // Always must respect max frame size, so usually packets can't be more than 8192Byte!
	MaxBandwidth          uint64 // Peer must respect this, otherwise connection will terminate and GP go to black list!
	ServiceCallCount      uint64 // Count successful or unsuccessful request.
	BytesSent             uint64 // Counts the bytes of payload data sent.
	PacketsSent           uint64 // Counts packets sent.
	BytesReceived         uint64 // Counts the bytes of payload data Receive.
	PacketsReceived       uint64 // Counts packets Receive.
	FailedPacketsReceived uint64 // Counts failed packets receive for firewalling server from some attack types!
	FailedServiceCall     uint64 // Counts failed service call e.g. data validation failed, ...
}

// MakeIncomeStream make and return the new stream with income ID!
// Never make Stream instance by hand, This function can improve by many ways!
func (conn *Connection) MakeIncomeStream(streamID uint32) (st *Stream, err error) {
	// TODO::: Check user can open new stream first as stream policy!

	// if given streamID is 0, return new incremental streamID from pool
	if streamID == 0 {
		streamID = conn.StreamPool.freeIncomeStreamID
		conn.StreamPool.freeIncomeStreamID += 2
	}

	st = &Stream{
		ID:           streamID,
		Connection:   conn,
		State:        StateOpen,
		StateChannel: make(chan state),
	}
	conn.StreamPool.RegisterStream(st)
	return
}

// MakeOutcomeStream make and return the new stream with outcome ID!
// Never make Stream instance by hand, This function can improve by many ways!
func (conn *Connection) MakeOutcomeStream(streamID uint32) (st *Stream, err error) {
	// TODO::: Check user can open new stream first as stream policy!

	// if given streamID is 0, return new incremental streamID from pool
	if streamID == 0 {
		streamID = conn.StreamPool.freeOutcomeStreamID
		conn.StreamPool.freeOutcomeStreamID += 2
	}

	st = &Stream{
		ID:           streamID,
		Connection:   conn,
		State:        StateOpen,
		StateChannel: make(chan state),
	}
	conn.StreamPool.RegisterStream(st)
	return
}

// MakeSubscriberStream make new Publish–Subscribe stream!
func (conn *Connection) MakeSubscriberStream() (st *Stream) {
	// TODO:::
	return
}
