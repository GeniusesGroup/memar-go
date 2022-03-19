/* For license and copyright information please see LEGAL file in repository */

package service

import (
	"../mediatype"
	"../protocol"
)

// Service store needed data for a service to implement protocol.Service when embed to other struct that implements other methods!
type Service struct {
	mediatype *mediatype.MediaType
	uri       string // Fill just if any http like type handler needed! Simple URI not variabale included! API services can set like "/m?{{.ServiceID}}" but it is not efficient, find services by ID.

	priority protocol.Priority // Use to queue requests by its priority
	weight   protocol.Weight   // Use to queue requests by its weights in the same priority

	// Authorization data to authorize incoming service
	id       uint64
	crud     protocol.CRUD // CRUD == Create, Read, Update, Delete
	userType protocol.UserType
}

// New returns a new error!
func New(uri string, mediatype *mediatype.MediaType) (s *Service) {
	if mediatype == nil || mediatype.ID() == 0 {
		// This condition must be true just in the dev phase.
		panic("Service must have a valid mediatype and ID. It is rule to add more detail about service before register it!")
	}

	s = &Service{
		mediatype: mediatype,
		uri:       uri,
		id:        mediatype.ID(),
	}
	return
}

func (s *Service) SetPriority(priority protocol.Priority, weight protocol.Weight) *Service {
	s.priority = priority
	s.weight = weight
	return s
}

func (s *Service) SetAuthorization(crud protocol.CRUD, userType protocol.UserType) *Service {
	s.crud = crud
	s.userType = userType
	return s
}

func (s *Service) MediaType() protocol.MediaType { return s.mediatype }
func (s *Service) URI() string                   { return s.uri }
func (s *Service) Priority() protocol.Priority   { return s.priority }
func (s *Service) Weight() protocol.Weight       { return s.weight }
func (s *Service) ID() uint64                    { return s.id }
func (s *Service) CRUDType() protocol.CRUD       { return s.crud }
func (s *Service) UserType() protocol.UserType   { return s.userType }

/*
*********** Handlers ***********
not-implemented handlers of the service.
*/

func (s *Service) ServeSRPC(st protocol.Stream) (err protocol.Error) {
	err = ErrServiceNotAcceptSRPC
	return
}
func (s *Service) ServeSRPCDirect(conn protocol.Connection, request []byte) (response []byte, err protocol.Error) {
	err = ErrServiceNotAcceptSRPCDirect
	return
}
func (s *Service) ServeHTTP(st protocol.Stream, httpReq protocol.HTTPRequest, httpRes protocol.HTTPResponse) (err protocol.Error) {
	err = ErrServiceNotAcceptHTTP
	return
}
