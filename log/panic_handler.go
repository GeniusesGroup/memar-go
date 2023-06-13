/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"fmt"

	"libgo/protocol"
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
			logEvent = ErrorEvent(&DefaultEvent_MediaType, message.Error())
		case string:
			logEvent = ErrorEvent(&DefaultEvent_MediaType, message)
		case protocol.Stringer:
			logEvent = ErrorEvent(&DefaultEvent_MediaType, message.ToString())
		case fmt.Stringer:
			logEvent = ErrorEvent(&DefaultEvent_MediaType, message.String())
		default:
			logEvent = ErrorEvent(&DefaultEvent_MediaType, fmt.Sprint(r))
		}
		Logger.Log(logEvent)
	}
}
