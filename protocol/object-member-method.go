/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// DataType_Method
type DataType_Method interface {
	DataType_Function

	Receiver() DataType // std/go/ast.FuncType.TypeParams
}
