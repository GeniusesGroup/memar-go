/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/event"
)

// Logger is default global  like window.console.log global variable in browsers.
// As suggested in protocol.Logger document, You can listen to notify about any log events occur.
var Logger event.EventTarget[*Event]
