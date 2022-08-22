/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	er "github.com/GeniusesGroup/libgo/error"
)

// Declare package errors
var (
	ErrNotFound                   er.Error
	ErrServiceNotAcceptSRPC       er.Error
	ErrServiceNotAcceptSRPCDirect er.Error
	ErrServiceNotAcceptHTTP       er.Error
)

func init() {
	ErrNotFound.Init("domain/geniuses.group; type=error; package=service; name=not_found")
	ErrServiceNotAcceptSRPC.Init("domain/geniuses.group; type=error; package=service; name=service_not_accept_srpc")
	ErrServiceNotAcceptSRPCDirect.Init("domain/geniuses.group; type=error; package=service; name=service_not_accept_direct_srpc")
	ErrServiceNotAcceptHTTP.Init("domain/geniuses.group; type=error; package=service; name=service_not_accept_http")
}
