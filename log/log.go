/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"../protocol"
)

type Logger struct{}

// PanicHandler recover from panics if exist to prevent app stop.
// Call it by defer in any goroutine due to >> https://github.com/golang/go/issues/20161
func (l *Logger) PanicHandler() {
	var r = recover()
	if r != nil {
		l.Log(r.(protocol.LogEvent))
	}
}

func (l *Logger) Log(event protocol.LogEvent) {
	// TODO::: it is a huge performance impact to check each logging, force caller to check before call log?
	if (!protocol.AppDevMode && event.Level() == protocol.Log_Unset) ||
		(!protocol.AppDebugMode && event.Level() == protocol.Log_DeepDebug) ||
		(!protocol.AppDeepDebugMode && event.Level() == protocol.Log_DeepDebug) {
		return
	}

	// TODO::: save log to storage with this ID: sha3.256(LogMediatypeID, NodeID, TimeRoundToDay)
}
