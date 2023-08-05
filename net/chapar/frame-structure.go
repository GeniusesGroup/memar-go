/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

// frameFormat represents Chapar frame structure.
// It is just to show protocol in better way, we never use this type.
// Read more about this protocol : https://github.com/GeniusesGroup/memar/blob/main/networking-osi_2-Chapar.md
// It is just to show protocol in better way, we never use this type.
// up-to 255 switch port number can be in a frame header.
// First hopNum is hopNum==1 not hopNum==0. Don't read hopNum==0 due to it is use for broadcast frame.
type frameFormat struct {
	FrameType       byte //
	HopCount        byte // the number of intermediate network devices indicate in frame.
	NextHop         byte // next hop number.
	FirstHopPortNum byte //
// SecondHopPortNum byte
// ...              byte
}
