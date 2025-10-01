/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
	operation_p "memar/process/operation/protocol"
	service_p "memar/process/service/protocol"
	request_p "memar/process/request/protocol"
	response_p "memar/process/response/protocol"
)

// Service implement protocol.Service when embed to other struct that implements other needed methods.
type Service struct {
	datatype.DataType
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (s *Service) Deinit() (err error_p.Error) { return }

//memar:impl memar/identifier/mediatype/protocol.MediaType
func (s *Service) MediaType() string { return "" }

//memar:impl memar/process/service/protocol.Field_ServiceID
func (s *Service) ServiceID() service_p.ID { return 0 }

//memar:impl memar/process/service/protocol.Authorization
func (s *Service) ActionType() operation_p.ActionType { return operation_p.ActionType_None }

//memar:impl memar/process/operation/protocol.Importance
func (s *Service) Priority() operation_p.Priority { return operation_p.Priority_Unset }
func (s *Service) Weight() operation_p.Weight     { return operation_p.Weight_Unset }

//memar:impl memar/process/service/protocol.Details
func (s *Service) Request() request_p.Request  { return nil }
func (s *Service) Response() response_p.Response { return nil }
