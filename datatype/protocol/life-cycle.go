/* For license and copyright information please see the LEGAL file in the code repository */

package datatype_p

type Field_LifeCycle interface {
	LifeCycle() LifeCycle
}

type LifeCycle uint8

// Software Status
// https://en.wikipedia.org/wiki/LifeCycle_release_life_cycle
const (
	LifeCycle_Unset LifeCycle = iota

	// LifeCycle_PreAlpha refers to all activities performed during the software project before formal testing.
	LifeCycle_PreAlpha
	// LifeCycle_Alpha is the first phase to begin software testing
	LifeCycle_Alpha
	// LifeCycle_Beta generally begins when the software is feature complete but likely to contain a number of known or unknown bugs.
	LifeCycle_Beta
	// LifeCycle_PerpetualBeta where new features are continually added to the software without establishing a final "stable" release.
	// This technique may allow a developer to delay offering full support and responsibility for remaining issues.
	LifeCycle_PerpetualBeta
	// LifeCycle_OpenBeta is for a larger group, or anyone interested.
	LifeCycle_OpenBeta
	// LifeCycle_ClosedBeta is released to a restricted group of individuals for a user test by invitation.
	LifeCycle_ClosedBeta
	// LifeCycle_ReleaseCandidate also known as "going silver", is a beta version with potential to be a stable product,
	// which is ready to release unless significant bugs emerge
	LifeCycle_ReleaseCandidate
	// LifeCycle_StableRelease Also called production release,
	// the stable release is the last release candidate (RC) which has passed all verifications / tests.
	LifeCycle_StableRelease
	// LifeCycle_EndOfLife indicate no longer supported and continue its existence until to ExpiryDate!
	LifeCycle_EndOfLife
)
