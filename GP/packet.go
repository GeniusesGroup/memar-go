/* For license and copyright information please see LEGAL file in repository */

package gp

const (
	// MinPacketLen is minimum packet length of GP packet
	// 296bit header + 128bit min payload + 256bit signature checksum
	MinPacketLen = 85
)

// CheckPacket will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
func CheckPacket(p []byte) error {
	if len(p) < MinPacketLen {
		return ErrPacketTooShort
	}
	return nil
}

// GetDestinationAddr returns full destination GP address.
func GetDestinationAddr(p []byte) (addr Addr) {
	copy(addr[:], p[:13])
	return
}

// GetSourceAddr returns full source GP address.
func GetSourceAddr(p []byte) (addr Addr) {
	copy(addr[:], p[14:])
	return
}

// GetPayloadLength returns payload length in bytes! Not include headers, padding or checksum!!
func GetPayloadLength(p []byte) uint16 {
	return uint16(p[28]) | uint16(p[29])<<8
}

// GetStreamID returns stream ID.
func GetStreamID(p []byte) uint32 {
	return uint32(p[30]) | uint32(p[31])<<8 | uint32(p[32])<<16 | uint32(p[33])<<24
}

// GetPacketID returns packet ID.
func GetPacketID(p []byte) uint32 {
	return uint32(p[34]) | uint32(p[35])<<8 | uint32(p[36])<<16 | uint32(p[37])<<24
}

// GetPayload returns payload without padding or checksum part.
func GetPayload(p []byte) []byte {
	return p[38 : 38+GetPayloadLength(p)]
}
