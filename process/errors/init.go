/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	errors_errs "memar/process/errors/errors"
)

func init() {
	Register(&errors_errs.NotFound)
	Register(&errors_errs.NotExist)

	Register(&errors_errs.NotProvideIdentifier)
	Register(&errors_errs.DuplicateIdentifier)
}
