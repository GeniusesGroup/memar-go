/* For license and copyright information please see the LEGAL file in the code repository */

package validation_p

import (
	error_p "memar/error/protocol"
)

// TODO::: why not validate per field??

// https://martinfowler.com/apsupp/spec.pdf
// https://github.com/nullexp/specp/blob/main/example/basic.go
type Satisfier[T any] interface {
	IsSatisfiedBy(value T) error_p.Error
}

type Specification[T any] interface {
	Satisfier[T]
	And(other Specification[T]) Specification[T]
	AndNot(other Specification[T]) Specification[T]
	Not() Specification[T]
	OrNot(other Specification[T]) Specification[T]
	Or(other Specification[T]) Specification[T]
}
