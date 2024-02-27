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
		var logEvent *Event
		switch message := r.(type) {
		case *Event:
			logEvent = message
		case protocol.LogEvent:
			var e Event
			e.Init(message.Domain(), message.Level(), message.Message(), message.Stack())
			logEvent = &e
		case protocol.Error:
			var msgStr, _ = message.ToString()
			logEvent = ErrorEvent(message, msgStr)
		case error:
			logEvent = ErrorEvent(&DT, nil)
			logEvent.MSG_string(message.Error())
		case string:
			logEvent = ErrorEvent(&DT, nil)
			logEvent.MSG_string(message)
		case protocol.Stringer[protocol.String]:
			var msgStr, _ = message.ToString()
			logEvent = ErrorEvent(&DT, msgStr)
		case fmt.Stringer:
			logEvent = ErrorEvent(&DT, nil)
			logEvent.MSG_string(message.String())
		default:
			logEvent = ErrorEvent(&DT, nil)
			logEvent.MSG_string(fmt.Sprint(r))
		}
		Logger.DispatchEvent(logEvent)
	}
}
