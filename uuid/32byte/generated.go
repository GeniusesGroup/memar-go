/* For license and copyright information please see the LEGAL file in the code repository */

package uuid

import "memar/protocol"

type Generated struct {
	UID
	id         protocol.ID
	idAsString string
}

func (g *Generated) NewHashString(data string) {
	g.NewHash((unsafeStringToByteSlice(data)))
	g.id = g.UID.ID()
	g.idAsString = g.UID.IDasString()
}

func (g *Generated) ID() protocol.ID    { return g.id }
func (g *Generated) IDasString() string { return g.idAsString }
