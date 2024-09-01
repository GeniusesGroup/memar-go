/* For license and copyright information please see the LEGAL file in the code repository */

package errors_p

import (
	"memar/protocol"
)

// Errors use to register errors to get them in a desire way e.g. ErrorID in http headers.
type Errors interface {
	Register(errorToRegister protocol.Error) (err protocol.Error)
	GetByID(id protocol.DataTypeID) (err protocol.Error)
	GetByMediaType(mt string) (err protocol.Error)
}
