/* For license and copyright information please see the LEGAL file in the code repository */

package hs

import (
	er "github.com/GeniusesGroup/libgo/error"
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainEnglish = "HTTP"
const domainPersian = "HTTP"

// Declare package errors
var (
	ErrNoConnection         er.Error
	ErrNotFound             er.Error
	ErrUnsupportedMediaType er.Error
)

func init() {
	ErrNoConnection.Init("domain/http.protocol; type=error; name=no-connection")
	ErrNoConnection.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"No Connection",
		"There is no connection to peer(server or client) to process request",
		"",
		"",
		nil)
	ErrNoConnection.SetDetail(protocol.LanguagePersian, domainPersian,
		"عدم وجود ارتباط",
		"هیچ راه ارتباطی با رایانه مقصد برای پردازش درخواست مورد نظر وجود ندارد",
		"",
		"",
		nil)

	ErrNotFound.Init("domain/http.protocol; type=error; name=not-found")
	ErrNotFound.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Not Found",
		"Requested HTTP URI Service is not found in this instance of app",
		"",
		"",
		nil)

	ErrUnsupportedMediaType.Init("domain/http.protocol; type=error; name=unsupported-media-type")
	ErrUnsupportedMediaType.SetDetail(protocol.LanguageEnglish, domainEnglish,
		"Unsupported Media Type",
		"Refuse to accept the request or response because the payload format or encoding is in an unsupported format",
		"",
		"",
		nil)
}
