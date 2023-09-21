//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "HTTP"

func init() {
	ErrNoConnection.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"No Connection",
		"There is no connection to peer(server or client) to process request",
		"",
		"",
		nil)

	ErrPacketTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"Received HTTP packet size is shorter than expected and can't use",
		"",
		"",
		nil)

	ErrPacketTooLong.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Long",
		"Received HTTP packet size is larger than expected and can't use",
		"",
		"",
		nil)
	
	ErrParseMethod.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Method",
		"Parsing received HTTP packet encounter unknown situation in Method part",
		"",
		"",
		nil)
	ErrParseVersion.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Version",
		"Parsing received HTTP packet encounter unknown situation in Version part",
		"",
		"",
		nil)
	ErrParseStatusCode.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Status Code",
		"Parsing received HTTP packet encounter unknown situation in StatusCode part",
		"",
		"",
		nil)
	ErrParseReasonPhrase.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Reason Phrase",
		"Parsing received HTTP packet encounter unknown situation in ReasonPhrase part",
		"",
		"",
		nil)

		// Their HTTP client may or may not be
			// able to read this if we're
			// responding to them and hanging up
			// while they're still writing their
			// request. Undefined behavior.
	ErrParseHeaderTooLarge.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Header Too Large",
		"Parsing received HTTP packet encounter situation that header part of http packet is larger than expected",
		"",
		"",
		nil)

	ErrCookieBadName.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Name",
		"Cookie name include illegal charecter by related RFC",
		"",
		"",
		nil)
	ErrCookieBadValue.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Value",
		"Cookie value include illegal charecter by related RFC",
		"",
		"",
		nil)
	ErrCookieBadPath.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Path",
		"Cookie path include illegal charecter by related RFC",
		"",
		"",
		nil)
	ErrCookieBadDomain.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Domain",
		"Cookie domain is not valid by related RFC",
		"",
		"",
		nil)

	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Requested HTTP URI Service is not found in this instance of app",
		"",
		"",
		nil)

		
	ErrUnsupportedMediaType.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Unsupported Media Type",
		"Refuse to accept the request or response because the payload format or encoding is in an unsupported format",
		"",
		"",
		nil)

			// Respond as per RFC 7230 Section 3.3.1 which says,
			//      A server that receives a request message with a
			//      transfer coding it does not understand SHOULD
			//      respond with 501 (Unimplemented).

	ErrServiceNotAcceptHTTP.SetDetail(detail.New(protocol.LanguageEnglish, domainEnglish).
		SetName("").
		SetAbbreviation("").
		SetAliases([]string{}).
		SetSummary("Service Not Accept HTTP").
		SetOverview("Requested service by given ID not accept HTTP protocol in this server").
		SetUserNote("Try other server or contact support of the software").
		SetDevNote("It is so easy to implement HTTP handler for a service! Take a time and do it!").
		SetTAGS([]string{})
	)
	
	ErrBodySizeMismatch.SetDetail(protocol.LanguageEnglish, domainEnglish,
			"Body Size Mismatch",
			"This error will occur when the body's length did not match the length from the Content-Length header.",
			"This typically occurs when the data is malformed, or when the Content-Length header was calculated based on characters instead of bytes.",
			"",
			nil)
}
