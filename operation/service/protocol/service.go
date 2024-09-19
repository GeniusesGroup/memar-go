/* For license and copyright information please see the LEGAL file in the code repository */

package service_p

import (
	object_p "memar/computer/language/object/protocol"
	datatype_p "memar/datatype/protocol"
	mediatype_p "memar/mediatype/protocol"
	operation_p "memar/operation/protocol"
	user_p "memar/user/protocol"
)

// Service is the interface that must implement by any struct to be a service.
type Service interface {
	Field_ServiceID

	Authorization
	Details

	operation_p.Importance
	object_p.LifeCycle

	datatype_p.DataType
	mediatype_p.MediaType
}

// Service authorization to authorize incoming service request
type Authorization interface {
	operation_p.Field_ActionType
	user_p.Field_UserType
}

type Details /*[REQ, RES DataType]*/ interface {
	Request() datatype_p.DataType
	Response() datatype_p.DataType
}
