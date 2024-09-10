/* For license and copyright information please see the LEGAL file in the code repository */

package function_p

import (
	datatype_p "memar/datatype/protocol"
)

// Function
type Function interface {
	// We believe fields MUST always access from inside the object,
	// So we MUST have this method just in methods not fields.
	Access() Access

	WhenToExecuted() // runtime, compile time

	Blocking() bool // TODO::: be method or as type??
	// TODO::: add more

	datatype_p.DataType
}
