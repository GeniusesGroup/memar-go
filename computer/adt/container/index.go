/* For license and copyright information please see the LEGAL file in the code repository */

package container

import (
	container_p "memar/computer/adt/container/protocol"
)

//memar:impl memar/computer/adt/container/protocol.Field_LastElementIndex
func (self *Container[EL]) LastElementIndex() container_p.ElementIndex {
	return container_p.ElementIndex(self.OccupiedLength() - 1)
}
