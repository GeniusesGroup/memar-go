//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package errs

const domainEnglish = "Achaemenid"

func init() {
	ErrSDKNotFound.SetDetail(lang_p.LanguageEnglish, domainEnglish, "SDK Not Found",
		"Requested software errors SDK in desire human language not found",
		"Contact administrator of software and ask add needed data",
		"You can use online translator and desire human language to errors details")

	ErrNotStandardStructureID.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Not Standard StructureID",
		"Given structure to register set non standard ID, It can cause some bad situation in your platform",
		"",
		"")

	ErrFileNameIsNotValid.SetDetail(lang_p.LanguageEnglish, domainEnglish, "File Name Is Not Valid",
		"Requested file name is not exist or in bad shape",
		"",
		"")

	ErrServiceNameIsNotValid.SetDetail(lang_p.LanguageEnglish, domainEnglish, "Service Name Is Not Valid",
		"Requested service name is not exist or in bad shape",
		"",
		"")
}
