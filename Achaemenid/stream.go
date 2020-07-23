/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "../errors"

// Stream use to send or receive data on specific connection.
// It can pass to logic layer to give data access to developer!
// Data flow can be up to down (parse raw income data) or down to up (encode app data with respect MTU)
type Stream struct {
	ID         uint32 // Odd number for server(who accept connection), Even number for Peer(who start connection).
	Connection *Connection
	ReqRes     *Stream // It point to Request||Response stream related to this stream
	ProtocolID uint16  // Something like TCP||UDP port to indicate data type structure in payload!
	ServiceID  uint32  // Can be ErrorID in response stream!
	Payload    []byte  // Income||Outcome data buffer. Will divide to n packet to respect network MTU!

	/* State */
	Err           error            // Decode||Encode by ErrorID
	State         streamState      // States locate in const of this file.
	StateChannel  chan streamState // States locate in const of this file.
	Weight        uint8            // 16 queue for priority weight of the streams exist.
	TimeSensitive bool             // If true we must call related service in each received packet. VoIP, IPTV, Sensors data, ...

	/* Metrics */
	TotalPacket     uint32 // Expected packets count that must received!
	PacketReceived  uint32 // Count of packets received!
	LastPacketID    uint32 // Last send or received Packet use to know order of packets!
	PacketDropCount uint8  // Count drop packets to prevent some attacks type!
}

type streamState uint8

// Stream Status
const (
	StreamStateClosed streamState = iota
	StreamStateClosing
	StreamStateRateLimited
	StreamStateOpened
	StreamStateOpening
	StreamStateBrokenPacket
	StreamStateReady
)

// Errors
var (
	ErrNoConnectionToNode = errors.New("NoConnectionToNode", "There is no connection to desire node to complete request")
)

// SetState change state of stream and send notification on stream StateChannel.
func (st *Stream) SetState(state streamState) {
	st.State = state
	// notify stream listener that stream ready to use!
	st.StateChannel <- StreamStateReady
}

// MakeResponse make new Response stream to convert a UnidirectionalStream to the BidirectionalStream!
func (st *Stream) MakeResponse() (resStream *Stream, err error) {
	resStream = &Stream{
		ID:           st.ID + 1,
		Connection:   st.Connection,
		ReqRes:       st,
		State:        StreamStateOpened,
		StateChannel: make(chan streamState),
	}
	st.ReqRes = resStream
	st.State = StreamStateOpened
	st.Connection.RegisterStream(resStream)
	return
}

// Send register stream in send pool. Usually use to send response stream.
func (st *Stream) Send() (err error) {
	// First Check st.Connection.Status to ability send stream over it

	// TODO::: remove this check when remove IP support
	if st.Connection.IPAddr != nil {
		// Send by IP
	} else if st.Connection.UDPAddr != nil {
		// Send by UDP
	} else {
		// Send stream by GP
	}

	// send stream by weight

	st.Connection.BytesSent += uint64(len(st.Payload))

	// Last Close stream!
	st.Connection.CloseStream(st)

	return
}

// SendReq register stream in send pool and block caller until response ready to read.
func (st *Stream) SendReq() (err error) {
	st.Send()

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus streamState = <-st.ReqRes.StateChannel
	if responseStatus == StreamStateReady {

	} else {

	}

	// Last Close response stream!
	st.Connection.CloseStream(st.ReqRes)

	return
}

// Authorize authorize request by data in related stream and connection.
func (st *Stream) Authorize() (err error) {
	err = st.Connection.AccessControl.authorizeWhen()
	if err != nil {
		return
	}
	err = st.Connection.AccessControl.authorizeWhich(st.ServiceID)
	if err != nil {
		return
	}
	err = st.Connection.AccessControl.authorizeWhere(st.Connection.SocietyID, st.Connection.RouterID)
	if err != nil {
		return
	}
	return
}

// MakeBidirectionalStream use to make new Request-Response stream without any connection exist yet!
// TODO::: due to IPv4&&IPv6 support we need this func! Remove it when remove those support!
func MakeBidirectionalStream() (reqStream, resStream *Stream, err error) {
	// TODO::: check server allow make new connection and has enough resource

	reqStream = &Stream{
		ID:           0,
		State:        StreamStateOpening,
		StateChannel: make(chan streamState),
	}
	resStream = &Stream{
		ID:           1,
		ReqRes:       reqStream,
		State:        StreamStateOpening,
		StateChannel: make(chan streamState),
	}
	reqStream.ReqRes = resStream
	return
}
