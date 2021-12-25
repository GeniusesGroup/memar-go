/* For license and copyright information please see LEGAL file in repository */

package connection

import (
	er "../error"
	"../protocol"
)

const domainEnglish = "connection"
const domainPersian = "ارتباط"

// Errors
var (
	ErrNoConnection = er.New("urn:giti:connection.protocol:error:no-connection").SetDetail(protocol.LanguageEnglish, domainEnglish, "No Connection",
		"No connection exist to complete request due to temporary or long term problem",
		"",
		"").
		SetDetail(protocol.LanguagePersian, domainPersian, "ارتباط قطع",
			"ارتباطی جهت انجام رخواست مورد نظر بدلیل وجود مشکل موقت یا دایم وجود ندارد",
			"",
			"").Save()

	ErrSendRequest = er.New("urn:giti:connection.protocol:error:send-request").SetDetail(protocol.LanguageEnglish, domainEnglish, "Send Request",
		"Send request encounter problem due to temporary or long term problem!",
		"",
		"").Save()

	ErrProtocolHandler = er.New("urn:giti:connection.protocol:error:protocol-handler").SetDetail(protocol.LanguageEnglish, domainEnglish, "Protocol Handler",
		"Protocol handler not exist to complete the request",
		"",
		"").Save()

	ErrGuestNotAllow = er.New("urn:giti:connection.protocol:error:guest-not-allow").SetDetail(protocol.LanguageEnglish, domainEnglish, "Guest Not Allow",
		"Guest users don't allow to make new connection",
		"",
		"").Save()

	ErrGuestMaxReached = er.New("urn:giti:connection.protocol:error:guest-max-reached").SetDetail(protocol.LanguageEnglish, domainEnglish, "Guest Max Reached",
		"Server not have enough resource to make new guest connection, try few minutes later or try other server",
		"",
		"").Save()
)
