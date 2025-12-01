/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

import (
	function_p "memar/computer/function/protocol"
	datatype_p "memar/computer/datatype/protocol"
)

type Capsule interface {
	Fields() []datatype_p.DataType
	Methods() []function_p.Method

	// datatype_p.DataType
	// LifeCycle
}
