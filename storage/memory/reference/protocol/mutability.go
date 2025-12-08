/* For license and copyright information please see the LEGAL file in the code repository */

package reference_p

import (
	capsule_p "memar/computer/capsule/protocol"
	// error_p "memar/process/error/protocol"
)

// Mutable
//
// Other protocols:::
// https://doc.rust-lang.org/rust-by-example/scope/borrow/mut.html
type Mutable[T any] interface {
	capsule_p.Assignment[T]
}

// Immutable
//
// Other protocols:::
// https://en.wikipedia.org/wiki/Immutable_object
type Immutable interface {
	// TODO::: Do we need runtime immutable mechanism?? If yes how about undo this action?? It seems can add many complex codes to codebases.
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/freeze
	// Freeze() (err error_p.Error)
	// Immutably() (err error_p.Error)
	// IsFrozen() (froze bool, err error_p.Error)
	// TODO::: immutable embedded capsule??
	// DeepFreeze() (err error_p.Error)
}
