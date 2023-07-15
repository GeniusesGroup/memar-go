/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Services is the interface that must implement by any Application.
type Services interface {
	// RegisterService use to register application services.
	// Due to minimize performance impact, This method isn't safe to use concurrently and
	// must register all service before use GetService methods.
	RegisterService(s Service) (err Error)

	Services() []Service
	GetServiceByID(id MediaTypeID) (ser Service, err Error)
	GetServiceByMediaType(mt string) (ser Service, err Error)
	GetServiceByURI(uri string) (ser Service, err Error)
}

// Service is the interface that must implement by any struct to be a service.
type Service interface {
	// Fill just if any http like type handler needed. Suggest use simple immutable path,
	// not variable data included in path describe here:https://www.rfc-editor.org/rfc/rfc6570 e.g. "/product?id=1" instead of "/product/1/"
	// API services can set like "/s?{{.ServiceID}}" but it is not efficient, instead find services by ID as base64
	URI() string // suggest use just URI.Path

	// Service authorization to authorize incoming service request
	CRUDType() CRUD
	UserType() UserType

	SRPCHandler

	OperationImportance
	ServiceDetails
	ObjectLifeCycle
}

type ServiceHandlers[ReqT, ResT any] interface {
	// Call service locally by import service package to other one
	Process(sk Socket, req ReqT) (res ResT, err Error)
	//
	// Call service remotely by preferred protocol.
	Do(req ReqT) (res ResT, err Error)
}

type ServiceDetails interface {
	Request() DataType
	Response() DataType

	Details
	MediaType
}
