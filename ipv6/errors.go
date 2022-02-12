/* For license and copyright information please see LEGAL file in repository */

package ipv6

import (
	er "../error"
	"../mediatype"
	"../protocol"
)

const domainEnglish = "IPv6"
const domainPersian = "IPv6"

// Errors
var (
	ErrPacketTooShort = er.New(mediatype.New("domain/ipv6.protocol.error; name=packet-too-short").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"IPv6 packet is empty or too short than standard minimum size. It must include at least 40Byte header",
		"",
		"",
		nil).
		Expired(0, nil))
)
