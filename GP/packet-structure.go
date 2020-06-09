/* For license and copyright information please see LEGAL file in repository */

package gp

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/GP.md
type packetStructure struct {
	DestinationSociety     [4]byte  // uint32
	DestinationRouter      [4]byte  // uint32
	DestinationUser        [4]byte  // uint32
	DestinationApp         [2]byte  // uint16
	DestinationAppProtocol [2]byte  // uint16
	SourceGP               [16]byte // 4+4+4+2+2
	PayloadLength          uint16
	StreamID               uint32
	PacketID               uint32
	Payload                []byte
	Padding                []byte
	Checksum               [32]byte
}
