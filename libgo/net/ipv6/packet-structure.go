/* For license and copyright information please see the LEGAL file in the code repository */

package ipv6

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
// https://en.wikipedia.org/wiki/IPv6_packet
type packetStructure struct {
	//  version of given IPv6 packet.
	Version uint8
	// traffic class of given IPv6 packet.
	TrafficClass uint8
	//  flow label of given IPv6 packet.
	FlowLabel [3]byte
	// payload length of given IPv6 packet.
	PayloadLength uint16
	// next header of given IPv6 packet.
	NextHeader uint8 // https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
	// hop limit of given IPv6 packet.
	HopLimit uint8
	// source IP address of given IPv6 packet.
	SourceIPAddress Addr
	// destination IP address of given IPv6 packet.
	DestinationIPAddress Addr
	//  payload of given IPv6 packet.
	Payload []byte
}
