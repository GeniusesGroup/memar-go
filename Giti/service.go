/* For license and copyright information please see LEGAL file in repository */

package giti

// Service is the interface that must implement by any struct to be a service!
// Set fields methods in this type must accept just once to prevent any mistake by change after set first!
type Service interface {
	URN() string
	ID() uint32
	Domain() string
	URI() string

	Status() ServiceStatus
	SetStatus(status ServiceStatus) // Just once

	// SRPCSyllabHandler method is sRPC handler of the service with Syllab codec data in the stream payload.
	SRPCSyllabHandler(st Stream) (err Error)
	// HTTPHandler method is HTP handler of the service.
	HTTPHandler(st Stream, httpReq HTTPRequest, httpRes HTTPResponse) (err Error)
	// CLIHandler method is the command-line interface (CLI) handler of the service
	CLIHandler() (err Error)

	// Handler method is main handler of the service but must implement as pure function outside service scope.
	// handler(Stream, ServiceRequest) (ServiceResponse, Error)

	ServiceRequest() ServiceRequest
	ServiceResponse() ServiceResponse

	JSON
}

// ServiceRequest is the interface that must implement by any struct to be a service request!
type ServiceRequest interface{}

// ServiceResponse is the interface that must implement by any struct to be a service request!
type ServiceResponse interface{}

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
