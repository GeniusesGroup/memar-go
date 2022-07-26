/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

func ToGoError(err protocol.Error) error {
	var exErr = err.(*Error)
	if exErr != nil {
		return exErr
	}
	return &errorString{msg: err.ToString()}
}

// errorString is a trivial implementation of error.
type errorString struct {
	msg string
}

func (e *errorString) Error() string { return e.msg }

func ToError(err error) protocol.Error {
	// TODO:::
	return nil
}
