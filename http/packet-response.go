/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"strings"
	"unsafe"
)

// PacketStructureResponse is represent response protocol structure!
// https://tools.ietf.org/html/rfc2616#section-6
type PacketStructureResponse struct {
	Version string // HTTP version
	Status  uint16 // Status code
	Header  Header
	Body    []byte // Packet payload
}

// ParsePacket reads and parses second phase of an incoming SCP packet.
func (p *PacketStructureResponse) ParsePacket(httpPacket []byte) (err error) {
	if len(httpPacket) < PacketLen {
		return ErrHTTPPacketTooShort
	} else if len(httpPacket) > MaxHTTPHeaderSize {
		return ErrHTTPPacketTooLong
	}

	var s = *(*string)(unsafe.Pointer(&httpPacket))

	var index1, index2 int
	index1 = strings.Index(s, " ")
	// First line: HTTP/1.0 200 OK
	p.Version = s[:index1]

	index2 = strings.Index(s[index1+1:], "\n")
	p.Status = uint16(httpPacket[index1+1]) | uint16(httpPacket[index2])<<8

	return nil
}
