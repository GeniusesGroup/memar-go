/* For license and copyright information please see LEGAL file in repository */

package gp

const (
	// PacketLen is minimum packet length of GP packet
	// 448bit header + 128bit min payload
	PacketLen = 60
)

// CheckPacket will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
func CheckPacket(p []byte) error {
	if len(p) < PacketLen {
		return ErrGPPacketTooShort
	}
	return nil
}

// GetDestinationGP will return DestinationGP in memory safe way!
func GetDestinationGP(p []byte) (GP [16]byte) {
	copy(GP[:], p[0:15])
	return
}

// GetDestinationXP will return DestinationXP in memory safe way!
func GetDestinationXP(p []byte) uint32 {
	return uint32(p[0]) | uint32(p[1])<<8 | uint32(p[2])<<16 | uint32(p[3])<<24
}

// GetDestinationRouter will return DestinationRouter in memory safe way!
func GetDestinationRouter(p []byte) uint32 {
	return uint32(p[4]) | uint32(p[5])<<8 | uint32(p[6])<<16 | uint32(p[7])<<24
}

// GetDestinationUser will return DestinationUser in memory safe way!
func GetDestinationUser(p []byte) uint32 {
	return uint32(p[8]) | uint32(p[9])<<8 | uint32(p[10])<<16 | uint32(p[11])<<24
}

// GetDestinationApp will return DestinationApp in memory safe way!
func GetDestinationApp(p []byte) uint16 {
	return uint16(p[12]) | uint16(p[13])<<8
}

// GetDestinationAppProtocol will return DestinationAppProtocol in memory safe way!
// DestinationAppProtocol use like TCP||UDP port that indicate payload protocol
func GetDestinationAppProtocol(p []byte) uint16 {
	return uint16(p[14]) | uint16(p[15])<<8
}

// GetSourceGP will return SourceGP in memory safe way!
func GetSourceGP(p []byte) (GP [16]byte) {
	copy(GP[:], p[16:31])
	return
}

// GetSourceXP will return SourceXP in memory safe way!
func GetSourceXP(p []byte) uint32 {
	return uint32(p[16]) | uint32(p[17])<<8 | uint32(p[18])<<16 | uint32(p[19])<<24
}

// GetSourceRouter will return SourceRouter in memory safe way!
func GetSourceRouter(p []byte) uint32 {
	return uint32(p[20]) | uint32(p[21])<<8 | uint32(p[22])<<16 | uint32(p[23])<<24
}

// GetSourceUser will return SourceUser in memory safe way!
func GetSourceUser(p []byte) uint32 {
	return uint32(p[24]) | uint32(p[25])<<8 | uint32(p[26])<<16 | uint32(p[27])<<24
}

// GetSourceApp will return SourceApp in memory safe way!
func GetSourceApp(p []byte) uint16 {
	return uint16(p[28]) | uint16(p[29])<<8
}

// GetSourceAppProtocol will return SourceAppProtocol in memory safe way!
// SourceAppProtocol use like TCP||UDP port that indicate payload protocol
func GetSourceAppProtocol(p []byte) uint16 {
	return uint16(p[30]) | uint16(p[31])<<8
}

// GetPayloadLength will return Payload length in bytes! Not include headers, padding or checksum!!
func GetPayloadLength(p []byte) uint16 {
	return uint16(p[32]) | uint16(p[33])<<8
}

// GetStreamID will return StreamID in memory safe way!
func GetStreamID(p []byte) uint32 {
	return uint32(p[34]) | uint32(p[35])<<8 | uint32(p[36])<<16 | uint32(p[37])<<24
}

// GetPacketID will return PacketID in memory safe way!
func GetPacketID(p []byte) uint32 {
	return uint32(p[38]) | uint32(p[39])<<8 | uint32(p[40])<<16 | uint32(p[41])<<24
}

// GetPayload will return Payload without padding or checksum part in memory safe way!
func GetPayload(p []byte) []byte {
	return p[42:GetPayloadLength(p)]
}
