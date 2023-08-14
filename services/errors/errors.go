/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	er "libgo/error"
)

// Declare package errors
var (
	ErrNotFound er.Error

	ErrServiceNotAcceptSRPC       er.Error
	ErrServiceNotAcceptSRPCDirect er.Error
	ErrServiceNotAcceptHTTP       er.Error

	// This conditions must be true just in the dev phase.
	ErrServiceNotProvideIdentifier er.Error
	ErrServiceDuplicateIdentifier  er.Error
)

func init() {
	ErrNotFound.Init("domain/libgo.scm.geniuses.group; type=error; package=service; name=not_found")

	ErrServiceNotAcceptSRPC.Init("domain/libgo.scm.geniuses.group; type=error; package=service; name=service_not_accept_srpc")
	ErrServiceNotAcceptSRPCDirect.Init("domain/libgo.scm.geniuses.group; type=error; package=service; name=service_not_accept_direct_srpc")
	ErrServiceNotAcceptHTTP.Init("domain/libgo.scm.geniuses.group; type=error; package=service; name=service_not_accept_http")

	ErrServiceNotProvideIdentifier.Init("domain/libgo.scm.geniuses.group; type=error; package=service; name=service_not_provide_identifier")
	ErrServiceDuplicateIdentifier.Init("domain/libgo.scm.geniuses.group; type=error; package=service; name=service_duplicate_identifier")
}
