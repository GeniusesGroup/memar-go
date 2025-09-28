/* For license and copyright information please see the LEGAL file in the code repository */

package service_p

import (
	capsule_p "memar/computer/capsule/protocol"
	datatype_p "memar/datatype/protocol"
	mediatype_p "memar/mediatype/protocol"
	operation_p "memar/operation/protocol"
)

type Field_Service interface {
	Service() Service
}

// Service is the interface that must implement by any struct to be a service.
type Service interface {
	capsule_p.LifeCycle

	operation_p.Field_OperationID
	operation_p.Field_ActionType
	operation_p.Field_Request
	operation_p.Field_Response
	
	operation_p.Importance
	operation_p.Authorize

	datatype_p.DataType
	mediatype_p.MediaType
}
