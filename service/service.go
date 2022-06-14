/* For license and copyright information please see LEGAL file in repository */

package service

import (
	"../mediatype"
	"../protocol"
)

// Service store needed data for a service to implement protocol.Service when embed to other struct that implements other methods!
type Service struct {
	uri string // Fill just if any http like type handler needed! Simple URI not variable included! API services can set like "/m?{{.ServiceID}}" but it is not efficient, find services by ID.

	priority protocol.Priority // Use to queue requests by its priority
	weight   protocol.Weight   // Use to queue requests by its weights in the same priority

	// Authorization data to authorize incoming service
	crud     protocol.CRUD // CRUD == Create, Read, Update, Delete
	userType protocol.UserType

	mediatype.MediaType
}

// func (s *Service) Init() {}

func (s *Service) SetURIRoutePath(uri string) { s.uri = uri }
func (s *Service) SetPriority(priority protocol.Priority, weight protocol.Weight) {
	s.priority = priority
	s.weight = weight
}
func (s *Service) SetAuthorization(crud protocol.CRUD, userType protocol.UserType) {
	s.crud = crud
	s.userType = userType
}

func (s *Service) URI() string                 { return s.uri }
func (s *Service) Priority() protocol.Priority { return s.priority }
func (s *Service) Weight() protocol.Weight     { return s.weight }
func (s *Service) CRUDType() protocol.CRUD     { return s.crud }
func (s *Service) UserType() protocol.UserType { return s.userType }

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
