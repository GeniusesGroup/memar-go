/* For license and copyright information please see the LEGAL file in the code repository */

package service_p

import (
	object_p "memar/computer/language/object/protocol"
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
	mediatype_p "memar/mediatype/protocol"
	operation_p "memar/operation/protocol"
	user_p "memar/user/protocol"
)

type ServiceID = datatype_p.ID

// Service is the interface that must implement by any struct to be a service.
type Service interface {
	// Usually easily return s.DataTypeID() or it can return some old way numbering like HTTP:80, HTTPS:443, ...
	ServiceID() ServiceID

	Service_Authorization
	operation_p.Importance
	Service_Details
	object_p.LifeCycle

	datatype_p.DataType
	mediatype_p.MediaType
}

// ServiceHandlers is just test (approver) interface and MUST NOT use directly in any signature.
// Due to Golang import cycle problem we can't use `net_p.Socket`
type ServiceHandlers[SK any /*net_p.Socket*/, ReqT, ResT datatype_p.DataType] interface {
	// Call service locally by import service package to other one
	Process(sk SK, req ReqT) (res ResT, err error_p.Error)
	//
	// Call service remotely by preferred(SDK generator choose) protocol.
	Do(req ReqT) (res ResT, err error_p.Error)
}

// Service authorization to authorize incoming service request
type Service_Authorization interface {
	operation_p.Field_CRUD
	UserType() user_p.Type
}

type Service_Details /*[REQ, RES DataType]*/ interface {
	Request() datatype_p.DataType
	Response() datatype_p.DataType
}
