/* For license and copyright information please see LEGAL file in repository */

package udp

import (
	er "../error"
	"../protocol"
)

const domainEnglish = "UDP"
const domainPersian = "UDP"

// Errors
var (
	ErrPacketTooShort    er.Error
	ErrPacketWrongLength er.Error
)

func init() {
	ErrPacketTooShort.Init("domain/udp.protocol.error; name=packet-too-short")
	ErrPacketTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"UDP packet is empty or too short than standard header. It must include at least 20Byte header",
		"",
		"",
		nil)

	ErrPacketWrongLength.Init("domain/udp.protocol.error; name=packet-wrong-length")
	ErrPacketWrongLength.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Wrong Length",
		"Data offset set in UDP packet header is not set correctly",
		"",
		"",
		nil)
}
