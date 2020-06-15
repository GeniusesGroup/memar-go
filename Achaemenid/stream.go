/* For license and copyright information please see LEGAL file in repository */

package achaemenid

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

// MakeResponseStream use to make new Response stream to convert a UnidirectionalStream to the BidirectionalStream!
func (st *Stream) MakeResponseStream() (resStream *Stream, err error) {
	resStream = &Stream{
		ID:           st.ID + 1,
		Connection:   st.Connection,
		ReqRes:       st,
		StateChannel: make(chan streamState),
	}
	st.ReqRes = resStream
	st.Connection.RegisterStream(resStream)
	return
}
