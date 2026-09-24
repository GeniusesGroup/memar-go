/* For license and copyright information please see the LEGAL file in the code repository */

package container

import (
	container_p "memar/computer/adt/container/protocol"
	// error_p "memar/process/error/protocol"
)

//memar:impl memar/computer/adt/protocol.OccupiedLength
func (self *Container[EL]) OccupiedLength() container_p.NumberOfElement {
	return container_p.NumberOfElement(len(self.container))
}
