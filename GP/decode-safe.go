/* For license and copyright information please see LEGAL file in repository */

package gp

import "errors"

const (
	// PacketLen is minimum packet length of GP packet
	// 448bit header + 128bit min payload
	PacketLen = 60
)

// Declare Errors Details
var (
	ErrGPPacketTooShort = errors.New("GP packet is empty or too short than standard header. It must include 44Byte header plus 16Byte min Payload")
)

// CheckPacket  check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic occur!
func CheckPacket(p []byte) error {
	if len(p) < PacketLen {
		return ErrGPPacketTooShort
	}
	return nil
}

// GetDestinationGP returns full destination GP address.
func GetDestinationGP(p []byte) (GP [16]byte) {
	copy(GP[:], p[0:15])
	return
}

// GetDestinationSociety returns destination society ID.
func GetDestinationSociety(p []byte) uint32 {
	return uint32(p[0]) | uint32(p[1])<<8 | uint32(p[2])<<16 | uint32(p[3])<<24
}

// GetDestinationRouter returns destination router ID.
func GetDestinationRouter(p []byte) uint32 {
	return uint32(p[4]) | uint32(p[5])<<8 | uint32(p[6])<<16 | uint32(p[7])<<24
}

// GetDestinationUser returns destination user ID!
func GetDestinationUser(p []byte) uint32 {
	return uint32(p[8]) | uint32(p[9])<<8 | uint32(p[10])<<16 | uint32(p[11])<<24
}

// GetDestinationApp returns destination app ID!
func GetDestinationApp(p []byte) uint16 {
	return uint16(p[12]) | uint16(p[13])<<8
}

// GetDestinationAppProtocol returns destination app protocol ID.
// app protocol ID usage is like TCP||UDP port that indicate payload protocol
func GetDestinationAppProtocol(p []byte) uint16 {
	return uint16(p[14]) | uint16(p[15])<<8
}

// GetSourceGP returns source GP address.
func GetSourceGP(p []byte) (GP [16]byte) {
	copy(GP[:], p[16:31])
	return
}

// GetSourceSociety returns source society ID.
func GetSourceSociety(p []byte) uint32 {
	return uint32(p[16]) | uint32(p[17])<<8 | uint32(p[18])<<16 | uint32(p[19])<<24
}

// GetSourceRouter returns source router ID.
func GetSourceRouter(p []byte) uint32 {
	return uint32(p[20]) | uint32(p[21])<<8 | uint32(p[22])<<16 | uint32(p[23])<<24
}

// GetSourceUser returns source user ID.
func GetSourceUser(p []byte) uint32 {
	return uint32(p[24]) | uint32(p[25])<<8 | uint32(p[26])<<16 | uint32(p[27])<<24
}

// GetSourceApp returns source app ID.
func GetSourceApp(p []byte) uint16 {
	return uint16(p[28]) | uint16(p[29])<<8
}

// GetSourceAppProtocol returns source app protocol ID.
// app protocol ID usage is like TCP||UDP port that indicate payload protocol
func GetSourceAppProtocol(p []byte) uint16 {
	return uint16(p[30]) | uint16(p[31])<<8
}

// GetPayloadLength returns payload length in bytes! Not include headers, padding or checksum!!
func GetPayloadLength(p []byte) uint16 {
	return uint16(p[32]) | uint16(p[33])<<8
}

// GetStreamID returns stream ID.
func GetStreamID(p []byte) uint32 {
	return uint32(p[34]) | uint32(p[35])<<8 | uint32(p[36])<<16 | uint32(p[37])<<24
}

// GetPacketID returns packet ID.
func GetPacketID(p []byte) uint32 {
	return uint32(p[38]) | uint32(p[39])<<8 | uint32(p[40])<<16 | uint32(p[41])<<24
}

// GetPayload returns payload without padding or checksum part.
func GetPayload(p []byte) []byte {
	return p[42:GetPayloadLength(p)]
}
