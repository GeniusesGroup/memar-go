/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// StreamHandler use to standard stream handler in any layer!
type StreamHandler func(*Server, *Stream)

// Stream : Can use for send or receive data on specific StreamID.
// It can pass to logic layer to give data access to developer!
// Data flow can be up to down (parse raw income data) or down to up (encode app data with respect MTU)
type Stream struct {
	Connection    *Connection
	ReqRes        *Stream    // It point to Request||Response stream related to this stream
	StreamID      uint32     // Odd number for server(who accept connection), Even number for Peer(who start connection).
	LastPacketID  uint32     // Last send or received Packet use to know order of packets!
	ProtocolID    uint16     // Something like TCP||UDP port to indicate data type structure in payload!
	ServiceID     uint32     // Can be ErrorID in response stream!
	Err           error      // Decode||Encode by ErrorID
	Status        chan uint8 // States locate in const of this file.
	Weight        uint8      // 16 queue for priority weight of the streams exist.
	TimeSensitive bool       // If true we must call related service in each received packet. VoIP, IPTV, ...
	Payload       []byte     // Income||Outcome data buffer. Will divide to n packet to respect network MTU!
}

// Stream Status
const (
	StreamStateClosed uint8 = iota
	StreamStateClosing
	StreamStateRateLimited
	StreamStateOpened
	StreamStateOpening
	StreamStateBrokenPacket
	StreamStateReady
)

// MakeResponseStream use to make new Response stream to make UnidirectionalStream to BidirectionalStream!
func (st *Stream) MakeResponseStream() (resStream *Stream) {
	resStream = &Stream{
		Connection: st.Connection,
		ReqRes:     st,
		Status:     make(chan uint8),
		StreamID:   st.StreamID + 1,
	}
	st.Connection.RegisterStream(resStream)
	return
}

// AddNewUIPPacket use to add new UIP packet payload to the stream!
func (st *Stream) AddNewUIPPacket(p []byte, packetID uint32) (err error) {
	// Handle packet received not by order
	if packetID < st.LastPacketID {
		err = ErrPacketArrivedPosterior
	} else if packetID > st.LastPacketID+1 {
		err = ErrPacketArrivedAnterior
	}

	st.LastPacketID = packetID
	copy(st.Payload, p)
	return err
}
