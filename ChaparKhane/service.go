/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// Service store needed data for service
type Service struct {
	ID                      uint32 // Handy ID or Hash of name!
	URI                     string // Fill just if any http like type handler needed! Simple URI not variabale included!
	Name                    string
	IssueDate               int64
	ExpiryDate              int64
	ExpireInFavorOf         string // Other service name
	ExpireInFavorOfID       uint32 // Other ServiceID! Handy ID or Hash of ExpireInFavorOf!
	Status                  uint8
	Handler                 StreamHandler
	MinExpectedRequestSize  uint64 // to improve performance by alloc stream buffer size
	MaxExpectedRequestSize  uint64 // to improve performance by alloc stream buffer size
	MinExpectedResponseSize uint64 // to improve performance by alloc stream buffer size
	MaxExpectedResponseSize uint64 // to improve performance by alloc stream buffer size
	Description             []string
	TAGS                    []string
}

// Service Status
// https://en.wikipedia.org/wiki/Software_release_life_cycle
const (
	// ServiceStatePreAlpha refers to all activities performed during the software project before formal testing.
	ServiceStatePreAlpha uint8 = iota
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
