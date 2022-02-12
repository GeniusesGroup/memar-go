/* For license and copyright information please see LEGAL file in repository */

package ipv4

import (
	er "../error"
	"../mediatype"
	"../protocol"
)

const domainEnglish = "IPv4"
const domainPersian = "IPv4"

// Errors
var (
	ErrPacketTooShort = er.New(mediatype.New("domain/ipv4.protocol.error; name=packet-too-short").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Too Short",
		"IPv4 packet is empty or too short than standard minimum size. It must include at least 20Byte header",
		"",
		"",
		nil).
		Expired(0, nil))

	ErrPacketWrongLength = er.New(mediatype.New("domain/ipv4.protocol.error; name=packet-wrong-length").SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Packet Wrong Length",
		"Data offset set in IPv4 packet header is not set correctly",
		"",
		"",
		nil).
		Expired(0, nil))
)
