/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"crypto/sha512"

	"../convert"
	er "../error"
)

const (
	// 8-BYTE as Service or Error ID
	MinLength = 8
)

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/sRPC.md
type packetStructure struct {
	ID      uint64 // request>>ServiceID || response>>ErrorID (0 means no error!)
	Payload []byte
}

// CheckPacket will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
// Anyway expectedMinLen can't be under MinLength!
func CheckPacket(p []byte, expectedMinLen int) *er.Error {
	if len(p) < expectedMinLen {
		return ErrPacketTooShort
	}
	return nil
}

// GetID decodes service||error ID from the payload buffer.
func GetID(p []byte) uint64 {
	return uint64(p[0]) | uint64(p[1])<<8 | uint64(p[2])<<16 | uint64(p[3])<<24 | uint64(p[4])<<32 | uint64(p[5])<<40 | uint64(p[6])<<48 | uint64(p[7])<<56
}

// GetPayload use to get payload of a packet
func GetPayload(p []byte) []byte {
	return p[8:]
}

// SetID encodes service||error ID to the payload buffer.
func SetID(p []byte, id uint64) {
	p[0] = byte(id)
	p[1] = byte(id >> 8)
	p[2] = byte(id >> 16)
	p[3] = byte(id >> 24)
	p[4] = byte(id >> 32)
	p[5] = byte(id >> 40)
	p[6] = byte(id >> 48)
	p[7] = byte(id >> 56)
}

// IDCalculator calculate service||error ID by given urn
func IDCalculator(urn string) (id uint64) {
	var hash = sha512.Sum512(convert.UnsafeStringToByteSlice(urn))
	id = GetID(hash[:])
	return
}
