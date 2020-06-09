/* For license and copyright information please see LEGAL file in repository */

package http

import "errors"

// Declare Errors Details
var (
	ErrHTTPPacketTooShort = errors.New("Received HTTP packet size is smaller than expected and can't use")
	ErrHTTPPacketTooLong  = errors.New("Received HTTP packet size is larger than expected and can't use")

	ErrParsedErrorOnMethod                    = errors.New("Parsing received HTTP packet encounter unknown situation in Method part")
	ErrParsedErrorOnURI                       = errors.New("Parsing received HTTP packet encounter unknown situation in URI part")
	ErrParsedErrorOnVersion                   = errors.New("Parsing received HTTP packet encounter unknown situation in Version part")
	ErrParsedErrorOnStatusCode                = errors.New("Parsing received HTTP packet encounter unknown situation in StatusCode part")
	ErrParsedErrorOnReasonPhrase              = errors.New("Parsing received HTTP packet encounter unknown situation in ReasonPhrase part")
	ErrParsedErrorRequestHeaderFieldsTooLarge = errors.New("Parsing received HTTP packet encounter situation that header part of http packet is larger than expected")

	ErrCookieBadName   = errors.New("Cookie name include illegal charecter by related RFC")
	ErrCookieBadValue  = errors.New("Cookie value include illegal charecter by related RFC")
	ErrCookieBadPath   = errors.New("Cookie path include illegal charecter by related RFC")
	ErrCookieBadDomain = errors.New("Cookie domain is not valid by related RFC")
)
