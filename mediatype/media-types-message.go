/* For license and copyright information please see LEGAL file in repository */

package mediatype

import "../protocol"

var (
	EML = New("message/rfc822", "eml").
		SetDetail(protocol.LanguageEnglish, "", "", []string{})
)
