/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	er "../error"
	"../protocol"
)

const errorEnglishDomain = "Achaemenid"
const errorPersianDomain = "فرمانده"

// Errors
var (
	ErrNoConnection = er.New("urn:giti:achaemenid.giti:error:no-connection").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "No Connection",
		"No connection exist to complete request due to temporary or long term problem",
		"",
		"").
		SetDetail(protocol.LanguagePersian, errorPersianDomain, "ارتباط قطع",
			"ارتباطی جهت انجام رخواست مورد نظر بدلیل وجود مشکل موقت یا دایم وجود ندارد",
			"",
			"").Save()

	ErrSendRequest = er.New("urn:giti:achaemenid.giti:error:send-request").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Send Request",
		"Send request encounter problem due to temporary or long term problem!",
		"",
		"").Save()

	ErrReceiveResponse = er.New("urn:giti:achaemenid.giti:error:receive-response").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Receive Respone",
		"Receive response encounter problem due to temporary or long term problem!",
		"",
		"").Save()

	ErrBadRequest = er.New("urn:giti:achaemenid.giti:error:bad-request").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Bad Request",
		"Some given data in request must be invalid or peer not accept them",
		"",
		"").Save()

	ErrBadResponse = er.New("urn:giti:achaemenid.giti:error:bad-response").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Bad Response",
		"Response data from peer is not valid",
		"",
		"").Save()

	ErrInternalError = er.New("urn:giti:achaemenid.giti:error:internal-error").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Internal Error",
		"Peer encounter problem due to temporary or long term problem!",
		"",
		"").Save()

	ErrProtocolHandler = er.New("urn:giti:achaemenid.giti:error:protocol-handler").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Protocol Handler",
		"Protocol handler not exist to complete the request",
		"",
		"").Save()

	ErrGuestConnectionNotAllow = er.New("urn:giti:achaemenid.giti:error:guest-connection-not-allow").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Guest Connection Not Allow",
		"Guest users don't allow to make new connection",
		"",
		"").Save()

	ErrGuestConnectionMaxReached = er.New("urn:giti:achaemenid.giti:error:guest-connection-max-reached").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Guest Connection Max Reached",
		"Server not have enough resource to make new guest connection, try few minutes later or try other server",
		"",
		"").Save()

		ErrNotStandardID = er.New("urn:giti:giti:error:not-standard-id").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Not Standard ID",
		"You set non standard ID for error||service||data-structure||..., It can cause some bad situation in your platform",
		"",
		"").Save()
)
