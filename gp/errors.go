/* For license and copyright information please see LEGAL file in repository */

package gp

import (
	er "../error"
	"../protocol"
)

const errorEnglishDomain = "Giti Network"
const errorPersianDomain = "شبکه گیتی"

// Errors
var (
	ErrPacketTooShort = er.New("urn:giti:gp.protocol:error:packet-too-short").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Packet Too Short",
		"Giti packet is empty or too short than standard header. It must include 44Byte header plus 16Byte min Payload",
		"",
		"").Save()

	ErrPacketArrivedAnterior = er.New("urn:giti:gp.protocol:error:packet-arrived-anterior").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Packet Arrived Anterior",
		"New packet arrive before some expected packet arrived. Usually cause of drop packet detection or high latency occur for some packet",
		"",
		"").Save()

	ErrPacketArrivedPosterior = er.New("urn:giti:gp.protocol:error:packet-arrived-posterior").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Packet Arrived Posterior",
		"New packet arrive after some expected packet arrived. Usually cause of drop packet detection or high latency occur for some packet",
		"",
		"").Save()
)
