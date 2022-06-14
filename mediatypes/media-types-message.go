/* For license and copyright information please see LEGAL file in repository */

package mediatypes

import (
	"../mediatype"
	"../protocol"
)

var (
	EML mediatype.MediaType
)

func init() {
	EML.Init("message/rfc822")
	EML.SetFileExtension("eml")
	EML.SetDetail(protocol.LanguageEnglish, "", "", "", "", "", []string{})
}
