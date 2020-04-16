/* For license and copyright information please see LEGAL file in repository */

package ipv6

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
type packetStructure struct {
	Version              uint8
	TrafficClass         uint8
	FlowLabel            uint32
	PayloadLength        uint16
	NextHeader           uint8 // https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
	HopLimit             uint8
	SourceIPAddress      Addr
	DestinationIPAddress Addr
	Payload              []byte
}
