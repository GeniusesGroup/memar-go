/* For license and copyright information please see LEGAL file in repository */

package ipv6

// Packet use for packet methods!
type Packet []byte

const (
	// Version of protocol
	Version = 6
	// HeaderLen is minimum header length of IPv6 header
	HeaderLen = 40
)

// CheckPacket will check packet for any bad situation!
// Always check packet before use any other packet methods otherwise panic may occur!
func (p *Packet) CheckPacket() (err error) {
	if len(*p) < HeaderLen {
		return PacketTooShort
	}

	return nil
}

// GetVersion will return packet version in memory safe way!
func (p *Packet) GetVersion() (Version uint8) {
	return uint8((*p)[0]) >> 4
}

// GetTrafficClass will return packet TrafficClass in memory safe way!
func (p *Packet) GetTrafficClass() (TrafficClass uint8) {
	return uint8((*p)[0]&0x0f)<<4 | uint8((*p)[1])>>4
}

// GetFlowLabel will return packet FlowLabel in memory safe way!
func (p *Packet) GetFlowLabel() (FlowLabel uint32) {
	return uint32((*p)[1]&0x0f)<<16 | uint32((*p)[2])<<8 | uint32((*p)[3])
}

// GetPayloadLength will return packet PayloadLength in memory safe way!
func (p *Packet) GetPayloadLength() (PayloadLength uint16) {
	return uint16((*p)[4]) | uint16((*p)[5])<<8
}

// GetNextHeader will return packet NextHeader in memory safe way!
func (p *Packet) GetNextHeader() (NextHeader uint8) {
	return uint8((*p)[6])
}

// GetHopLimit will return packet HopLimit in memory safe way!
func (p *Packet) GetHopLimit() (HopLimit uint8) {
	return uint8((*p)[7])
}

// GetSourceIP will return SourceIPAddress in memory safe way!
func (p *Packet) GetSourceIP() (SourceIPAddress [16]byte) {
	copy(SourceIPAddress[:], (*p)[8:24])
	return SourceIPAddress
}

// GetDestinationIP will return DestinationIPAddress in memory safe way!
func (p *Packet) GetDestinationIP() (DestinationIPAddress [16]byte) {
	copy(DestinationIPAddress[:], (*p)[24:40])
	return DestinationIPAddress
}

// GetPayload will return Payload in memory safe way!
func (p *Packet) GetPayload() (Payload []byte) {
	// TODO : Check first payload location by GetNextHeader method!
	return (*p)[40:]
}
