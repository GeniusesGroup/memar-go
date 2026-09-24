//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package chapar_errs

const domainEnglish = "Chapar"

func init() {
	ErrShortFrameLength.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Short Frame Length",
		"Chapar frame is too short(<12) than standard",
		"",
		"",
		nil)
	ErrLongFrameLength.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Long Frame Length",
		"Chapar frame is too long(>8192) than standard",
		"",
		"",
		nil)
	ErrMTU.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Maximum Transmission Unit - MTU",
		"Chapar frame isn't legal due to MTU is not respected by payload!",
		"",
		"",
		nil)
	ErrPortNotExist.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Port Not Exist",
		"Chapar frame can't be handle due to frame want to switch to a port that not exist in network",
		"",
		"",
		nil)
	ErrPathAlreadyUse.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Path Already Use",
		"Path already use as main chapar connection path",
		"",
		"",
		nil)
}
