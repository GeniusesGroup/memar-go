/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	errors_errs "memar/process/errors/errors"
)

func init() {
	Register(&errors_errs.ErrNotFound)
	Register(&errors_errs.ErrNotExist)

	Register(&errors_errs.ErrNotProvideIdentifier)
	Register(&errors_errs.ErrDuplicateIdentifier)
}
