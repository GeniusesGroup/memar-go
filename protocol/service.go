/* For license and copyright information please see LEGAL file in repository */

package protocol

// Services is the interface that must implement by any Application!
type Services interface {
	RegisterService(s Service)
	GetServiceByID(urnID uint64) (ser Service, err Error)
	GetServiceByURN(urn string) (ser Service, err Error)
	GetServiceByURI(uri string) (ser Service, err Error)
}

// Service is the interface that must implement by any struct to be a service!
// Set fields methods in this type must accept just once to prevent any mistake by change after set first!
type Service interface {
	Detail(lang LanguageID) ServiceDetail
	URN() GitiURN
	URI() string // HTTPURI.Path
	Status() SoftwareStatus
	IssueDate() TimeUnixSec  // TODO::: Temporary use TimeUnixSec instead of Time
	ExpiryDate() TimeUnixSec // TODO::: Temporary use TimeUnixSec instead of Time
	ExpireInFavorOf() GitiURN
	Weight() Weight // Use to queue requests by services weights

	// Service Authorization
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

// ServiceDetail return locale detail about the service.
type ServiceDetail interface {
	Language() LanguageID
	// Domain return locale domain name that service belongs to it.
	Domain() string
	// Summary return locale general summary service text that gives the main points in a concise form.
	Summary() string
	// Overview return locale general service text that gives the main ideas without explaining all the details.
	Overview() string
	// Description return locale service text that gives the main ideas with explaining all the details and purposes.
	Description() string
	// TAGS  return locale service tags to sort service in groups for any purpose e.g. in GUI to help org manager to give service delegate authorization to staffs.
	TAGS() []string
}
