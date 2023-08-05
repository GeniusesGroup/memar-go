//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package ipv6

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "IPv6"

func init() {
	ErrPacketTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"IPv6 packet is empty or too short than standard minimum size. It must include at least 40Byte header",
		"",
		"",
		nil)
}
