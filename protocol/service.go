/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type ServiceID = MediaTypeID

// Service is the interface that must implement by any struct to be a service.
type Service interface {
	ID() ServiceID // Usually easily return s.MT.ID() or it can return some old way numbering like HTTP:80, HTTPS:443, ...

	Service_Authorization
	OperationImportance
	ServiceDetails
	ObjectLifeCycle
}

type ServiceHandlers[ReqT, ResT any] interface {
	// Call service locally by import service package to other one
	Process(sk Socket, req ReqT) (res ResT, err Error)
	//
	// Call service remotely by preferred(SDK generator choose) protocol.
	Do(req ReqT) (res ResT, err Error)
}

// Service authorization to authorize incoming service request
type Service_Authorization interface {
	CRUDType() CRUD
	UserType() UserType
}

type ServiceDetails interface {
	Request() DataType
	Response() DataType

	Detail
	MediaType
}
