/* For license and copyright information please see LEGAL file in repository */

package protocol

// Services is the interface that must implement by any Application!
type Services interface {
	RegisterService(s Service)
	GetServiceByID(urnID uint64) (ser Service, err Error)
	GetServiceByURN(urn string) (ser Service, err Error)
	GetServiceByURI(domain, uri string) (ser Service, err Error) // register in URN.Domain group, due to multi domain servises need to register
}

// Service is the interface that must implement by any struct to be a service!
// Set fields methods in this type must accept just once to prevent any mistake by change after set first!
type Service interface {
	Detail(lang LanguageID) ServiceDetail
	URN() GitiURN
	URI() string
	Status() SoftwareStatus
	IssueDate() Time
	ExpiryDate() Time
	ExpireInFavorOf() GitiURN

	// Service Authorization
	CRUDType() CRUD
	UserType() UserType

	// DirectHandler use to call a service without need to open any stream.
	// It can also use when service request data is smaller than network MTU.
	// Or use for time sensitive data like audio and video that streams shape in app layer
	DirectHandler(conn Connection, request []byte) (response []byte, err Error)
	SRPCHandler
	HTTPHandler
	CLIHandler
	// Can't standardize here, must implement as pure private(package scope) function outside service scope.
	// Handler(Stream, interface{}) (interface{}, Error)

	// JSON
}

// ServiceDetail return locale detail about the service!
type ServiceDetail interface {
	Name() string
	Description() string
	TAGS() []string
}
