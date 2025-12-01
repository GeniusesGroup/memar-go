/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

import (
	error_p "memar/process/error/protocol"
)

// Assignment use to set or store other capsule to current one.
// If a capsule not provide this interface it means it is `Readonly`.
// 
// https://en.wikipedia.org/wiki/Assignment_(computer_science)
type Assignment[T any] interface {
	// It will check(validate) given value and return proper error for
	Set(new T) (err error_p.Error)

	// STORE use in assembly languages
	// STORE()
	// STR()

	// ReSet()
}

type Assignment_DefaultValue interface {
	// SetDefaultValue CAN return some types of error e.g. READONLY capsule.
	SetDefaultValue() (err error_p.Error)
}
