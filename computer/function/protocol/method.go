/* For license and copyright information please see the LEGAL file in the code repository */

package function_p

import (
	datatype_p "memar/computer/datatype/protocol"
)

type Field_Method interface {
	Method() Method
}

type Field_Methods interface {
	Methods() []Method
}

// Method
type Method interface {
	Function

	Receiver() datatype_p.DataType // std/go/ast.FuncType.TypeParams
}
