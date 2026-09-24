/* For license and copyright information please see the LEGAL file in the code repository */

package container

import (
	container_p "memar/computer/adt/container/protocol"
)

type Container[EL container_p.Element] struct {
	container []EL
}
