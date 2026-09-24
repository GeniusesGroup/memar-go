/* For license and copyright information please see the LEGAL file in the code repository */

package container

import (
	container_p "memar/computer/adt/container/protocol"
	error_p "memar/process/error/protocol"
)

//memar:impl memar/computer/adt/container/protocol.GetElement
func (self *Container[EL]) GetElement(i container_p.ElementIndex) (el EL, err error_p.Error) {
	// TODO::: Check i not more than self.OccupiedLength()
	// if !exist {
	// err = TODO:::
	// }

	el = self.container[i]
	return
}
