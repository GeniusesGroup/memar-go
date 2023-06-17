//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package connection

import (
	"libgo/protocol"
)

const domainEnglish = "connection"

func init() {
	ErrNoConnection.SetDetail(protocol.LanguageEnglish, domainEnglish, "No Connection",
		"No connection exist to complete request due to temporary or long term problem",
		"",
		"",
		nil)

	ErrSendRequest.SetDetail(protocol.LanguageEnglish, domainEnglish, "Send Request",
		"Send request encounter problem due to temporary or long term problem!",
		"",
		"",
		nil)

	ErrProtocolHandler.SetDetail(protocol.LanguageEnglish, domainEnglish, "Protocol Handler",
		"Protocol handler not exist to complete the request",
		"",
		"",
		nil)

	ErrGuestNotAllow.SetDetail(protocol.LanguageEnglish, domainEnglish, "Guest Not Allow",
		"Guest users don't allow to make new connection",
		"",
		"",
		nil)

	ErrGuestMaxReached.SetDetail(protocol.LanguageEnglish, domainEnglish, "Guest Max Reached",
		"Server not have enough resource to make new guest connection, try few minutes later or try other server",
		"",
		"",
		nil)
}
