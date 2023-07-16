/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"libgo/protocol"
)

// frameStructure is represent protocol frame structure.
// It is just to show protocol in better way, we never use this type.
// Read more about this protocol : https://github.com/GeniusesGroup/RFCs/blob/master/networking-osi_3-Giti-Network.md
type frameStructure struct {
	FrameID protocol.Network_FrameID //

	// DestinationGPAddr Addr
	DestinationPlanet  [2]byte // uint16
	DestinationSociety [4]byte // uint32
	DestinationRouter  [4]byte // uint32
	DestinationUser    [4]byte // uint32
	DestinationApp     [2]byte // uint16

	//SourceGPAddr Addr
	SourcePlanet  [2]byte // uint16
	SourceSociety [4]byte // uint32
	SourceRouter  [4]byte // uint32
	SourceUser    [4]byte // uint32
	SourceApp     [2]byte // uint16
}
