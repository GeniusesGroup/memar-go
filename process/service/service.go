/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"memar/datatype"
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
	operation_p "memar/operation/protocol"
	service_p "memar/operation/service/protocol"
)

// Service implement protocol.Service when embed to other struct that implements other needed methods.
type Service struct {
	datatype.DataType
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (s *Service) Deinit() (err error_p.Error) { return }

//memar:impl memar/mediatype/protocol.MediaType
func (s *Service) MediaType() string { return "" }

//memar:impl memar/operation/service/protocol.Field_ServiceID
func (s *Service) ServiceID() service_p.ID { return 0 }

//memar:impl memar/operation/service/protocol.Authorization
func (s *Service) ActionType() operation_p.ActionType { return operation_p.ActionType_None }

//memar:impl memar/operation/protocol.Importance
func (s *Service) Priority() operation_p.Priority { return operation_p.Priority_Unset }
func (s *Service) Weight() operation_p.Weight     { return operation_p.Weight_Unset }

//memar:impl memar/operation/service/protocol.Details
func (s *Service) Request() datatype_p.DataType  { return nil }
func (s *Service) Response() datatype_p.DataType { return nil }
