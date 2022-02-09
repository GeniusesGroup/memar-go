/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"fmt"

	"../mediatype"
	"../protocol"
)

// New returns a new error
// "urn:giti:{{domain-name}}:error:{{error-name}}"
func New(mediatype *mediatype.MediaType) *Error {
	if mediatype.ID() == 0 {
		// This condition will just be true in the dev phase.
		panic("Error must have valid ID to save it in platform errors pools")
	}

	var err = Error{
		id:        mediatype.ID(),
		mediatype: mediatype,
	}

	// Save finalize needed logic on given error and register in the application
	err.updateStrings()
	// Force to check by runtime check, due to testing package not let us by any const!
	if protocol.App != nil {
		protocol.App.RegisterError(&err)
	}
	return &err
}

// GetID return id of error if err id exist.
func GetID(err error) uint64 {
	if err == nil {
		return 0
	}
	var exErr *Error = err.(*Error)
	if exErr != nil {
		return exErr.mediatype.ID()
	}
	// if error not nil but not Error, pass biggest number!
	return 18446744073709551615
}

// Error is a extended implementation of error.
// Never change urn due to it adds unnecessary complicated troubleshooting errors on SDK.
type Error struct {
	id        uint64
	mediatype *mediatype.MediaType

	stringMethod string
	errorMethod  string
}

func (e *Error) ID() uint64                    { return e.id }
func (e *Error) MediaType() protocol.MediaType { return e.mediatype }
func (e *Error) ToString() string              { return e.stringMethod }

// Equal compare two Error.
func (e *Error) Equal(err protocol.Error) bool {
	if e == nil && err == nil {
		return true
	}
	if e != nil && err != nil && e.mediatype.ID() == err.MediaType().ID() {
		return true
	}
	return false
}

// IsEqual compare two error.
func (e *Error) IsEqual(err protocol.Error) bool {
	var exErr = err.(*Error)
	if exErr != nil && e.mediatype.ID() == exErr.mediatype.ID() {
		return true
	}
	return false
}

// Go compatibility methods. Unwrap provides compatibility for Go 1.13 error chains.
func (e *Error) Error() string { return e.errorMethod }
func (e *Error) Cause() error  { return e }
func (e *Error) Unwrap() error { return e }

func (e *Error) updateStrings() {
	e.stringMethod = "Error ID: " + e.mediatype.IDasString()
	var localeDetail = e.mediatype.Detail(protocol.AppLanguage)
	e.errorMethod = fmt.Sprintf("Error ID: %s\n	Summary: %s\n	Overview: %s\n", e.mediatype.IDasString(), localeDetail.Summary(), localeDetail.Overview())
}
