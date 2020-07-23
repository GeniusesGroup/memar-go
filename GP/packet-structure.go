/* For license and copyright information please see LEGAL file in repository */

package gp

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/GP.md
type packetStructure struct {
 // DestinationGP          [16]byte
	DestinationSociety     [4]byte  // uint32
	DestinationRouter      [4]byte  // uint32
	DestinationUser        [4]byte  // uint32
	DestinationApp         [2]byte  // uint16
	DestinationProtocol    [2]byte  // uint16

 // SourceGP               [16]byte
	SourceSociety          [4]byte  // uint32
	SourceRouter           [4]byte  // uint32
	SourceUser             [4]byte  // uint32
	SourceApp              [2]byte  // uint16
	SourceProtocol         [2]byte  // uint16

	PayloadLength          [2]byte  // uint16
	StreamID               [4]byte  // uint32
	PacketID               [4]byte  // uint32

	Payload                []byte
	Padding                []byte
	Checksum               [32]byte
}
