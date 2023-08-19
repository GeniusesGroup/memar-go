/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"fmt"

	"memar/protocol"
)

// PanicHandler recover from panics if exist to prevent app stop.
// Call it by defer in any goroutine due to >> https://github.com/golang/go/issues/20161
func PanicHandler() {
	var r = recover()
	if r != nil {
		var logEvent protocol.LogEvent
		switch message := r.(type) {
		case protocol.LogEvent:
			logEvent = message
		case protocol.Error:
			logEvent = ErrorEvent(message, message.ToString())
		case error:
			logEvent = ErrorEvent(&domain, message.Error())
		case string:
			logEvent = ErrorEvent(&domain, message)
		case protocol.Stringer:
			logEvent = ErrorEvent(&domain, message.ToString())
		case fmt.Stringer:
			logEvent = ErrorEvent(&domain, message.String())
		default:
			logEvent = ErrorEvent(&domain, fmt.Sprint(r))
		}
		Logger.Log(logEvent)
	}
}
