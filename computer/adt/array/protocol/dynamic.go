/* For license and copyright information please see the LEGAL file in the code repository */

package array_p

import (
	container_p "memar/adt/container/protocol"
)

// Slice is dynamically size array
// https://en.wikipedia.org/wiki/Dynamic_array
type Dynamic[ELEMENT container_p.Element] interface {
	container_p.Container[ELEMENT]
}
