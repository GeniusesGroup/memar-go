/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type SoftwareStatus uint8

// Software Status
// https://en.wikipedia.org/wiki/Software_release_life_cycle
const (
	Software_Unset SoftwareStatus = iota

	// Software_PreAlpha refers to all activities performed during the software project before formal testing.
	Software_PreAlpha
	// Software_Alpha is the first phase to begin software testing
	Software_Alpha
	// Software_Beta generally begins when the software is feature complete but likely to contain a number of known or unknown bugs.
	Software_Beta
	// Software_PerpetualBeta where new features are continually added to the software without establishing a final "stable" release.
	// This technique may allow a developer to delay offering full support and responsibility for remaining issues.
	Software_PerpetualBeta
	// Software_OpenBeta is for a larger group, or anyone interested.
	Software_OpenBeta
	// Software_ClosedBeta is released to a restricted group of individuals for a user test by invitation.
	Software_ClosedBeta
	// Software_ReleaseCandidate also known as "going silver", is a beta version with potential to be a stable product,
	// which is ready to release unless significant bugs emerge
	Software_ReleaseCandidate
	// Software_StableRelease Also called production release,
	// the stable release is the last release candidate (RC) which has passed all verifications / tests.
	Software_StableRelease
	// Software_EndOfLife indicate no longer supported and continue its existence until to ExpiryDate!
	Software_EndOfLife
)
