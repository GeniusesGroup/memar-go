/* For license and copyright information please see LEGAL file in repository */

package ipv6

import (
	"../protocol"
)

/*
// Routing is IPv6 extension header with NextHeader==43
type Routing struct {
	NextHeader       uint8
	HdrExtLen        uint8
	RoutingType      uint8
	SegmentsLeft     uint8
	TypeSpecificData [4]byte
	Optional         []byte // more type-specific data...
}
*/
type extensionRouting []byte

func (er extensionRouting) NextHeader() byte          { return o[0] }
func (er extensionRouting) Length() byte              { return o[1] }
func (er extensionRouting) RoutingType() byte         { return o[2] }
func (er extensionRouting) SegmentsLeft() byte        { return o[3] }
func (er extensionRouting) TypeSpecificData() [4]byte { return o[4:] }
func (er extensionRouting) Optional() []byte          { return o[8:] }
func (er extensionRouting) NextHeader() []byte        { return o[8+er.Length():] }

func (er extensionRouting) Process(conn *Connection) (err protocol.Error) {
	return
}
