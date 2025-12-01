/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

import (
	error_p "memar/process/error/protocol"
)

type Accessor[T any] interface {
	Get() T

	// LOAD use in assembly languages
	// LOAD()
	// LDR()
}

type DefaultValue[T any] interface {
	DefaultValue() T
}

// Other protocols:::
// https://www.typescriptlang.org/docs/handbook/utility-types.html#requiredtype
type Optional interface {
	// Base on data or function false means:
	// - data required and must be exist.
	// 		e.g. user gender is optional, user email is not optional
	// - Function (usually a service) MUST be successful in chain of calls or MUST rollback all calls.
	// 		e.g. SendSMS is optional but WithdrawMoney is not optional in TransferMoney chain
	// 		SendSMS is not optional in SendOTP chain
	Optional() bool
}
