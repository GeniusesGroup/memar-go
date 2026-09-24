/* For license and copyright information please see the LEGAL file in the code repository */

package service

import (
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
	operation_p "memar/process/operation/protocol"
	request_p "memar/process/request/protocol"
	response_p "memar/process/response/protocol"
)

// Service implement protocol.Service when embed to other struct that implements other needed methods.
type Service struct {
	datatype.DataType
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Service) Deinit() (err error_p.Error) { return }

//memar:impl memar/identifier/mediatype/protocol.MediaType
func (self *Service) MediaType() string { return "ERROR::: Service MUST indicate MediaType for itself" }

//memar:impl memar/process/operation/protocol.Field_OperationID
func (self *Service) OperationID() operation_p.ID { return 0 }

//memar:impl memar/process/service/protocol.Authorization
func (self *Service) ActionType() operation_p.ActionType { return operation_p.ActionType_None }

//memar:impl memar/process/operation/protocol.Importance
func (self *Service) OperationPriority() operation_p.Priority { return operation_p.Priority_Unset }
func (self *Service) OperationWeight() operation_p.Weight     { return operation_p.Weight_Unset }

//memar:impl memar/process/service/protocol.Details
func (self *Service) Request() request_p.Request    { return nil }
func (self *Service) Response() response_p.Response { return nil }

func (self *Service) authorize() error_p.Error { return nil }
