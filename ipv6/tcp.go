/* For license and copyright information please see the LEGAL file in the code repository */

package ipv6

import (
	"libgo/protocol"
	"libgo/tcp"
)

// TODO::: due to below map use by many core, in insert and expand time it must lock globally,
// we must implement internal hash table to improve performance by lock in bucket and grow table faster
var tcpSockets = make(map[ipv6SocketKey]*tcp.Stream, 1024)

type ipv6SocketKey struct {
	SourceIP        Addr
	SourcePort      uint16
	DestinationIP   Addr
	DestinationPort uint16
}

// ReceiveTCPOverIPv6 hold packet for some times, So sender must know to not reuse packet memory location.
// ReceiveTCPOverIPv6 Don't hold packet, So sender can reuse packet slice for any purpose.
// It must be non blocking and just route packet not to wait for anything else.
func ReceiveTCPOverIPv6(srcIPAddr, desIPAddr Addr, tcpRawSegment []byte) (err protocol.Error) {
	var tcpSegment = tcp.Segment(tcpRawSegment)
	var srcPort = tcpSegment.SourcePort()
	var desPort = tcpSegment.DestinationPort()
	// Find proper socket or make new one if allow by some rules
	var sKey = ipv6SocketKey{
		SourceIP:        srcIPAddr,
		SourcePort:      srcPort,
		DestinationIP:   desIPAddr,
		DestinationPort: desPort,
	}

	// Check application support requested protocol
	var ProtocolID = protocol.NetworkApplication_ProtocolID(desPort)
	var protocolHandler protocol.NetworkApplication_Handler = protocol.App.GetNetworkApplicationHandler(ProtocolID)
	if protocolHandler == nil {
		// Send response or just ignore packet
		// TODO::: DDOS!!??
		return
	}

	if handleMode_Listener {
		// TODO:::
	}

	var st = tcpSockets[sKey]
	if st == nil {
		st, err = newSocketOverIPv6(tcpSegment, sKey)
		if err != nil {
			return
		}
	}

	err = st.Receive(tcpSegment)
	return
}

func newSocketOverIPv6(tcpSegment tcp.Segment, sKey ipv6SocketKey) (st *tcp.Stream, err protocol.Error) {
	// TODO:::
	tcpSockets[sKey] = st
	return
}
