/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	WOFF = New("font/woff", "woff").
		SetDetail(protocol.LanguageEnglish, "Web Open Font Format", "", []string{})
	WOFF2 = New("font/woff2", "woff2").
		SetDetail(protocol.LanguageEnglish, "Web Open Font Format version 2", "", []string{})

	TTF = New("font/ttf", "ttf").
		SetDetail(protocol.LanguageEnglish, "TrueType Font", "", []string{})
)
