/* For license and copyright information please see LEGAL file in repository */

package chapar

// chapar represents Chapar frame structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md
type chapar struct {
	Header  header // Up to 258 Byte length!
	Payload []byte // Up to 7934 Byte length!
}

// It is just to show protocol in better way, we never use this type!
// up-to 255 switch port number can be in a frame header!
type header struct {
	NextHop          byte
	HopCount         byte
	NextHeader       byte
	FirstHopPortNum  byte
 // SecondHopPortNum byte
 // ...              byte
	Payload          []byte
}
