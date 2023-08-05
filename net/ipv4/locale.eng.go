//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package ipv4

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "IPv4"

func init() {
	ErrPacketTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"IPv4 packet is empty or too short than standard minimum size. It must include at least 20Byte header",
		"",
		"",
		nil)
	ErrPacketWrongLength.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Wrong Length",
		"Data offset set in IPv4 packet header is not set correctly",
		"",
		"",
		nil)
}
