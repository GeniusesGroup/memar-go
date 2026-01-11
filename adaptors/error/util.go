/* For license and copyright information please see the LEGAL file in the code repository */

package error_adaptor

import (
	error_p "memar/process/error/protocol"
)

func ToGoError(memarErr error_p.Error) error {
	if memarErr == nil {
		return nil
	}

	var errStr errorString
	errStr.msg = memarErr.Summary()
	return &errStr
}

// errorString is a trivial implementation of error.
type errorString struct {
	msg string
}

func (self *errorString) Error() string { return self.msg }

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
