/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	logic_p "memar/math/logic/protocol"
	memory_p "memar/storage/memory/protocol"
)

// Character is base of any character encoding
type Character interface {
	logic_p.Equivalence[Character]
	memory_p.Clone[Character]
	memory_p.Copy[Character]
}
