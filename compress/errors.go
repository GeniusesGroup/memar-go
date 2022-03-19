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
	ErrNotFound = er.New(mediatype.New("domain/compress.protocol.error; name=not-found").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Can't find requested compression||decompression algorithm",
		"",
		"",
		nil))

	ErrSourceNotChangeable = er.New(mediatype.New("domain/compress.protocol.error; name=source-not-changeable").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Source not Changeable",
		"Can't read from other source than source given in compression||decompression creation",
		"",
		"",
		nil))
)
