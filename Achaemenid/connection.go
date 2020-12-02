/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	gp "../GP"
	ip "../IP"
	"../authorization"
	"../crypto"
	etime "../earth-time"
	er "../error"
)

// Connection can use by any type users itself or delegate to other users to act as the owner!
// Each user in each device need unique connection to another party.
type Connection struct {
	/* Connection data */
	Server     *Server
	ID         [32]byte
	State      state
	Weight     Weight
	StreamPool StreamPool

	/* Peer data */
	// Peer Location
	GPAddr  gp.Addr
	IPAddr  ip.Addr // TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
	ThingID [32]byte
	// Peer Identifiers
	UserID           [32]byte // Can't change on StateLoaded!
	UserType         authorization.UserType
	DelegateUserID   [32]byte // Can't change on StateLoaded!
	DelegateUserType authorization.UserType

	/* Security data */
	PeerPublicKey [32]byte
	Cipher        crypto.Cipher // Selected cipher algorithms https://en.wikipedia.org/wiki/Cipher_suite
	AccessControl authorization.AccessControl

	/* Metrics data */
	LastUsage             etime.Time // Last use of this connection
	PacketPayloadSize     uint16     // Always must respect max frame size, so usually packets can't be more than 8192Byte!
	MaxBandwidth          uint64     // Peer must respect this, otherwise connection will terminate and GP go to black list!
	ServiceCallCount      uint64     // Count successful or unsuccessful request.
	BytesSent             uint64     // Counts the bytes of payload data sent.
	PacketsSent           uint64     // Counts packets sent.
	BytesReceived         uint64     // Counts the bytes of payload data Receive.
	PacketsReceived       uint64     // Counts packets Receive.
	FailedPacketsReceived uint64     // Counts failed packets receive for firewalling server from some attack types!
	FailedServiceCall     uint64     // Counts failed service call e.g. data validation failed, ...
}

// MakeIncomeStream make and return the new stream with income ID!
// Never make Stream instance by hand, This function can improve by many ways!
func (conn *Connection) MakeIncomeStream(streamID uint32) (st *Stream, err *er.Error) {
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
func (conn *Connection) MakeOutcomeStream(streamID uint32) (st *Stream, err *er.Error) {
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

// ServiceCallOK tell achaemenid that the service request occur on this connection
func (conn *Connection) ServiceCallOK() (st *Stream) {
	// TODO::: Any job??
	conn.ServiceCallCount++
	return
}

// ServiceCallFail tell achaemenid that bad service request occur on this connection
func (conn *Connection) ServiceCallFail() (st *Stream) {
	// TODO::: Attack?? tel router to block
	conn.FailedServiceCall++
	return
}

// SetThingID set thingID only if it is not set before
func (conn *Connection) SetThingID(thingID [32]byte) {
	if conn.ThingID == [32]byte{} {
		conn.ThingID = thingID
	}
}
