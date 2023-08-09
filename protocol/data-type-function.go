/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// DataType_Function
type DataType_Function interface {
	// We believe fields MUST always access from inside the object,
	// So we MUST have this method just in methods not fields.
	Access() DataType_Access

	WhenToExecuted() // runtime, compile time

	Blocking() bool // TODO::: be method or as type??
	// TODO::: add more

	DataType
}
