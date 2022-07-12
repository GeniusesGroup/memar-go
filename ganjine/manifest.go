/* For license and copyright information please see LEGAL file in repository */

package ganjine

// Manifest store Ganjine manifest data
type Manifest struct {
	DataCentersClass uint8 // 0:FirstClass 256:Low-Quality default:0

	TotalZones       uint8  // ReplicationNumber, deafult:3
	TotalNodesInZone uint32 // not count replicated nodes, just one of them count.

	TransactionTimeOut uint16 // in ms, default:500ms, Max 65.535s timeout
	NodeFailureTimeOut uint16 // in minute, default:60m, other corresponding node same range will replace failed node! not use in network failure, it is handy proccess!

	CachePercent uint8 // GC cached records when reach this size limit!
}

// DataStructure State
// https://en.wikipedia.org/wiki/Software_release_life_cycle
const (
	// DataStructureStatePreAlpha refers to all activities performed during the software project before formal testing.
	DataStructureStatePreAlpha uint8 = iota
	// DataStructureStateAlpha is the first phase to begin software testing
	DataStructureStateAlpha
	// DataStructureStateBeta generally begins when the software is feature complete but likely to contain a number of known or unknown bugs.
	DataStructureStateBeta
	// DataStructureStatePerpetualBeta where new features are continually added to the software without establishing a final "stable" release.
	// This technique may allow a developer to delay offering full support and responsibility for remaining issues.
	DataStructureStatePerpetualBeta
	// DataStructureStateOpenBeta is for a larger group, or anyone interested.
	DataStructureStateOpenBeta
	// DataStructureStateClosedBeta is released to a restricted group of individuals for a user test by invitation.
	DataStructureStateClosedBeta
	// DataStructureStateReleaseCandidate also known as "going silver", is a beta version with potential to be a stable product,
	// which is ready to release unless significant bugs emerge
	DataStructureStateReleaseCandidate
	// DataStructureStateStableRelease Also called production release,
	// the stable release is the last release candidate (RC) which has passed all verifications / tests.
	DataStructureStateStableRelease
	// DataStructureStateEndOfLife indicate no longer supported and continue its existence until to ExpiryDate!
	DataStructureStateEndOfLife
)
