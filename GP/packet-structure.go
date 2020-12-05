/* For license and copyright information please see LEGAL file in repository */

package gp

// packetStructure is represent protocol structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/GP.md
type packetStructure struct {
	//DestinationGPAddr[14]byte
	DestinationSociety [4]byte // uint32
	DestinationRouter  [4]byte // uint32
	DestinationUser    [4]byte // uint32
	DestinationApp     [2]byte // uint16

	//SourceGPAddr[14]byte
	SourceSociety [4]byte // uint32
	SourceRouter  [4]byte // uint32
	SourceUser    [4]byte // uint32
	SourceApp     [2]byte // uint16

	PayloadLength [2]byte // uint16
	StreamID      [2]byte // uint16
	PacketID      [2]byte // uint16 also can act as offset

	Payload   []byte
	Signature [32]byte // Checksum
	Padding   []byte   // If needed
}
