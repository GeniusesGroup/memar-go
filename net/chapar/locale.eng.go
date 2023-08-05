//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "Chapar"

func init() {
	ErrShortFrameLength.SetDetail(protocol.LanguageEnglish, domainEnglish, "Short Frame Length",
		"Chapar frame is too short(<12) than standard",
		"",
		"",
		nil)
	ErrLongFrameLength.SetDetail(protocol.LanguageEnglish, domainEnglish, "Long Frame Length",
		"Chapar frame is too long(>8192) than standard",
		"",
		"",
		nil)
	ErrMTU.SetDetail(protocol.LanguageEnglish, domainEnglish, "Maximum Transmission Unit - MTU",
		"Chapar frame isn't legal due to MTU is not respected by payload!",
		"",
		"",
		nil)
	ErrPortNotExist.SetDetail(protocol.LanguageEnglish, domainEnglish, "Port Not Exist",
		"Chapar frame can't be handle due to frame want to switch to a port that not exist in network",
		"",
		"",
		nil)
	ErrPathAlreadyUse.SetDetail(protocol.LanguageEnglish, domainEnglish, "Path Already Use",
		"Path already use as main chapar connection path",
		"",
		"",
		nil)
}
