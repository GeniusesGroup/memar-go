//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
	"libgo/protocol"
)

const domainEnglish = "HTTP"

func init() {
	ErrBadRequest.SetDetail(protocol.LanguageEnglish, domainEnglish, "Bad Request",
		"Some given data in request must be invalid or peer not accept them",
		"",
		"",
		nil)

	ErrBadResponse.SetDetail(protocol.LanguageEnglish, domainEnglish, "Bad Response",
		"Response data from peer is not valid",
		"",
		"",
		nil)

	ErrInternalError.SetDetail(protocol.LanguageEnglish, domainEnglish, "Internal Error",
		"Peer encounter problem due to temporary or long term problem!",
		"",
		"",
		nil)

	ErrProtocolHandler.SetDetail(protocol.LanguageEnglish, domainEnglish, "Protocol Handler",
		"Protocol handler not exist to complete the request",
		"",
		"",
		nil)

	ErrGuestConnectionNotAllow.SetDetail(protocol.LanguageEnglish, domainEnglish, "Guest Connection Not Allow",
		"Guest users don't allow to make new connection",
		"",
		"",
		nil)

	ErrGuestConnectionMaxReached.SetDetail(protocol.LanguageEnglish, domainEnglish, "Guest Connection Max Reached",
		"Server not have enough resource to make new guest connection, try few minutes later or try other server",
		"",
		"",
		nil)

	ErrNotStandardID.SetDetail(protocol.LanguageEnglish, domainEnglish, "Not Standard ID",
		"You set non standard ID for error||service||data-structure||..., It can cause some bad situation in your platform",
		"",
		"",
		nil)
}
