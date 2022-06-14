/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"../mediatype"
	"../protocol"
)

// Error implements protocol.Error
// Never change MediaType due to it adds unnecessary complicated troubleshooting errors on SDK.
type Error struct {
	internal  bool
	temporary bool

	mediatype.MediaType
}

// func (e *Error) Init() {}

// RegisterError will register in the application.
func (e *Error) RegisterError() {
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
	return false
}

func (e *Error) Internal() bool   { return e.internal }
func (e *Error) Temporary() bool  { return e.temporary }
func (e *Error) ToString() string { return e.IDasString() }

func (e *Error) SetInternal()  { e.internal = true }
func (e *Error) SetTemporary() { e.temporary = true }

func (e *Error) Notify() {
	// TODO:::
}
