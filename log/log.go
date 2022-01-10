/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"fmt"

	"../protocol"
)

type Logger struct{}

// PanicHandler recover from panics if exist to prevent app stop.
// Call it by defer in any goroutine due to >> https://github.com/golang/go/issues/20161
func (l *Logger) PanicHandler() {
	var r = recover()
	if r != nil {
		var logEvent, ok = r.(protocol.LogEvent)
		if !ok {
			switch message := r.(type) {
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
		}
		l.Log(logEvent)
	}
}

func (l *Logger) Log(event protocol.LogEvent) {
	// TODO::: it is a huge performance impact to check each logging, force caller to check before call log?
	// This func can inline means constants check on compile time?
	if (!protocol.AppMode_Dev && event.Level() == protocol.Log_Unset) ||
		(!protocol.LogMode_Debug && event.Level() == protocol.Log_Debug) ||
		(!protocol.LogMode_DeepDebug && event.Level() == protocol.Log_DeepDebug) {
		return
	}

	l.saveLog(event)
}

func (l *Logger) saveLog(event protocol.LogEvent) {
	// TODO::: Implement some mechanism to listen to important logs, ...

	// TODO::: First save locally as cache to reduce network trip for non important data??

	// TODO::: Each day, Save logs to storage in one object with this ID: sha3.256(mediatype.LOG.ID(), NodeID, TimeRoundToDay)
	// protocol.App.Objects().
}
