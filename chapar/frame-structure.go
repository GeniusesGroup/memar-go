/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

// frameFormat represents Chapar frame structure.
// It is just to show protocol in better way, we never use this type.
// Read more about this protocol : https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md
type frameFormat struct {
	Header  frameHeaderFormat // Up to 258 Byte length
	Payload []byte            // Up to 7934 Byte length
}

// It is just to show protocol in better way, we never use this type.
// up-to 255 switch port number can be in a frame header.
// First hopNum is hopNum==1 not hopNum==0. Don't read hopNum==0 due to it is use for broadcast frame.
type frameHeaderFormat struct {
	NextHop          byte // next hop number.
	HopCount         byte // the number of intermediate network devices indicate in frame.
	NextHeader       byte // next header ID of the frame.
	FirstHopPortNum  byte //
 // SecondHopPortNum byte
 // ...              byte
}
