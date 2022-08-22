/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Services is the interface that must implement by any Application.
type Services interface {
	// RegisterService use to register application services.
	// Due to minimize performance impact, This method isn't safe to use concurrently and
	// must register all service before use GetService methods.
	RegisterService(s Service)

	GetServiceByID(id MediaTypeID) (ser Service, err Error)
	GetServiceByMediaType(mt string) (ser Service, err Error)
	GetServiceByURI(uri string) (ser Service, err Error)
}

// Service is the interface that must implement by any struct to be a service.
// Set fields methods in this type must accept just once to prevent any mistake by change after set first.
type Service interface {
	// Fill just if any http like type handler needed! Simple URI not variable included.
	// API services can set like "/m?{{.ServiceID}}" but it is not efficient, instead find services by ID as integer.
	URI() string // suggest use just URI.Path

	Priority() Priority // Use to queue requests by its priority
	Weight() Weight     // Use to queue requests by its weights in the same priority

	// Service authorization to authorize incoming service request
	CRUDType() CRUD
	UserType() UserType

	// Handlers, Due to specific args and returns, we can't uncomment some of them
	//
	// Call service locally by import service package to other one
	// Process(st Stream, req interface{}) (res interface{}, err Error)
	//
	// Call service remotely by preferred protocol.
	// Do(req interface{}) (res interface{}, err Error)

	SRPCHandler
	HTTPHandler // Some other protocol like gRPC, SOAP, ... must implement inside HTTP, If they are use HTTP as a transfer protocol.

	ServiceDetails
}

type ServiceDetails interface {
	Request() []Field
	Response() []Field

	Details
	MediaType
}
