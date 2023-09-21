//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/detail"
	"memar/protocol"
)

const domainEnglish = "HTTP"

func init() {
	MediaType.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Hypertext Transfer Protocol",
		"An application layer protocol in the Internet protocol suite model for distributed, collaborative, hypermedia information",
		"",
		"",
		[]string{})

	MediaTypeRequest.SetDetail(protocol.LanguageEnglish, domainEnglish, "Hypertext Transfer Protocol Request", "", "", "", []string{})

	MediaTypeResponse.SetDetail(protocol.LanguageEnglish, domainEnglish, "Hypertext Transfer Protocol Response", "", "", "", []string{})
}
