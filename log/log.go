/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"fmt"

	"../event"
	"../protocol"
)

type Logger struct {
	event.EventTarget
}

// PanicHandler recover from panics if exist to prevent app stop.
// Call it by defer in any goroutine due to >> https://github.com/golang/go/issues/20161
func (l *Logger) PanicHandler() {
	var r = recover()
	if r != nil {
		var logEvent protocol.LogEvent
		switch message := r.(type) {
		case protocol.LogEvent:
			logEvent = message
		case protocol.Error:
			logEvent = PanicEvent("Unknown protocol.Error Domain", message.ToString())
		case error:
			logEvent = PanicEvent("Unknown Error Domain", message.Error())
		case string:
			logEvent = PanicEvent("Unknown string Domain", message)
		case protocol.Stringer:
			logEvent = PanicEvent("Unknown Stringer Domain", message.ToString())
		default:
			logEvent = PanicEvent("Unknown Domain", fmt.Sprint(r))
		}
		l.Log(logEvent)
	}
}

func (l *Logger) Log(event protocol.LogEvent) {
	// TODO::: it is a huge performance impact to check each logging, force caller to check before call log?
	// This func can inline means constants check on compile time?
	if (!protocol.AppMode_Dev && event.Level() == protocol.LogEvent_Unset) ||
		(!protocol.LogMode_Debug && event.Level() == protocol.LogEvent_Debug) ||
		(!protocol.LogMode_DeepDebug && event.Level() == protocol.LogEvent_DeepDebug) {
		return
	}

	l.DispatchEvent(event)
	l.saveEvent(event)
}

func (l *Logger) saveEvent(event protocol.LogEvent) {
	// TODO::: First save locally as cache to reduce network trip for non important data??

	// TODO::: Each day, Save logs to storage in one object with this ID: sha3.256(mediatype.LOG.ID(), NodeID, TimeRoundToDay)
	// protocol.App.Objects().
}
