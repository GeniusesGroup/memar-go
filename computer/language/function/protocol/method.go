/* For license and copyright information please see the LEGAL file in the code repository */

package function_p

import (
	datatype_p "memar/datatype/protocol"
)

// Method
type Method interface {
	Function

	Receiver() datatype_p.DataType // std/go/ast.FuncType.TypeParams
}
