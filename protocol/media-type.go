/* For license and copyright information please see LEGAL file in repository */

package protocol

type MediaTypes interface {
	RegisterMediaType(mt MediaType)
	GetMediaTypeByID(urnID uint64) MediaType
	GetMediaTypeByURN(urnURI string) MediaType
	GetMediaTypeByFileExtension(ex string) MediaType
	GetMediaTypeByType(mt string) MediaType
}

// MediaType is standard shape of any coding media-type
// MediaType or MimeType standrad list can found here:
// http://www.iana.org/assignments/media-types/media-types.xhtml
// https://en.wikipedia.org/wiki/Media_type
// https://tools.ietf.org/html/rfc6838
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
type MediaType interface {
	URN() GitiURN
	MediaType() string
	MainType() string
	SubType() string
	FileExtension() string
	Status() MediaTypeStatus
	IssueDate() Time
	ExpiryDate() Time
	ExpireInFavorOf() MediaType
}

type MediaTypeStatus uint8

// MediaType Status
// https://en.wikipedia.org/wiki/Software_release_life_cycle
const (
	// MediaTypeStatePreAlpha refers to all activities performed during the software project before formal testing.
	MediaTypeStatePreAlpha MediaTypeStatus = iota
	// MediaTypeStateAlpha is the first phase to begin software testing
	MediaTypeStateAlpha
	// MediaTypeStateBeta generally begins when the software is feature complete but likely to contain a number of known or unknown bugs.
	MediaTypeStateBeta
	// MediaTypeStatePerpetualBeta where new features are continually added to the software without establishing a final "stable" release.
	// This technique may allow a developer to delay offering full support and responsibility for remaining issues.
	MediaTypeStatePerpetualBeta
	// MediaTypeStateOpenBeta is for a larger group, or anyone interested.
	MediaTypeStateOpenBeta
	// MediaTypeStateClosedBeta is released to a restricted group of individuals for a user test by invitation.
	MediaTypeStateClosedBeta
	// MediaTypeStateReleaseCandidate also known as "going silver", is a beta version with potential to be a stable product,
	// which is ready to release unless significant bugs emerge
	MediaTypeStateReleaseCandidate
	// MediaTypeStateStableRelease Also called production release,
	// the stable release is the last release candidate (RC) which has passed all verifications / tests.
	MediaTypeStateStableRelease
	// MediaTypeStateEndOfLife indicate no longer supported and continue its existence until to ExpiryDate!
	MediaTypeStateEndOfLife
)
