/* For license and copyright information please see LEGAL file in repository */

package protocol

// Logger save logs as records in time chain for the node that create log.
// Suggest that RecordID time chain by sha3.256(LogMediatypeID, NodeID, TimeRoundToDay)
type Logger interface {
	PanicHandler()

	Log(event LogEvent)
	// Due to expect Fatal terminate app and it brake the app, Dev must design it in the app architecture with panic and log the event with Log_Fatal
	// LogFatal(event LogEvent)
}

type LogEvent interface {
	Level() LogType
	Time() TimeUnixMilli
	Domain() string
	Message() string // save formated data e.g. fmt.Sprintf("Panic Exception: %s\nDebug Stack: %s", r, debug.Stack())
	Stack() []byte   // if log need to trace, specially in panic situation
}

// LogType indicate log level
type LogType uint8

const (
	Log_Unset LogType = iota
	Log_Information
	Log_Notice
	Log_Debug // Detailed information on the flow through the system. Expect these to be written to logs only. Generally speaking, most lines logged by your application should be written as DEBUG.
	Log_DeepDebug
	Log_Warning // Use of deprecated APIs, poor use of API, 'almost' errors, other runtime situations that are undesirable or unexpected, but not necessarily "wrong". Expect these to be immediately visible on a status console.
	Log_Error   // Other runtime errors or unexpected conditions
	Log_Alert
	Log_Panic // in panic() it will add debug stack to trace more easily panic errors
	Log_Critical
	Log_Emergency
	Log_Fatal // Severe errors that cause premature termination
	Log_Security
	Log_Confidential
)

// If any below mode disabled, logger must not save that log type.
const (
	LogMode_Debug     = true
	LogMode_DeepDebug = true
)
