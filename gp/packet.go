/* For license and copyright information please see LEGAL file in repository */

package gp

import "../protocol"

const (
	// MinPacketLen is minimum packet length of GP packet
	// 320bit header + 128bit min payload
	MinPacketLen = 56
)

type packet []byte

// CheckPacket will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
func CheckPacket(p []byte) protocol.Error {
	if len(p) < MinPacketLen {
		return ErrPacketTooShort
	}
	return nil
}

// GetDestinationAddr returns full destination GP address.
func GetDestinationAddr(p []byte) (addr Addr) {
	copy(addr[:], p[:])
	return
}

// GetSourceAddr returns full source GP address.
func GetSourceAddr(p []byte) (addr Addr) {
	copy(addr[:], p[16:])
	return
}

// GetPacketNumber returns packet number
func GetPacketNumber(p []byte) uint64 {
	return uint64(p[32]) | uint64(p[33])<<8 | uint64(p[34]<<16) | uint64(p[35])<<24 | uint64(p[36]<<32) | uint64(p[37])<<40 | uint64(p[38]<<48) | uint64(p[39])<<56
}

// GetPayload returns payload that means all data after packetNumber
func GetPayload(p []byte) []byte {
	return p[40:]
}
