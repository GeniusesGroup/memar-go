/* For license and copyright information please see LEGAL file in repository */

package http

import (
	er "../error"
	lang "../language"
)

const errorEnglishDomain = "HTTP"
const errorPersianDomain = "HTTP"

// Declare Errors Details
var (
	ErrNoConnection = er.New("urn:giti:http.ietf.org:error:no-connection").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "No Connection",
		"There is no connection to peer(server or client) to proccess request",
		"",
		"").
		SetDetail(lang.LanguagePersian, errorPersianDomain, "عدم وجود ارتباط",
			"هیچ راه ارتباطی با رایانه مقصد برای پردازش درخواست مورد نظر وجود ندارد",
			"",
			"").Save()

	ErrPacketTooShort = er.New("urn:giti:http.ietf.org:error:packet-too-short").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Packet Too Short",
		"Received HTTP packet size is shorter than expected and can't use",
		"",
		"").Save()

	ErrPacketTooLong = er.New("urn:giti:http.ietf.org:error:packet-too-long").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Packet Too Long",
		"Received HTTP packet size is larger than expected and can't use",
		"",
		"").Save()

	ErrParseMethod = er.New("urn:giti:http.ietf.org:error:parse-method").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Parse Method",
		"Parsing received HTTP packet encounter unknown situation in Method part",
		"",
		"").Save()

	ErrParseURI = er.New("urn:giti:http.ietf.org:error:parse-uri").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Parse URI",
		"Parsing received HTTP packet encounter unknown situation in URI part",
		"",
		"").Save()

	ErrParseVersion = er.New("urn:giti:http.ietf.org:error:parse-version").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Parse Version",
		"Parsing received HTTP packet encounter unknown situation in Version part",
		"",
		"").Save()

	ErrParseStatusCode = er.New("urn:giti:http.ietf.org:error:parse-status-code").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Parse Status Code",
		"Parsing received HTTP packet encounter unknown situation in StatusCode part",
		"",
		"").Save()

	ErrParseReasonPhrase = er.New("urn:giti:http.ietf.org:error:parse-reason-phrase").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Parse Reason Phrase",
		"Parsing received HTTP packet encounter unknown situation in ReasonPhrase part",
		"",
		"").Save()

	ErrParseHeaderTooLarge = er.New("urn:giti:http.ietf.org:error:parse-header-too-large").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Parse Header Too Large",
		"Parsing received HTTP packet encounter situation that header part of http packet is larger than expected",
		"",
		"").Save()

	ErrCookieBadName = er.New("urn:giti:http.ietf.org:error:cookie-bad-name").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Cookie Bad Name",
		"Cookie name include illegal charecter by related RFC",
		"",
		"").Save()

	ErrCookieBadValue = er.New("urn:giti:http.ietf.org:error:cookie-bad-value").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Cookie Bad Value",
		"Cookie value include illegal charecter by related RFC",
		"",
		"").
		Save()

	ErrCookieBadPath = er.New("urn:giti:http.ietf.org:error:cookie-bad-path").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Cookie Bad Path",
		"Cookie path include illegal charecter by related RFC",
		"",
		"").Save()

	ErrCookieBadDomain = er.New("urn:giti:http.ietf.org:error:cookie-bad-domain").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Cookie Bad Domain",
		"Cookie domain is not valid by related RFC",
		"",
		"").Save()

	ErrNotFound = er.New("urn:giti:http.ietf.org:error:not-found").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Not Found",
		"Requested HTTP URI Service is not found in this instance of app",
		"",
		"").Save()

	ErrUnsupportedMediaType = er.New("urn:giti:http.ietf.org:error:unsupported-media-type").SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Unsupported Media Type",
		"Refuse to accept the request or response because the payload format or encoding is in an unsupported format",
		"",
		"").Save()
)
