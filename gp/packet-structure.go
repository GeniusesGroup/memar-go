/* For license and copyright information please see LEGAL file in repository */

package gp

// packetStructure is represent protocol structure.
// It is just to show protocol in better way, we never use this type.
// Read more about this protocol : https://github.com/GeniusesGroup/RFCs/blob/master/Giti-Network.md
type packetStructure struct {
	//DestinationGPAddr[16]byte
	DestinationPlanet  [2]byte // uint16
	DestinationSociety [4]byte // uint32
	DestinationRouter  [4]byte // uint32
	DestinationUser    [4]byte // uint32
	DestinationApp     [2]byte // uint16

	//SourceGPAddr[16]byte
	SourcePlanet  [2]byte // uint16
	SourceSociety [4]byte // uint32
	SourceRouter  [4]byte // uint32
	SourceUser    [4]byte // uint32
	SourceApp     [2]byte // uint16

	PacketNumber [8]byte // uint64
	Frames       []byte
	Signature    []byte // Checksum, MAC, Tag, ...
}
