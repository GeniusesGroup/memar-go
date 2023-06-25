/* For license and copyright information please see the LEGAL file in the code repository */

package ipv6

/*
	Routing is IPv6 extension header with NextHeader==43
	type Routing struct {
		NextHeader       byte
		HdrExtLen        byte
		RoutingType      byte
		SegmentsLeft     byte
		TypeSpecificData [4]byte
		Optional         []byte // more type-specific data...
	}
*/
type extensionRouting []byte

func (er extensionRouting) NextHeader() byte               { return er[0] }
func (er extensionRouting) Length() byte                   { return er[1] }
func (er extensionRouting) RoutingType() byte              { return er[2] }
func (er extensionRouting) SegmentsLeft() byte             { return er[3] }
func (er extensionRouting) TypeSpecificData() (sd [4]byte) { copy(sd[:], er[4:]); return }
func (er extensionRouting) Optional() []byte               { return er[8:er.Length()] }
func (er extensionRouting) NextFrame() []byte              { return er[8+er.Length():] }

// func (er extensionRouting) Process(conn *Connection) (err protocol.Error) {
// 	return
// }
