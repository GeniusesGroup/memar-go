/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import "../crypto"

// Connection can use by any type users itself or delegate to other users to act as the owner!
// Each user in each device need unique connection to another party.
type Connection struct {
	/* Connection data */
	ConnectionID [16]byte
	Status       uint8              // States locate in const of this file.
	Weight       uint8              // 16 queue for priority weight of the connections exist.
	StreamPool   map[uint32]*Stream // StreamID

	/* Peer data */
	DomainID       [16]byte // Usually use for server to server connections that peer has domainID!
	UIPAddress     [16]byte
	ThingID        [16]byte
	UserID         [16]byte // Can't change after first set. initial is 0 as Guest!
	UserType       uint8    // 0:Guest, 1:Registered 2:Person, 3:Org, 4:App, ...
	DelegateUserID [16]byte // Can't change after first set. Guest={1}

	/* Security data */
	PeerPublicKey [32]byte
	Cipher        crypto.Cipher // Selected cipher algorithms https://en.wikipedia.org/wiki/Cipher_suite
	AccessControl AccessControl

	/* Metrics data */
	PacketPayloadSize     uint16 // Always must respect max frame size, so usually packets can't be more than 8192Byte!
	MaxBandwidth          uint64 // Peer must respect this, otherwise connection will terminate and UIP go to black list!
	ServiceCallCount      uint64 // Count successful or unsuccessful request.
	BytesSent             uint64 // Counts the bytes of payload data sent.
	PacketsSent           uint64 // Counts packets sent.
	BytesReceived         uint64 // Counts the bytes of payload data Receive.
	PacketsReceived       uint64 // Counts packets Receive.
	FailedPacketsReceived uint64 // Counts failed packets receive for firewalling server from some attack types!
	FailedServiceCall     uint64 // Counts failed service call e.g. data validation failed, ...
}

// Connection Status
const (
	// ConnectionStateClosed indicate connection had been closed
	ConnectionStateClosed uint8 = iota
	// ConnectionStateOpen indicate connection is open and ready to use
	ConnectionStateOpen
	// ConnectionStateRateLimited indicate connection limited due to higher usage than permitted!
	ConnectionStateRateLimited
	// ConnectionStateOpening indicate connection plan to open and not ready to accept stream!
	ConnectionStateOpening
	// ConnectionStateClosing indicate connection plan to close and not accept new stream
	ConnectionStateClosing
	// ConnectionStateNotResponse indicate peer not response to recently send request!
	ConnectionStateNotResponse
)

// MakeUnidirectionalStream use to make a new one way stream!
// Never make Stream instance by hand, This function can improve by many ways!
func (conn *Connection) MakeUnidirectionalStream(streamID uint32) (st *Stream, err error) {
	// TODO::: Check user can open new stream first as stream policy!

	if streamID == 0 {
		// TODO::: Get new incremental streamID from pool
	}

	st = &Stream{
		StreamID:      streamID,
		StatusChannel: make(chan uint8),
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

	resStream, err = reqStream.MakeResponseStream()
	return
}

// MakeSubscriberStream use to make new Publish–Subscribe stream!
func (conn *Connection) MakeSubscriberStream() (st *Stream) {
	return
}

// GetStreamByID use to get exiting stream in the stream pool of a connection!
func (conn *Connection) GetStreamByID(streamID uint32) (st *Stream, ok bool) {
	st, ok = conn.StreamPool[streamID]
	// TODO::: Check stream isn't closed!!
	return
}

// RegisterStream use to register new stream in the stream pool of a connection!
func (conn *Connection) RegisterStream(st *Stream) {
	conn.StreamPool[st.StreamID] = st
}

// CloseStream use to close the stream on other side requested or finished!
func (conn *Connection) CloseStream(st *Stream) {
	delete(conn.StreamPool, st.StreamID)
}
