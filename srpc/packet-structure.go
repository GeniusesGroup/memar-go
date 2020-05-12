/* For license and copyright information please see LEGAL file in repository */

package srpc

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/sRPC.md
type packetStructure struct {
	ID      uint32 // request>>ServiceID || response>>ErrorID (0 means no error!)
	Payload []byte
}

// CheckPacket will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
// Anyway expectedMinLen can't be under 4-BYTE in any situation!!
func CheckPacket(p []byte, expectedMinLen int) error {
	if len(p) < expectedMinLen {
		return ErrSRPCPacketTooShort
	}
	return nil
}

// GetID decodes service||error ID from the payload buffer.
func GetID(p []byte) uint32 {
	return uint32(p[0]) | uint32(p[1])<<8 | uint32(p[2])<<16 | uint32(p[3])<<24
}

// GetPayload use to get payload of a packet
func GetPayload(p []byte) []byte {
	return p[4:]
}

// SetID encodes service||error ID to the payload buffer.
func SetID(p []byte, id uint32) {
	p[0] = byte(id)
	p[1] = byte(id >> 8)
	p[2] = byte(id >> 16)
	p[3] = byte(id >> 24)
}
