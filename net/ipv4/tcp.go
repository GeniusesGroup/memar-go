/* For license and copyright information please see the LEGAL file in the code repository */

package ipv4

import (
	"libgo/protocol"
	"libgo/tcp"
)

var tcpSockets = make(map[ipv4SocketKey]*tcp.Stream, 1024)

type ipv4SocketKey struct {
	SourceIP        [4]byte
	SourcePort      uint16
	DestinationIP   [4]byte
	DestinationPort uint16
}

// ReceiveOverIPv4 hold packet for some times, So sender must know to not reuse packet memory location.
// ReceiveOverIPv4 Don't hold packet, So sender can reuse packet slice for any purpose.
// It must be non blocking and just route packet not to wait for anything else.
func ReceiveOverIPv4(tcpRawSegment []byte, srcIPAddr, desIPAddr [4]byte) (err protocol.Error) {
	var tcpSegment = tcp.Segment(tcpRawSegment)
	// Find proper stream or make new one if allow by some rules
	var sKey = ipv4SocketKey{
		SourceIP:        srcIPAddr,
		SourcePort:      tcpSegment.SourcePort(),
		DestinationIP:   desIPAddr,
		DestinationPort: tcpSegment.DestinationPort(),
	}
	var stream = tcpSockets[sKey]
	if stream == nil {
		stream, err = newSocketOverIPv4(tcpSegment, sKey)
		if err != nil {
			return
		}
	}

	err = stream.Receive(tcpSegment)
	return
}

func newSocketOverIPv4(tcpFrame tcp.Segment, sKey ipv4SocketKey) (stream *tcp.Stream, err protocol.Error) {
	// TODO:::
	tcpSockets[sKey] = stream
	return
}
