/* For license and copyright information please see LEGAL file in repository */

package protocol

type SoftwareStatus uint8

// Software Status
// https://en.wikipedia.org/wiki/Software_release_life_cycle
const (
	// SoftwareStatePreAlpha refers to all activities performed during the software project before formal testing.
	SoftwareStatePreAlpha SoftwareStatus = iota
	// SoftwareStateAlpha is the first phase to begin software testing
	SoftwareStateAlpha
	// SoftwareStateBeta generally begins when the software is feature complete but likely to contain a number of known or unknown bugs.
	SoftwareStateBeta
	// SoftwareStatePerpetualBeta where new features are continually added to the software without establishing a final "stable" release.
	// This technique may allow a developer to delay offering full support and responsibility for remaining issues.
	SoftwareStatePerpetualBeta
	// SoftwareStateOpenBeta is for a larger group, or anyone interested.
	SoftwareStateOpenBeta
	// SoftwareStateClosedBeta is released to a restricted group of individuals for a user test by invitation.
	SoftwareStateClosedBeta
	// SoftwareStateReleaseCandidate also known as "going silver", is a beta version with potential to be a stable product,
	// which is ready to release unless significant bugs emerge
	SoftwareStateReleaseCandidate
	// SoftwareStateStableRelease Also called production release,
	// the stable release is the last release candidate (RC) which has passed all verifications / tests.
	SoftwareStateStableRelease
	// SoftwareStateEndOfLife indicate no longer supported and continue its existence until to ExpiryDate!
	SoftwareStateEndOfLife
)
