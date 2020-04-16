/* For license and copyright information please see LEGAL file in repository */

package uip

const (
	// PacketLen is minimum packet length of UIP packet
	// 448bit header + 128bit min payload
	PacketLen = 60
)

// CheckPacket will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
func CheckPacket(p []byte) error {
	if len(p) < PacketLen {
		return ErrUIPPacketTooShort
	}
	return nil
}

// GetDestinationUIP will return DestinationUIP in memory safe way!
func GetDestinationUIP(p []byte) (uip [16]byte) {
	copy(uip[:], p[0:15])
	return
}

// GetDestinationAppProtocol will return DestinationAppProtocol in memory safe way!
// DestinationAppProtocol use like TCP||UDP port that indicate payload protocol
func GetDestinationAppProtocol(p []byte) (destinationAppProtocol uint16) {
	return uint16(p[14]) | uint16(p[15])<<8
}

// GetSourceUIP will return SourceUIP in memory safe way!
func GetSourceUIP(p []byte) (uip [16]byte) {
	copy(uip[:], p[16:31])
	return
}

// GetSourceAppProtocol will return SourceAppProtocol in memory safe way!
// SourceAppProtocol use like TCP||UDP port that indicate payload protocol
func GetSourceAppProtocol(p []byte) (sourceSessionID uint16) {
	return uint16(p[30]) | uint16(p[31])<<8
}

// GetStreamID will return StreamID in memory safe way!
func GetStreamID(p []byte) uint32 {
	return uint32(p[32]) | uint32(p[33])<<8 | uint32(p[34])<<16 | uint32(p[35])<<24
}

// GetPacketID will return PacketID in memory safe way!
func GetPacketID(p []byte) uint32 {
	return uint32(p[36]) | uint32(p[37])<<8 | uint32(p[38])<<16 | uint32(p[39])<<24
}

// GetPayload will return Payload in memory safe way!
func GetPayload(p []byte) []byte {
	return p[40 : len(p)-4]
}

// GetPayloadLength will return Payload length!
func GetPayloadLength(p []byte) int {
	return len(p) - 44
}
