/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"libgo/protocol"
)

// Logger is default global protocol.Logger like window.console.log global variable in browsers.
// You must assign to it by any object implement protocol.Logger on your main.go file. Suggestion:
// log.Logger = &esteghrar.Logger
var Logger protocol.Logger
