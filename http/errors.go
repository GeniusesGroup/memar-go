/* For license and copyright information please see LEGAL file in repository */

package http

import (
	er "../error"
	"../protocol"
)

const errorEnglishDomain = "HTTP"
const errorPersianDomain = "HTTP"

// Declare Errors Details
var (
	ErrNoConnection = er.New("urn:giti:http.protocol:error:no-connection").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "No Connection",
		"There is no connection to peer(server or client) to proccess request",
		"",
		"").
		SetDetail(protocol.LanguagePersian, errorPersianDomain, "عدم وجود ارتباط",
			"هیچ راه ارتباطی با رایانه مقصد برای پردازش درخواست مورد نظر وجود ندارد",
			"",
			"").Save()

	ErrPacketTooShort = er.New("urn:giti:http.protocol:error:packet-too-short").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Packet Too Short",
		"Received HTTP packet size is shorter than expected and can't use",
		"",
		"").Save()

	ErrPacketTooLong = er.New("urn:giti:http.protocol:error:packet-too-long").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Packet Too Long",
		"Received HTTP packet size is larger than expected and can't use",
		"",
		"").Save()

	ErrParseMethod = er.New("urn:giti:http.protocol:error:parse-method").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Parse Method",
		"Parsing received HTTP packet encounter unknown situation in Method part",
		"",
		"").Save()

	ErrParseURI = er.New("urn:giti:http.protocol:error:parse-uri").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Parse URI",
		"Parsing received HTTP packet encounter unknown situation in URI part",
		"",
		"").Save()

	ErrParseVersion = er.New("urn:giti:http.protocol:error:parse-version").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Parse Version",
		"Parsing received HTTP packet encounter unknown situation in Version part",
		"",
		"").Save()

	ErrParseStatusCode = er.New("urn:giti:http.protocol:error:parse-status-code").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Parse Status Code",
		"Parsing received HTTP packet encounter unknown situation in StatusCode part",
		"",
		"").Save()

	ErrParseReasonPhrase = er.New("urn:giti:http.protocol:error:parse-reason-phrase").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Parse Reason Phrase",
		"Parsing received HTTP packet encounter unknown situation in ReasonPhrase part",
		"",
		"").Save()

	ErrParseHeaderTooLarge = er.New("urn:giti:http.protocol:error:parse-header-too-large").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Parse Header Too Large",
		"Parsing received HTTP packet encounter situation that header part of http packet is larger than expected",
		"",
		"").Save()

	ErrCookieBadName = er.New("urn:giti:http.protocol:error:cookie-bad-name").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Cookie Bad Name",
		"Cookie name include illegal charecter by related RFC",
		"",
		"").Save()

	ErrCookieBadValue = er.New("urn:giti:http.protocol:error:cookie-bad-value").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Cookie Bad Value",
		"Cookie value include illegal charecter by related RFC",
		"",
		"").
		Save()

	ErrCookieBadPath = er.New("urn:giti:http.protocol:error:cookie-bad-path").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Cookie Bad Path",
		"Cookie path include illegal charecter by related RFC",
		"",
		"").Save()

	ErrCookieBadDomain = er.New("urn:giti:http.protocol:error:cookie-bad-domain").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Cookie Bad Domain",
		"Cookie domain is not valid by related RFC",
		"",
		"").Save()

	ErrNotFound = er.New("urn:giti:http.protocol:error:not-found").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Not Found",
		"Requested HTTP URI Service is not found in this instance of app",
		"",
		"").Save()

	ErrUnsupportedMediaType = er.New("urn:giti:http.protocol:error:unsupported-media-type").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Unsupported Media Type",
		"Refuse to accept the request or response because the payload format or encoding is in an unsupported format",
		"",
		"").Save()
)
