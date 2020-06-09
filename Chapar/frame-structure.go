/* For license and copyright information please see LEGAL file in repository */

package chapar

// frameStructure represents Chapar frame structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/Chapar.md
// up-to 255 switch port number can be in a frame header!
type frameStructure struct {
	NextHop          byte
	HopCount         byte
	NextHeader       byte
	FirstHopPortNum  byte
 // SecondHopPortNum byte
 // ...              byte
	Payload          []byte
}
