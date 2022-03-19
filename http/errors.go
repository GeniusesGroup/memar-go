/* For license and copyright information please see LEGAL file in repository */

package http

import (
	er "../error"
	"../mediatype"
	"../protocol"
)

const domainEnglish = "HTTP"
const domainPersian = "HTTP"

// Declare Errors Details
var (
	ErrNoConnection = er.New(mediatype.New("domain/http.protocol.error; name=no-connection").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"No Connection",
		"There is no connection to peer(server or client) to proccess request",
		"",
		"",
		nil).SetDetail(protocol.LanguagePersian, domainPersian,
		"عدم وجود ارتباط",
		"هیچ راه ارتباطی با رایانه مقصد برای پردازش درخواست مورد نظر وجود ندارد",
		"",
		"",
		nil))

	ErrPacketTooShort = er.New(mediatype.New("domain/http.protocol.error; name=packet-too-short").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"Received HTTP packet size is shorter than expected and can't use",
		"",
		"",
		nil))

	ErrPacketTooLong = er.New(mediatype.New("domain/http.protocol.error; name=packet-too-long").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Long",
		"Received HTTP packet size is larger than expected and can't use",
		"",
		"",
		nil))

	ErrParseMethod = er.New(mediatype.New("domain/http.protocol.error; name=parse-method").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Method",
		"Parsing received HTTP packet encounter unknown situation in Method part",
		"",
		"",
		nil))

	ErrParseURI = er.New(mediatype.New("domain/http.protocol.error; name=parse-uri").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse URI",
		"Parsing received HTTP packet encounter unknown situation in URI part",
		"",
		"",
		nil))

	ErrParseVersion = er.New(mediatype.New("domain/http.protocol.error; name=parse-version").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Version",
		"Parsing received HTTP packet encounter unknown situation in Version part",
		"",
		"",
		nil))

	ErrParseStatusCode = er.New(mediatype.New("domain/http.protocol.error; name=parse-status-code").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Status Code",
		"Parsing received HTTP packet encounter unknown situation in StatusCode part",
		"",
		"",
		nil))

	ErrParseReasonPhrase = er.New(mediatype.New("domain/http.protocol.error; name=parse-reason-phrase").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Reason Phrase",
		"Parsing received HTTP packet encounter unknown situation in ReasonPhrase part",
		"",
		"",
		nil))

	ErrParseHeaderTooLarge = er.New(mediatype.New("domain/http.protocol.error; name=parse-header-too-large").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Parse Header Too Large",
		"Parsing received HTTP packet encounter situation that header part of http packet is larger than expected",
		"",
		"",
		nil))

	ErrCookieBadName = er.New(mediatype.New("domain/http.protocol.error; name=cookie-bad-name").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Name",
		"Cookie name include illegal charecter by related RFC",
		"",
		"",
		nil))

	ErrCookieBadValue = er.New(mediatype.New("domain/http.protocol.error; name=cookie-bad-value").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Value",
		"Cookie value include illegal charecter by related RFC",
		"",
		"",
		nil))

	ErrCookieBadPath = er.New(mediatype.New("domain/http.protocol.error; name=cookie-bad-path").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Path",
		"Cookie path include illegal charecter by related RFC",
		"",
		"",
		nil))

	ErrCookieBadDomain = er.New(mediatype.New("domain/http.protocol.error; name=cookie-bad-domain").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Cookie Bad Domain",
		"Cookie domain is not valid by related RFC",
		"",
		"",
		nil))

	ErrNotFound = er.New(mediatype.New("domain/http.protocol.error; name=not-found").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Requested HTTP URI Service is not found in this instance of app",
		"",
		"",
		nil))

	ErrUnsupportedMediaType = er.New(mediatype.New("domain/http.protocol.error; name=unsupported-media-type").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Unsupported Media Type",
		"Refuse to accept the request or response because the payload format or encoding is in an unsupported format",
		"",
		"",
		nil))
)
