/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	"memar/datatype"
	"memar/mediatype"
	"memar/protocol"
)

// Err is the same as the Error.
// Use this type when embed in other struct to solve field & method same name problem(Error struct and Error() method) to satisfy interfaces.
type Err = Error

// Error implements protocol.Error
type Error struct {
	ErrorType

	mediatype.MT
	datatype.DataType
}

// Init initialize Error that implement protocol.Error
// Never change MediaType due to it adds unnecessary complicated troubleshooting errors on SDK.
//
//memar:impl memar/protocol.ObjectLifeCycle
func (e *Error) Init(mediatype string) (err protocol.Error) {
	err = e.MT.Init(mediatype)
	// if err != nil {
	// 	return
	// }
	return
}

// Equal compare two Error.
func (e *Error) Equal(err protocol.Error) bool {
	if e == nil && err == nil {
		return true
	}
	if e != nil && err != nil && e.ID() == err.ID() {
		return true
	}
	// TODO::: check err as chain error
	return false
}

func (e *Error) Type() protocol.ErrorType             { return protocol.ErrorType(e.ErrorType) }
func (e *Error) CheckType(et protocol.ErrorType) bool { return e.ErrorType.Check(et) }

func (e *Error) Notify() {
	// TODO:::
}

// Go compatibility methods. Unwrap provides compatibility for Go 1.13 error chains.
func (e *Error) Error() string { return e.MT.MediaType() }
func (e *Error) Cause() error  { return nil }
func (e *Error) Unwrap() error { return nil }

// func (e *Error) Is(error) bool
// func (e *Error) As(any) bool
