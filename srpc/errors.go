/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	er "../error"
	lang "../language"
)

const errorEnglishDomain = "sRPC"
const errorPersianDomain = "sRPC"

// Errors
var (
	ErrPacketTooShort = er.New("urn:giti:srpc.libgo:error:packet-too-short").
				SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Packet Too Short", "Received sRPC Packet size is smaller than expected and can't use").
				Save()

	ErrServiceNotFound = er.New("urn:giti:srpc.libgo:error:service-not-found").
				SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Service Not Found", "Requested service by given ID not found in this server").
				Save()
)
