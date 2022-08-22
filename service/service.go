/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

// Service implement protocol.Service when embed to other struct that implements other needed methods.
type Service struct{}

//libgo:impl protocol.Service
func (s *Service) URI() string                 { return "" }
func (s *Service) Priority() protocol.Priority { return protocol.Priority_Unset }
func (s *Service) Weight() protocol.Weight     { return protocol.Weight_Unset }
func (s *Service) CRUDType() protocol.CRUD     { return protocol.CRUD_None }
func (s *Service) UserType() protocol.UserType { return protocol.UserType_Unset }

//libgo:impl protocol.ServiceDetails
func (s *Service) Request() []protocol.Field  { return nil }
func (s *Service) Response() []protocol.Field { return nil }

/*
*********** Handlers ***********
not-implemented handlers of the service.
*/

func (s *Service) ServeSRPC(st protocol.Stream) (err protocol.Error) {
	err = &ErrServiceNotAcceptSRPC
	return
}
func (s *Service) ServeSRPCDirect(conn protocol.Connection, request []byte) (response []byte, err protocol.Error) {
	err = &ErrServiceNotAcceptSRPCDirect
	return
}
func (s *Service) ServeHTTP(st protocol.Stream, httpReq protocol.HTTPRequest, httpRes protocol.HTTPResponse) (err protocol.Error) {
	err = &ErrServiceNotAcceptHTTP
	return
}
