/* For license and copyright information please see the LEGAL file in the code repository */

package time_p

import (
	datatype_p "memar/computer/datatype/protocol"
	logic_p "memar/math/logic/protocol"
	string_p "memar/codec/string/protocol"
)

type Field_Epoch interface {
	Epoch() Epoch
}

// Epoch is the interface that must implement by any time capsule.
// It is base on Epoch and Second terms to work anywhere (in any planet in the universe).
// https://en.wikipedia.org/wiki/Epoch
type Epoch interface {
	datatype_p.DataType
	string_p.Stringer_To

	logic_p.Equivalence[Epoch]
}
