/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	er "../error"
	"../mediatype"
	"../protocol"
)

const domainEnglish = "TCP"
const domainPersian = "TCP"

// Errors
var (
	ErrPacketTooShort = er.New(mediatype.New("domain/tcp.protocol.error; name=packet-too-short").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"TCP packet is empty or too short than standard header. It must include at least 20Byte header",
		"",
		"",
		nil).
		Expired(0, nil))
		
	ErrPacketWrongLength = er.New(mediatype.New("domain/tcp.protocol.error; name=packet-wrong-length").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Wrong Length",
		"Data offset set in TCP packet header is not set correctly",
		"",
		"",
		nil).
		Expired(0, nil))
)
