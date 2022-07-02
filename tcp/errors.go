/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	er "../error"
	"../protocol"
)

const domainEnglish = "TCP"
const domainPersian = "TCP"

// Errors
var (
	ErrPacketTooShort    er.Error
	ErrPacketWrongLength er.Error
)

func init() {
	ErrPacketTooShort.Init("domain/tcp.protocol.error; name=packet-too-short")
	ErrPacketTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"TCP packet is empty or too short than standard header. It must include at least 20Byte header",
		"",
		"",
		nil)

	ErrPacketWrongLength.Init("domain/tcp.protocol.error; name=packet-wrong-length")
	ErrPacketTooShort.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Wrong Length",
		"Data offset set in TCP packet header is not set correctly",
		"",
		"",
		nil)
}
