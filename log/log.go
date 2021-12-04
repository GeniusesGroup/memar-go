/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"../protocol"
)

type Logger struct{}

// PanicHandler recover from panics if exist to prevent app stop.
// Call it by defer in any goroutine due to >> https://github.com/golang/go/issues/20161
func (l *Logger) PanicHandler() {
	var r = recover()
	if r != nil {
		l.LogF(protocol.LogType_Panic, "Panic Exception: %s\nDebug Stack: %s", r, debug.Stack())
	}
}

func (l *Logger) Log(logType protocol.LogType, usefulData ...interface{}) {
	if (!protocol.AppDevMode && logType == protocol.LogType_Unset) ||
		(!protocol.AppDebugMode && logType == protocol.LogType_DeepDebug) ||
		(!protocol.AppDeepDebugMode && logType == protocol.LogType_DeepDebug) {
		return
	}

	// TODO::: save log to storage with this ID: sha3.256(LogMediatypeID, NodeID, LogType, TimeRoundToDay)

	if protocol.AppScreenMode {
		var log = fmt.Sprintln("[logType]", time.Now().Format(timeFormat), usefulData)
		os.Stderr.WriteString(log)
	}
}

func (l *Logger) LogF(logType protocol.LogType, format string, v ...interface{}) {
	if (!protocol.AppDevMode && logType == protocol.LogType_Unset) ||
		(!protocol.AppDebugMode && logType == protocol.LogType_DeepDebug) ||
		(!protocol.AppDeepDebugMode && logType == protocol.LogType_DeepDebug) {
		return
	}

	// TODO::: save log to storage

	if protocol.AppScreenMode {
		// fmt.Sprintf(format, v...)
		var log = fmt.Sprintf("[logType] %s %s\n", time.Now().Format(timeFormat), v)
		os.Stderr.WriteString(log)
	}
}

func (l *Logger) LogFatal(any ...interface{}) {}

// type logHeader struct {
// 	logType protocol.LogType
// 	time    int64
// }
