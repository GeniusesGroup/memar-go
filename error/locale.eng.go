//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainEnglish = "Error"

func init() {
	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"An error occurred but it is not registered yet to show more detail to you!",
		"Sorry it's us not your fault! Contact administrator of platform!",
		"Find error by its URN and save it for further use by any UserInterfaces",
		nil)

	ErrNotExist.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Exist",
		"Given Error is not exist",
		"Sorry it's us not your fault! Contact administrator of platform",
		"Trace error by enable panic recovery to find nil error detection problem",
		nil)
}
