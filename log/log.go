/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"fmt"

	"github.com/GeniusesGroup/libgo/event"
	"github.com/GeniusesGroup/libgo/protocol"
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
			logEvent = PanicEvent("Unknown error Domain", message.Error())
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

func (l *Logger) Log(event protocol.LogEvent) (err protocol.Error) {
	l.DispatchEvent(event)
	// Due to LogEvent is not cancelable, we don't need to check event.DefaultPrevented() before save log.
	// if event.DefaultPrevented() {}
	err = l.saveEvent(event)
	return
}

func (l *Logger) saveEvent(event protocol.LogEvent) (err protocol.Error) {
	if !CheckLevelEnabled(event.Level()) {
		return
	}

	// TODO::: First save locally as cache to reduce network trip for non important data??

	// TODO::: Each day, Save logs to storage(mediatype.LOG.ID()) with NodeID as object ID
	// protocol.App.Objects().
	return
}
