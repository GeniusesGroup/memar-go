/* For license and copyright information please see LEGAL file in repository */

package protocol

// Services is the interface that must implement by any Application!
type Services interface {
	// RegisterService use to register application services.
	// Due to minimize performance impact, This method isn't safe to use concurrently and
	// must register all service before use GetService methods.
	RegisterService(s Service)

	GetServiceByID(mtID uint64) (ser Service, err Error)
	GetServiceByMediaType(mt string) (ser Service, err Error)
	GetServiceByURI(uri string) (ser Service, err Error)
}

// Service is the interface that must implement by any struct to be a service!
// Set fields methods in this type must accept just once to prevent any mistake by change after set first!
type Service interface {
	// Request() MediaType
	// Response() MediaType
	URI() string // HTTPURI.Path

	Priority() Priority // Use to queue requests by its priority
	Weight() Weight     // Use to queue requests by its weights in the same priority

	// Service Authorization
	CRUDType() CRUD
	UserType() UserType

	// Handlers, Due to specific args and returns, we can't uncomment some of them
	// Handle(st Stream, req interface{}) (res interface{}, err Error)	Call service locally by import service package to other one
	SRPCHandler
	HTTPHandler // Some other protocol like gRPC, SOAP, ... must implement inside HTTP, If they are use HTTP as a transfer protocol.
	// Do(req interface{}) (res interface{}, err Error)			Call service remotely by preferred protocol.
	// DoSRPC(req interface{}) (res interface{}, err Error)		Call service remotely by sRPC protocol
	// DoHTTP(req interface{}) (res interface{}, err Error)		Call service remotely by HTTP protocol

	MediaType
}
