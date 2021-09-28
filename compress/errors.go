/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	er "../error"
	"../protocol"
)

const errorEnglishDomain = "Compress"
const errorPersianDomain = "فشرده سازی"

// Errors
var (
	ErrSourceNotChangeable = er.New("urn:giti:compress.protocol:error:source-not-changeable").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Source not Changeable",
		"Can't read from other source than source given in compression||decompression creation",
		"",
		"").Save()
)
