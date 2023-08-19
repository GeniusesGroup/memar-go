/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"memar/protocol"
)

// Service implement protocol.Service when embed to other struct that implements other needed methods.
type Service struct{}

//memar:impl memar/protocol.Service
func (s *Service) ID() protocol.ServiceID { return 0 }

//memar:impl memar/protocol.Service_Authorization
func (s *Service) CRUDType() protocol.CRUD     { return protocol.CRUD_None }
func (s *Service) UserType() protocol.UserType { return protocol.UserType_Unset }

//memar:impl memar/protocol.OperationImportance
func (s *Service) Priority() protocol.Priority { return protocol.Priority_Unset }
func (s *Service) Weight() protocol.Weight     { return protocol.Weight_Unset }

//memar:impl memar/protocol.ServiceDetails
func (s *Service) Request() protocol.DataType  { return nil }
func (s *Service) Response() protocol.DataType { return nil }
