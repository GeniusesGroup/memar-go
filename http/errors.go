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
	ErrNoConnection = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "No Connection",
		"There is no connection to peer(server or client) to proccess request").SetDetail(
		lang.PersianLanguage, errorPersianDomain, "عدم وجود ارتباط",
		"هیچ راه ارتباطی با رایانه مقصد برای پردازش درخواست مورد نظر وجود ندارد").Save()

	ErrPacketTooShort = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Packet Too Short",
		"Received HTTP packet size is smaller than expected and can't use").Save()

	ErrPacketTooLong = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Packet Too Long",
		"Received HTTP packet size is larger than expected and can't use").Save()

	ErrParseMethod = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Parse Method",
		"Parsing received HTTP packet encounter unknown situation in Method part").Save()

	ErrParseURI = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Parse URI",
		"Parsing received HTTP packet encounter unknown situation in URI part").Save()

	ErrParseVersion = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Parse Version",
		"Parsing received HTTP packet encounter unknown situation in Version part").Save()

	ErrParseStatusCode = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Parse Status Code",
		"Parsing received HTTP packet encounter unknown situation in StatusCode part").Save()

	ErrParseReasonPhrase = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Parse Reason Phrase",
		"Parsing received HTTP packet encounter unknown situation in ReasonPhrase part").Save()

	ErrParseHeaderTooLarge = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Parse Header Too Large",
		"Parsing received HTTP packet encounter situation that header part of http packet is larger than expected").Save()

	ErrCookieBadName = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Cookie Bad Name",
		"Cookie name include illegal charecter by related RFC").Save()

	ErrCookieBadValue = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Cookie Bad Value",
		"Cookie value include illegal charecter by related RFC").Save()

	ErrCookieBadPath = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Cookie Bad Path",
		"Cookie path include illegal charecter by related RFC").Save()

	ErrCookieBadDomain = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Cookie Bad Domain",
		"Cookie domain is not valid by related RFC").Save()

	ErrNotFound = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Not Found",
		"Requested HTTP URI Service is not found in this instance of app").Save()

	ErrUnsupportedMediaType = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Unsupported Media Type",
		"Refuse to accept the request or response because the payload format or encoding is in an unsupported format").Save()
)
