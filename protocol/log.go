/* For license and copyright information please see LEGAL file in repository */

package protocol

// Logger save logs as records in time chain for the node that create log.
// Suggest that RecordID time chain by sha3.256(LogMediatypeID, NodeID, TimeRoundToDay)
type Logger interface {
	// PanicHandler recover from panics in a goroutine if exist, to prevent the application unexpected stopping.
	PanicHandler()

	Log(event LogEvent)
	// Due to expect Fatal terminate app and it brake the app, Dev must design it in the app architecture with panic and log the event with Log_Fatal
	// LogFatal(event LogEvent)

	RegisterListener(level LogType, domain string, ll LogListener) (err Error)
}

type LogEvent interface {
	Level() LogType
	Time() TimeUnixMilli
	Domain() string
	Message() string // save formated data e.g. fmt.Sprintf("Panic Exception: %s\nDebug Stack: %s", r, debug.Stack())
	Stack() []byte   // if log need to trace, specially in panic situation
}

// LogListener Usually implement on a distributed service that will carry log event to desire node and
// - Show on screen in control room of the software
// - Notify to related person about critical log that must check as soon as possible
// But it can also implement by local GUI application to notify the dev users in AppMode_Dev
type LogListener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	LogEventHandler(le LogEvent)
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
