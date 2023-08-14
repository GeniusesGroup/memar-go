/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	er "memar/error"
)

var (
	ErrNotFound er.Error

	// This conditions must be checked just in the dev phase.
	ErrServiceNotProvideIdentifier er.Error
	ErrServiceDuplicateIdentifier  er.Error
)

func init() {
	ErrNotFound.Init("domain/memar.scm.geniuses.group; type=error; package=service; name=not_found")

	ErrServiceNotProvideIdentifier.Init("domain/memar.scm.geniuses.group; type=error; package=service; name=service_not_provide_identifier")
	ErrServiceDuplicateIdentifier.Init("domain/memar.scm.geniuses.group; type=error; package=service; name=service_duplicate_identifier")
}
