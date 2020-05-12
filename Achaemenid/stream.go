/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// StreamHandler use to standard stream handler in any layer!
type StreamHandler func(*Server, *Stream)

// Stream : Can use for send or receive data on specific StreamID.
// It can pass to logic layer to give data access to developer!
// Data flow can be up to down (parse raw income data) or down to up (encode app data with respect MTU)
type Stream struct {
	Connection      *Connection
	ReqRes          *Stream    // It point to Request||Response stream related to this stream
	StreamID        uint32     // Odd number for server(who accept connection), Even number for Peer(who start connection).
	TotalPacket     uint32     // Expected packets count that must received!
	PacketReceived  uint32     // Count of packets received!
	LastPacketID    uint32     // Last send or received Packet use to know order of packets!
	PacketDropCount uint8      // Count drop packets to prevent some attacks type!
	ProtocolID      uint16     // Something like TCP||UDP port to indicate data type structure in payload!
	ServiceID       uint32     // Can be ErrorID in response stream!
	Err             error      // Decode||Encode by ErrorID
	Status          uint8      // States locate in const of this file.
	StatusChannel   chan uint8 // States locate in const of this file.
	Weight          uint8      // 16 queue for priority weight of the streams exist.
	TimeSensitive   bool       // If true we must call related service in each received packet. VoIP, IPTV, Sensors data, ...
	Payload         []byte     // Income||Outcome data buffer. Will divide to n packet to respect network MTU!
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

// MakeResponseStream use to make new Response stream to convert a UnidirectionalStream to the BidirectionalStream!
func (st *Stream) MakeResponseStream() (resStream *Stream, err error) {
	resStream = &Stream{
		Connection:    st.Connection,
		ReqRes:        st,
		StatusChannel: make(chan uint8),
		StreamID:      st.StreamID + 1,
	}
	st.ReqRes = resStream
	st.Connection.RegisterStream(resStream)
	return
}

// AddNewGPPacket use to add new GP packet payload to the stream!
func (st *Stream) AddNewGPPacket(p []byte, packetID uint32) (err error) {
	// Handle packet received not by order
	if packetID < st.LastPacketID {
		st.Status = StreamStateBrokenPacket
		err = ErrPacketArrivedPosterior
	} else if packetID > st.LastPacketID+1 {
		st.Status = StreamStateBrokenPacket
		err = ErrPacketArrivedAnterior
		// TODO::: send request to sender about not received packets!!
	} else if packetID+1 == st.LastPacketID {
		st.LastPacketID = packetID
	}
	// TODO::: non of above cover for packet 0||1 drop situation!

	// Use PacketID 0||1 for request||response to set stream settings!
	if packetID < 2 {
		st.setStreamSettings(p)
	} else {
		// TODO::: can't easily copy this way!!
		copy(st.Payload, p)
	}

	// Check stream ready situation!
	if st.TotalPacket == st.PacketReceived {
		st.Status = StreamStateReady
	}

	return
}

// Just to show transfer data for completeStreamBySRPC()! We never use this type!
type completeStreamBySRPCReq struct {
	TotalPacket   uint32 // Expected packets count that send over this stream!
	PayloadSize   uint64
	TimeSensitive bool  // If true we must call related service in each received packet. VoIP, IPTV, ...
	Weight        uint8 // 16 queue for priority weight of the streams exist.
}

// setStreamSettings use to set stream settings like time sensitive use in VoIP, IPTV, ...
func (st *Stream) setStreamSettings(p []byte) {
	// TODO::: allow multiple settings set??

	// Dropping packets is preferable to waiting for packets delayed due to retransmissions.
	// Developer can ask to complete data for offline usage after first data usage.
}
