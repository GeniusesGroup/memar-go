/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	er "github.com/GeniusesGroup/libgo/error"
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainEnglish = "HTTP"
const domainPersian = "HTTP"

// Declare package errors
var (
	ErrNoConnection         er.Error
	ErrPacketTooShort       er.Error
	ErrPacketTooLong        er.Error
	ErrParseMethod          er.Error
	ErrParseVersion         er.Error
	ErrParseStatusCode      er.Error
	ErrParseReasonPhrase    er.Error
	ErrParseHeaderTooLarge  er.Error
	ErrCookieBadName        er.Error
	ErrCookieBadValue       er.Error
	ErrCookieBadPath        er.Error
	ErrCookieBadDomain      er.Error
	ErrNotFound             er.Error
	ErrUnsupportedMediaType er.Error
)

func init() {
	ErrNoConnection.Init("domain/http.protocol; type=error; name=no-connection")
	ErrNoConnection.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"No Connection",
		"There is no connection to peer(server or client) to process request",
		"",
		"",
		nil)
	ErrNoConnection.SetDetail(protocol.LanguagePersian, domainPersian,
		"عدم وجود ارتباط",
		"هیچ راه ارتباطی با رایانه مقصد برای پردازش درخواست مورد نظر وجود ندارد",
		"",
		"",
		nil)

	ErrPacketTooShort.Init("domain/http.protocol; type=error; name=packet-too-short")
	ErrPacketTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"Received HTTP packet size is shorter than expected and can't use",
		"",
		"",
		nil)

	ErrPacketTooLong.Init("domain/http.protocol; type=error; name=packet-too-long")
	ErrPacketTooLong.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Long",
		"Received HTTP packet size is larger than expected and can't use",
		"",
		"",
		nil)

	ErrParseMethod.Init("domain/http.protocol; type=error; name=parse-method")
	ErrParseMethod.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Method",
		"Parsing received HTTP packet encounter unknown situation in Method part",
		"",
		"",
		nil)

	ErrParseVersion.Init("domain/http.protocol; type=error; name=parse-version")
	ErrParseVersion.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Version",
		"Parsing received HTTP packet encounter unknown situation in Version part",
		"",
		"",
		nil)

	ErrParseStatusCode.Init("domain/http.protocol; type=error; name=parse-status-code")
	ErrParseStatusCode.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Status Code",
		"Parsing received HTTP packet encounter unknown situation in StatusCode part",
		"",
		"",
		nil)

	ErrParseReasonPhrase.Init("domain/http.protocol; type=error; name=parse-reason-phrase")
	ErrParseReasonPhrase.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Reason Phrase",
		"Parsing received HTTP packet encounter unknown situation in ReasonPhrase part",
		"",
		"",
		nil)

	ErrParseHeaderTooLarge.Init("domain/http.protocol; type=error; name=parse-header-too-large")
	ErrParseHeaderTooLarge.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Header Too Large",
		"Parsing received HTTP packet encounter situation that header part of http packet is larger than expected",
		"",
		"",
		nil)

	ErrCookieBadName.Init("domain/http.protocol; type=error; name=cookie-bad-name")
	ErrCookieBadName.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Name",
		"Cookie name include illegal charecter by related RFC",
		"",
		"",
		nil)

	ErrCookieBadValue.Init("domain/http.protocol; type=error; name=cookie-bad-value")
	ErrCookieBadValue.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Value",
		"Cookie value include illegal charecter by related RFC",
		"",
		"",
		nil)

	ErrCookieBadPath.Init("domain/http.protocol; type=error; name=cookie-bad-path")
	ErrCookieBadPath.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Path",
		"Cookie path include illegal charecter by related RFC",
		"",
		"",
		nil)

	ErrCookieBadDomain.Init("domain/http.protocol; type=error; name=cookie-bad-domain")
	ErrCookieBadDomain.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Domain",
		"Cookie domain is not valid by related RFC",
		"",
		"",
		nil)

	ErrNotFound.Init("domain/http.protocol; type=error; name=not-found")
	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Requested HTTP URI Service is not found in this instance of app",
		"",
		"",
		nil)

	ErrUnsupportedMediaType.Init("domain/http.protocol; type=error; name=unsupported-media-type")
	ErrUnsupportedMediaType.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Unsupported Media Type",
		"Refuse to accept the request or response because the payload format or encoding is in an unsupported format",
		"",
		"",
		nil)
}
