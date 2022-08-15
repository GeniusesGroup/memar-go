/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	"github.com/GeniusesGroup/libgo/detail"
	"github.com/GeniusesGroup/libgo/mediatype"
	"github.com/GeniusesGroup/libgo/protocol"
)

// New return new Error that implement protocol.Error
// Never change MediaType due to it adds unnecessary complicated troubleshooting errors on SDK.
// TODO::: escapes to heap problem of return value, How prevent it??
// func New(mediatype string) (err Error) { err.Init(mediatype); return }

// Err is the same as the Error.
// Use this type when embed in other struct to solve field & method same name problem(Error struct and Error() method) to satisfy interfaces.
type Err = Error

// Error implements protocol.Error
type Error struct {
	internal  bool
	temporary bool

	detail.DS
	mediatype.MT
}

// Init initialize Error that implement protocol.Error
// Never change MediaType due to it adds unnecessary complicated troubleshooting errors on SDK.
func (e *Error) Init(mediatype string) {
	e.MT.Init(mediatype)

	// RegisterError will register in the application.
	// Force to check by runtime check, due to testing package not let us by any const!
	if protocol.App != nil {
		protocol.App.RegisterError(e)
	}
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

func (e *Error) Internal() bool  { return e.internal }
func (e *Error) Temporary() bool { return e.temporary }

func (e *Error) SetInternal()  { e.internal = true }
func (e *Error) SetTemporary() { e.temporary = true }

func (e *Error) Notify() {
	// TODO:::
}

// Go compatibility methods. Unwrap provides compatibility for Go 1.13 error chains.
func (e *Error) Error() string { return e.MT.MediaType() }
func (e *Error) Cause() error  { return nil }
func (e *Error) Unwrap() error { return nil }
// func (e *Error) Is(error) bool
// func (e *Error) As(any) bool 
