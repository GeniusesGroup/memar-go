//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/protocol"
)

const domainEnglish = "TCP"

func init() {
	ErrSegmentTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Segment Too Short",
		"TCP packet is empty or too short than standard header. It must include at least 20Byte header",
		"",
		"",
		nil)
	ErrSegmentWrongLength.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Segment Wrong Length",
		"Data offset set in TCP packet header is not set correctly",
		"",
		"",
		nil)
}
