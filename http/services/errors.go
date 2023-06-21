/* For license and copyright information please see the LEGAL file in the code repository */

package hs

import (
	er "libgo/error"
)

// Declare package errors
var (
	ErrNoConnection         er.Error
	ErrNotFound             er.Error
	ErrUnsupportedMediaType er.Error

	ErrBadHost er.Error
)

func init() {
	ErrNoConnection.Init("domain/httpwg.org; type=error; name=no-connection")
	ErrNotFound.Init("domain/httpwg.org; type=error; name=not-found")
	ErrUnsupportedMediaType.Init("domain/httpwg.org; type=error; name=unsupported-media-type")

	ErrBadHost.Init("domain/httpwg.org; type=error; name=unsupported-media-type")
}
