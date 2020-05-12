/* For license and copyright information please see LEGAL file in repository */

package chapar

// frameStructure is represent switching frame structure!
// It is just to show protocol in better way, we never use this type!
// Read more about this protocol : https://github.com/SabzCity/RFCs/blob/master/Chapar.md
// up-to 256 switch port number can be here in frame!
type frameStructure struct {
	NextHop        byte
	TotalHop       byte
	NextHeader     byte
	Switch1PortNum byte
 // Switch2PortNum byte
 // ...            byte
	Payload        []byte
}
