/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	error_p "memar/error/protocol"
)

func ToGoError(err error_p.Error) error {
	if err == nil {
		return nil
	}

	var errStr errorString
	errStr.msg = err.Summary()
	return &errStr
}

// errorString is a trivial implementation of error.
type errorString struct {
	msg string
}

func (e *errorString) Error() string { return e.msg }

func ToError(err error) error_p.Error {
	if err == nil {
		return nil
	}

	var exErr = err.(error_p.Error)
	if exErr != nil {
		return exErr
	}

	// TODO:::
	return nil
}
