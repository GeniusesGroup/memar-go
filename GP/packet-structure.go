/* For license and copyright information please see LEGAL file in repository */

package gp

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/GP.md
type packetStructure struct {
	DestinationXP          [4]byte
	DestinationRouter      [4]byte
	DestinationUser        [4]byte
	DestinationApp         [2]byte
	DestinationAppProtocol [2]byte
	SourceGP               [16]byte
	PayloadLength          uint16
	StreamID               uint32
	PacketID               uint32
	Payload                []byte
	Padding                []byte
	Checksum               [32]byte
}
