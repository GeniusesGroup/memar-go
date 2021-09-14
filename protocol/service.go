/* For license and copyright information please see LEGAL file in repository */

package protocol

// Services is the interface that must implement by any Application!
type Services interface {
	RegisterService(s Service)
	UnRegisterService(s Service)
	GetServiceByID(urnID uint64) Service
	GetServiceByURN(urn string) Service
	GetServiceByURI(uri string) Service
}

// Service is the interface that must implement by any struct to be a service!
// Set fields methods in this type must accept just once to prevent any mistake by change after set first!
type Service interface {
	Detail(lang LanguageID) ServiceDetail
	URN() GitiURN
	URI() string
	Status() ServiceStatus
	IssueDate() Time
	ExpiryDate() Time

	// Service Authorization
	CRUDType() CRUD
	UserType() UserType

	SRPCHandler
	HTTPHandler
	CLIHandler
	// Service method is the handler of the service but we suggest to implement it as pure private(package scope) function outside service scope.
	// Service(Stream, interface{}) (interface{}, Error)

	JSON
}

// ServiceDetail return locale detail about the service!
type ServiceDetail interface {
	Name() string
	Description() string
	TAGS() []string
}

type ServiceStatus uint8

// Service Status
// https://en.wikipedia.org/wiki/Software_release_life_cycle
const (
	// ServiceStatePreAlpha refers to all activities performed during the software project before formal testing.
	ServiceStatePreAlpha ServiceStatus = iota
	// ServiceStateAlpha is the first phase to begin software testing
	ServiceStateAlpha
	// ServiceStateBeta generally begins when the software is feature complete but likely to contain a number of known or unknown bugs.
	ServiceStateBeta
	// ServiceStatePerpetualBeta where new features are continually added to the software without establishing a final "stable" release.
	// This technique may allow a developer to delay offering full support and responsibility for remaining issues.
	ServiceStatePerpetualBeta
	// ServiceStateOpenBeta is for a larger group, or anyone interested.
	ServiceStateOpenBeta
	// ServiceStateClosedBeta is released to a restricted group of individuals for a user test by invitation.
	ServiceStateClosedBeta
	// ServiceStateReleaseCandidate also known as "going silver", is a beta version with potential to be a stable product,
	// which is ready to release unless significant bugs emerge
	ServiceStateReleaseCandidate
	// ServiceStateStableRelease Also called production release,
	// the stable release is the last release candidate (RC) which has passed all verifications / tests.
	ServiceStateStableRelease
	// ServiceStateEndOfLife indicate no longer supported and continue its existence until to ExpiryDate!
	ServiceStateEndOfLife
)
