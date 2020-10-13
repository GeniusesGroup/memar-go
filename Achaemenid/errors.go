/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	errorr "../error"
	lang "../language"
)

// Errors
var (
	ErrAchaemenidNoConnectionToNode = errorr.New().
					SetDetail(lang.EnglishLanguage, "Achaemenid - No Connection To Node",
			"There is no connection to desire node to complete request").Save()

	ErrAchaemenidProtocolHandler = errorr.New().
					SetDetail(lang.EnglishLanguage, "Achaemenid - Protocol Handler",
			"Protocol handler not exist to complete the request").Save()

	ErrAchaemenidGuestConnectionNotAllow = errorr.New().
						SetDetail(lang.EnglishLanguage, "Achaemenid - Guest Connection Not Allow",
			"Guest users don't allow to make new connection").Save()

	ErrAchaemenidGuestConnectionMaxReached = errorr.New().
						SetDetail(lang.EnglishLanguage, "Achaemenid - Guest Connection Max Reached",
			"This server not have enough resource to make new guest connection, register or try other server").Save()
)
