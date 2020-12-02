/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	er "../error"
	lang "../language"
)

const errorEnglishDomain = "Achaemenid"
const errorPersianDomain = "فرمانده"

// Errors
var (
	ErrNoConnection = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "No Connection",
		"No connection exist to complete request due to temporary or long term problem").Save()

	ErrSendRequest = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Send Request",
		"Send request encounter problem due to temporary or long term problem!").Save()

	ErrReceiveResponse = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Receive Response",
		"Receive response encounter problem due to temporary or long term problem!").Save()

	ErrBadRequest = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Bad Request",
		"Some given data in request must be invalid or peer not accept them").Save()

	ErrBadResponse = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Bad Response",
		"Response data from peer is not valid").Save()

	ErrInternalError = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Internal Error",
		"Peer encounter problem due to temporary or long term problem!").Save()

	ErrProtocolHandler = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Protocol Handler",
		"Protocol handler not exist to complete the request").Save()

	ErrGuestConnectionNotAllow = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Guest Connection Not Allow",
		"Guest users don't allow to make new connection").Save()

	ErrGuestConnectionMaxReached = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Guest Connection Max Reached",
		"Server not have enough resource to make new guest connection, try few minutes later or try other server").Save()
)
