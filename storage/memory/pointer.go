/* For license and copyright information please see the LEGAL file in the code repository */

package memory

import (
	"memar/math/boolean"
)

type PTR uintptr

//memar:impl memar/math/logic/protocol.Equivalence
func (p *PTR) Equal(with PTR) boolean.Boolean {
	return *p == with
}
