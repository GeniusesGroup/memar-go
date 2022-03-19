/* For license and copyright information please see LEGAL file in repository */

package protocol

// Services is the interface that must implement by any Application!
type Services interface {
	RegisterService(s Service)
	GetServiceByID(mtID uint64) (ser Service, err Error)
	GetServiceByMediaType(mt string) (ser Service, err Error)
	GetServiceByURI(uri string) (ser Service, err Error)
}

// Service is the interface that must implement by any struct to be a service!
// Set fields methods in this type must accept just once to prevent any mistake by change after set first!
type Service interface {
	MediaType() MediaType
	// Request() MediaType
	// Response() MediaType
	URI() string // HTTPURI.Path

	Priority() Priority // Use to queue requests by its priority
	Weight() Weight     // Use to queue requests by its weights in the same priority

	// Service Authorization
	ID() uint64 // copy of MediaType().ID() to improve authorization mechanism performance
	CRUDType() CRUD
	UserType() UserType

	// Handlers
	SRPCHandler
	HTTPHandler // Some other protocol like gRPC, SOAP, ... must implement inside HTTP, If they are use HTTP as a transfer protocol.
	// Due to specific args and returns, we can't standardize here.
	// Do(st Stream, req interface{}) (res interface{}, err Error)	Call service locally by import service package to other one
	// DoSRPC(req interface{}) (res interface{}, err Error)			Call service remotely by sRPC protocol
	// DoHTTP(req interface{}) (res interface{}, err Error)			Call service remotely by HTTP protocol
}
