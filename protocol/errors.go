/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Errors use to register errors to get them in a desire way e.g. ErrorID in http headers.
type Errors interface {
	Register(errorToRegister Error) (err Error)
	GetByID(id MediaTypeID) (err Error)
	GetByMediaType(mt string) (err Error)
}
