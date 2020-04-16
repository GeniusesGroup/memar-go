/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"net/url"
	"strings"
	"unsafe"
)

// PacketStructureRequest is represent request protocol structure!
// https://tools.ietf.org/html/rfc2616#section-5
type PacketStructureRequest struct {
	Method  string   // GET, PUT, POST, PATCH, CONNECT, PRI,
	Path    *url.URL //
	Version string   // HTTP version
	Header  Header
	Body    []byte // Packet payload
}

const (
	// PacketLen is minimum Packet length of HTTP Packet.
	PacketLen = 64

	// MaxHTTPHeaderSize is max HTTP header size.
	MaxHTTPHeaderSize = 8192

	// TimeFormat is the time format to use when generating times in HTTP
	// headers. It is like time.RFC1123 but hard-codes GMT as the time
	// zone. The time being formatted must be in UTC for Format to
	// generate the correct format.
	TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
)

// ParsePacket reads and parses second phase of an incoming SCP packet.
func (p *PacketStructureRequest) ParsePacket(httpPacket []byte) (err error) {
	if len(httpPacket) < PacketLen {
		return ErrHTTPPacketTooShort
	} else if len(httpPacket) > MaxHTTPHeaderSize {
		return ErrHTTPPacketTooLong
	}

	var s = *(*string)(unsafe.Pointer(&httpPacket))

	var index1, index2 int
	index1 = strings.Index(s, " ")
	// First line: GET /index.html HTTP/1.0
	p.Method = s[:index1]

	index2 = strings.Index(s[index1+1:], " ")
	p.Path, _ = url.ParseRequestURI(s[index1+1 : index2])

	index1 = strings.Index(s[index2+1:], " ")
	p.Version = s[index2+1 : index1]

	return nil
}
