//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package hs

import (
	"libgo/protocol"
)

const domainEnglish = "HTTP Services"

func init() {
	ErrNoConnection.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"No Connection",
		"There is no connection to peer(server or client) to process request",
		"",
		"",
		nil)
	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Requested HTTP URI Service is not found in this instance of app",
		"",
		"",
		nil)
	ErrUnsupportedMediaType.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Unsupported Media Type",
		"Refuse to accept the request or response because the payload format or encoding is in an unsupported format",
		"",
		"",
		nil)
}

func init() {
	ServeWWWService.DS.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Serve WWW",
		"",
		"",
		"",
		nil)
	MuxService.DS.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Service Multiplexer",
		"Multiplex services by its ID with impressive performance",
		"",
		"",
		nil)
	HostSupportedService.DS.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Host Supported",
		"Service to check if requested host is valid or not",
		"",
		"",
		nil)
}
