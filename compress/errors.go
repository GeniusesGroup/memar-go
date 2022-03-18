/* For license and copyright information please see LEGAL file in repository */

package compress

import (
	er "../error"
	"../mediatype"
	"../protocol"
)

const domainEnglish = "Compress"
const domainPersian = "فشرده سازی"

// Errors
var (
	ErrSourceNotChangeable = er.New(mediatype.New("domain/compress.protocol.error; name=source-not-changeable").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Source not Changeable",
		"Can't read from other source than source given in compression||decompression creation",
		"",
		"",
		nil))
)
