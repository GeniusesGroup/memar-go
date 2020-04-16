/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// Stream : Can use for send or receive data on specific StreamID.
// It can pass to logic layer to give data access to developer!
// Data flow can be up to down (parse raw income data) or down to up (encode app data with respect MTU)
type Stream struct {
	Connection    *Connection
	ReqRes        *Stream // It point to Request||Response stream related to this stream
	StreamID      uint32  // Odd number for server(who accept connection), Even number for Peer(who start connection).
	LastPacketID  uint32  // Last send or received Packet use to know order of packets!
	ProtocolID    uint16  // Something like TCP||UDP port to indicate data type structure in payload!
	ServiceID     uint32
	Status        uint8  // States locate in const of this file.
	Weight        uint8  // 16 queue for priority weight of the streams exist.
	TimeSensitive bool   // If true we must call related service in each received packet. VoIP, IPTV, ...
	Payload       []byte // Income||Outcome data buffer. Will divide to n packet to respect network MTU!
}

// Stream Status
const (
// 0:close 1:open 2:rate-limited 3:closing 4:opening 5:BrokenPacket
)

// NewStream use to make new stream!
// Never make Stream instance by hand, This function can improve by many ways!
func NewStream() *Stream {
	return &Stream{}
}

// AddNewPacket use to add new packet payload to the stream!
func (st *Stream) AddNewPacket(p []byte) {
	// Handle packet received not by order
}

// StreamHandler use to standard stream handler in any layer!
type StreamHandler func(*Server, *Stream)
