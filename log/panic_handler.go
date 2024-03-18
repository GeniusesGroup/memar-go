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
			e.Init(message.Domain(), message.LogLevel(), message.LogMessage(), message.RuntimeStack())
			logEvent = &e
		case protocol.LogEvent_Message:
			var e Event_UTF8
			e.Init(&DT, protocol.LogLevel_Error, "", true)
			e.message = message.LogMessage()
			logEvent = &e.Event
		case protocol.Error:
			var e Event_UTF8
			e.Init(message, protocol.LogLevel_Error, message.DataTypeID_string(), true)
			logEvent = &e.Event
		case error:
			var e Event_UTF8
			e.Init(&DT, protocol.LogLevel_Error, message.Error(), true)
			logEvent = &e.Event
		case string:
			var e Event_UTF8
			e.Init(&DT, protocol.LogLevel_Error, message, true)
			logEvent = &e.Event
		case protocol.Stringer_To[protocol.String]:
			var msgStr, _ = message.ToString()
			var e Event_UTF8
			e.Init(&DT, protocol.LogLevel_Error, "", true)
			e.message = msgStr
			logEvent = &e.Event
		case fmt.Stringer:
			var e Event_UTF8
			e.Init(&DT, protocol.LogLevel_Error, message.String(), true)
			logEvent = &e.Event
		default:
			var e Event_UTF8
			e.Init(&DT, protocol.LogLevel_Error, fmt.Sprint(r), true)
			logEvent = &e.Event
		}

		Logger.DispatchEvent(logEvent)
	}
}
