/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net"

	"../crypto"
)

// Connection can use by any type users itself or delegate to other users to act as the owner!
// Each user in each device need unique connection to another party.
type Connection struct {
	/* Connection data */
	Server       *Server
	ID           [16]byte
	State        connectionState
	Weight       uint8              // 16 queue for priority weight of the connections exist.
	StreamPool   map[uint32]*Stream // key is Stream.ID
	freeStreamID uint32

	/* Peer data */
	DomainID       [16]byte // Usually use for server to server connections that peer has domainID!
	SocietyID      uint32
	RouterID       uint32
	GPAddr         [16]byte
	IPAddr         *net.IPAddr  // TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
	UDPAddr        *net.UDPAddr // TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
	ThingID        [16]byte
	UserID         [16]byte // Can't change after first set. initial is 0 as Guest!
	UserType       uint8    // 0:Guest, 1:Registered 2:Person, 3:Org, 4:App, ...
	DelegateUserID [16]byte // Can't change after first set

	/* Security data */
	PeerPublicKey [32]byte
	Cipher        crypto.Cipher // Selected cipher algorithms https://en.wikipedia.org/wiki/Cipher_suite
	AccessControl AccessControl

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

type connectionState uint8

// Connection State
const (
	// ConnectionStateClosed indicate connection had been closed
	ConnectionStateClosed connectionState = iota
	// ConnectionStateClosing indicate connection plan to close and not accept new stream
	ConnectionStateClosing
	// ConnectionStateNotResponse indicate peer not response to recently send request!
	ConnectionStateNotResponse
	// ConnectionStateOpen indicate connection is open and ready to use
	ConnectionStateOpen
	// ConnectionStateOpening indicate connection plan to open and not ready to accept stream!
	ConnectionStateOpening
	// ConnectionStateRateLimited indicate connection limited due to higher usage than permitted!
	ConnectionStateRateLimited
)

// MakeUnidirectionalStream use to make a new one way stream!
// Never make Stream instance by hand, This function can improve by many ways!
func (conn *Connection) MakeUnidirectionalStream(streamID uint32) (st *Stream, err error) {
	// TODO::: Check user can open new stream first as stream policy!

	// if given streamID is 0, return new incremental streamID from pool
	if streamID == 0 {
		streamID = conn.freeStreamID
		conn.freeStreamID += 2
	}

	st = &Stream{
		ID:         streamID,
		Connection: conn,
		State:      StreamStateOpened,
	}
	conn.RegisterStream(st)
	return
}

// MakeBidirectionalStream use to make new Request-Response stream!
func (conn *Connection) MakeBidirectionalStream(streamID uint32) (reqStream, resStream *Stream, err error) {
	reqStream, err = conn.MakeUnidirectionalStream(streamID)
	if err != nil {
		return
	}

	resStream, err = reqStream.MakeResponse()
	return
}

// MakeSubscriberStream use to make new Publish–Subscribe stream!
func (conn *Connection) MakeSubscriberStream() (st *Stream) {
	return
}

// GetStreamByID use to get exiting stream in the stream pool of a connection!
func (conn *Connection) GetStreamByID(streamID uint32) *Stream {
	// TODO::: Check stream isn't closed!!
	return conn.StreamPool[streamID]
}

// RegisterStream use to register new stream in the stream pool of a connection!
func (conn *Connection) RegisterStream(st *Stream) {
	conn.StreamPool[st.ID] = st
}

// CloseStream use to close the stream on other side requested or finished!
func (conn *Connection) CloseStream(st *Stream) {
	delete(conn.StreamPool, st.ID)
}

// CloseBidirectionalStream use to close the bidirectional stream!
func (conn *Connection) CloseBidirectionalStream(st *Stream) {
	delete(conn.StreamPool, st.ID)
	delete(conn.StreamPool, st.ReqRes.ID)
}
