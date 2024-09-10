/* For license and copyright information please see the LEGAL file in the code repository */

package errors_p

import (
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
)

// Errors use to register errors to get them in a desire way e.g. ErrorID in http headers.
type Errors interface {
	Register(errorToRegister error_p.Error) (err error_p.Error)
	GetByID(id datatype_p.ID) (err error_p.Error)
	GetByMediaType(mt string) (err error_p.Error)
}
