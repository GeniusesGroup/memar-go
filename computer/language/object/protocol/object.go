/* For license and copyright information please see the LEGAL file in the code repository */

package object_p

import (
	function_p "memar/computer/language/function/protocol"
	datatype_p "memar/datatype/protocol"
)

type Object interface {
	Fields() []datatype_p.DataType
	Methods() []function_p.Method

	// datatype_p.DataType
	// LifeCycle
}
