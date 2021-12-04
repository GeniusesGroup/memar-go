/* For license and copyright information please see LEGAL file in repository */

package protocol

// Logger save logs as objects in time chain for the node that create log.
// Suggest that ObjectID create by sha3.256(LogMediatypeID, NodeID, LogType, TimeRoundToDay)
type Logger interface {
	PanicHandler()

	Log(logType LogType, usefulData ...interface{})
	LogF(logType LogType, format string, v ...interface{})
	// It will terminate app. if dev don't wat to terminate app, just log error with LogType_Fatal
	LogFatal(any ...interface{})
}

// LogType indicate log level
type LogType uint8

const (
	LogType_Unset LogType = iota
	LogType_Confidential
	LogType_Information
	LogType_Notice
	LogType_Warning
	LogType_Error
	LogType_Debug
	LogType_DeepDebug
	LogType_Panic // in panic() it will add debug stack to trace more easily panic errors
	LogType_Critical
	LogType_Alert
	LogType_Emergency
	LogType_Fatal
)
