/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	operation_p "memar/operation/protocol"
	service_p "memar/operation/service/protocol"
	"memar/protocol"
	user_p "memar/user/protocol"
)

// Service implement protocol.Service when embed to other struct that implements other needed methods.
type Service struct{}

//memar:impl memar/protocol.Service
func (s *Service) ID() service_p.ServiceID { return 0 }

//memar:impl memar/protocol.Service_Authorization
func (s *Service) CRUDType() operation_p.CRUD { return operation_p.CRUD_None }
func (s *Service) UserType() user_p.Type      { return user_p.Type_Unset }

//memar:impl memar/operation/protocol.Importance
func (s *Service) Priority() operation_p.Priority { return operation_p.Priority_Unset }
func (s *Service) Weight() operation_p.Weight     { return operation_p.Weight_Unset }

//memar:impl memar/protocol.ServiceDetails
func (s *Service) Request() protocol.DataType  { return nil }
func (s *Service) Response() protocol.DataType { return nil }
