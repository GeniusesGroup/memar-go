/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	er "../../error"
	lang "../../language"
)

func getErrorByResCode(code int64) (err *er.Error) {
	switch code {
	case 0:
		return nil
	case 1:
		return ErrNoActionAfterReadCard
	case 2:
		return ErrAmountMinimum
	case 3:
		return ErrPOSNotReachable
	case 4:
		return ErrPOSNotValidData
	case 9:
		return ErrPOSNotValidTerminalID
	case 51:
		return ErrPOSNotEnoughBalance
	case 55:
		return ErrPOSCardPassword
	case 96:
		return ErrPOSNotIndicated
	case 99:
		return ErrPOSNotResponse
	default:
		return ErrInternalError
	}
}

// Errors
var (
	ErrBadPOSSettings = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - Bad POS Settings",
		"Can't find 'sep.ir-pos.json' file in 'secret' folder in top of repository or has not enough information").Save()

	ErrBadTerminalID = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - Bad Terminal ID",
		"Requested Terminal ID is not valid by platform settings").Save()

	ErrNoConnection = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - No Connection",
		"No connection exist to SEP servers due to temporary or long term problem").Save()

	ErrBadRequest = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - Bad Request",
		"Some given data in request must be invalid or server not accept them").Save()

	ErrInternalError = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - Internal Error",
		"SEP server encounter problem due to temporary or long term problem!").Save()

	ErrBadResponse = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - Bad Response",
		"Response data from SEP server is not valid").Save()

	// Server send errors

	ErrNoActionAfterReadCard = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - No Action After Read Card",
		"Transaction canceled due to no action received after read card by POS device").Save()

	ErrAmountMinimum = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - Amount Minimum",
		"Transaction canceled due to below legal minimum amount sent").Save()

	ErrPOSNotReachable = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - POS Not Reachable",
		"Transaction canceled due to selected POS not reachable").Save()

	ErrPOSNotValidData = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - POS Not Valid Data",
		"Transaction canceled due to not valid data received").Save()

	ErrPOSNotValidTerminalID = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - POS Not Valid TerminalID",
		"Transaction canceled due to TerminalID not valid").Save()

	ErrPOSNotEnoughBalance = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - POS Not Enough Balance",
		"Transaction canceled due to user don't has enough balance").Save()

	ErrPOSCardPassword = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - POS Card Password",
		"Transaction canceled due to user not enter right password of its card").Save()

	ErrPOSNotIndicated = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - POS Not Indicated",
		"Transaction canceled due to not indicate POS order service error").Save()

	ErrPOSNotResponse = er.New().SetDetail(lang.EnglishLanguage, "sep.ir - POS Not Response",
		"Transaction canceled due to POS not response in proper time").Save()
)
