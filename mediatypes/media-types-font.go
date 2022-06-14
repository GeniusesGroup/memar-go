/* For license and copyright information please see LEGAL file in repository */

package mediatypes

import (
	"../mediatype"
	"../protocol"
)

var (
	WOFF  mediatype.MediaType
	WOFF2 mediatype.MediaType

	TTF mediatype.MediaType
)

func init() {
	WOFF.Init("font/woff")
	WOFF.SetFileExtension("woff")
	WOFF.SetDetail(protocol.LanguageEnglish, "Web Open Font Format", "", "", "", "", []string{})

	WOFF2.Init("font/woff2")
	WOFF2.SetFileExtension("woff2")
	WOFF2.SetDetail(protocol.LanguageEnglish, "Web Open Font Format version 2", "", "", "", "", []string{})

	TTF.Init("font/ttf")
	TTF.SetFileExtension("ttf")
	TTF.SetDetail(protocol.LanguageEnglish, "TrueType Font", "", "", "", "", []string{})
}
