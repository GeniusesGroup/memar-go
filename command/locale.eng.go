//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	er "github.com/GeniusesGroup/libgo/error"
	"github.com/GeniusesGroup/libgo/protocol"
)

const domainEnglish = "Command"

func init() {
	ErrServiceNotAcceptCLI.SetDetail(protocol.LanguageEnglish, domainEnglish, "Service Not Accept CLI",
		"Requested service not accept CLI protocol in this server",
		"Try other server or contact support of the software",
		"It is so easy to implement CLI handler for a service! Take a time and do it!",
		nil)
}
