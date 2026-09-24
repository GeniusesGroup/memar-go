/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	marker_p "memar/computer/adt/marker/protocol"
	logic_p "memar/math/logic/protocol"
)

// Character is base of any character encoding
type Character interface {
	logic_p.Equivalence[Character]

	marker_p.Clone[Character]
	marker_p.Copy[Character]
}
