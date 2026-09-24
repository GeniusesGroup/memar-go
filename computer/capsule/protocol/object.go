/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

import (
	adt_p "memar/computer/adt/protocol"
	datatype_p "memar/computer/datatype/protocol"
	function_p "memar/computer/function/protocol"
)

type Capsule interface {
	Field_Fields
	function_p.Field_Methods

	datatype_p.DataType
	// LifeCycle
	// Authorization

	// logic_p.Equivalence[Capsule]
	adt_p.ADT
}
