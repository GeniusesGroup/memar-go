/* For license and copyright information please see the LEGAL file in the code repository */

package memar_uuid

import (
	datatype_p "memar/computer/datatype/protocol"
)

type Hash32_Generated struct {
	Hash32
	id         datatype_p.ID
	idAsString string
}

func (self *Hash32_Generated) NewHashString(data string) {
	self.NewHash((unsafeStringToByteSlice(data)))
	self.id = self.Hash32.ID()
	self.idAsString = self.Hash32.IDasString()
}

func (self *Hash32_Generated) ID() datatype_p.ID  { return self.id }
func (self *Hash32_Generated) IDasString() string { return self.idAsString }
