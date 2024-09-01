/* For license and copyright information please see the LEGAL file in the code repository */

package service_p

import (
	operation_p "memar/operation/protocol"
	"memar/protocol"
	user_p "memar/user/protocol"
)

type ServiceID = protocol.DataTypeID

// Service is the interface that must implement by any struct to be a service.
type Service interface {
	// Usually easily return s.DataTypeID() or it can return some old way numbering like HTTP:80, HTTPS:443, ...
	ServiceID() ServiceID

	Service_Authorization
	operation_p.Importance
	Service_Details
	protocol.ObjectLifeCycle

	protocol.DataType
	protocol.MediaType
}

// ServiceHandlers is just test (approver) interface and MUST NOT use directly in any signature.
// Due to Golang import cycle problem we can't use `net_p.Socket`
type ServiceHandlers[SK any /*net_p.Socket*/, ReqT, ResT protocol.DataType] interface {
	// Call service locally by import service package to other one
	Process(sk SK, req ReqT) (res ResT, err protocol.Error)
	//
	// Call service remotely by preferred(SDK generator choose) protocol.
	Do(req ReqT) (res ResT, err protocol.Error)
}

// Service authorization to authorize incoming service request
type Service_Authorization interface {
	operation_p.Field_CRUD
	UserType() user_p.Type
}

type Service_Details /*[REQ, RES DataType]*/ interface {
	Request() protocol.DataType
	Response() protocol.DataType
}
