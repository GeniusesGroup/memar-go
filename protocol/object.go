/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Object interface {
	Fields() []DataType
	Methods() []DataType_Method

	// DataType
	// ObjectLifeCycle
}
