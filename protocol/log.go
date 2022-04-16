/* For license and copyright information please see LEGAL file in repository */

package protocol

// Logger provide logging mechanism to prevent application from runtime crashes and
// save details about runtime events to check by developers to fix bugs or develope better features.
type Logger interface {
	// PanicHandler recover from panics in a goroutine if exist, to prevent the application unexpected stopping.
	PanicHandler()

	// Log save logs as log event in time chain for the node that create log.
	// Suggest that RecordID time chain by sha3.256(LogMediatypeID, NodeID, TimeRoundToDay)
	Log(event LogEvent)
	// Due to expect Fatal terminate app and it brake the app, Dev must design it in the app architecture with panic and log the event with LogEvent_Fatal
	// LogFatal(event LogEvent)

	// Log listener mechanism usually implement on some kind of services that:
	// - Carry log event to desire node and show on screen e.g. in control room of the organization
	// - Notify to related person about critical log that must check as soon as possible by pager, sms, email, web notification, user GUI app, ...
	// - Local GUI application to notify the developers in AppMode_Dev
	EventTarget
}

type LogEvent interface {
	Event

	Level() LogType
	Time() Time
	Domain() string
	Message() string // save formated data e.g. fmt.Sprintf("Panic Exception: %s\nDebug Stack: %s", r, debug.Stack())
	Stack() []byte   // if log need to trace, specially in panic situation
}

// LogType indicate log level that will also use as EventSubType too.
type LogType uint8

const (
	LogEvent_Unset LogType = iota
	LogEvent_Information
	LogEvent_Notice
	LogEvent_Debug // Detailed information on the flow through the system. Expect these to be written to logs only. Generally speaking, most lines logged by your application should be written as DEBUG.
	LogEvent_DeepDebug
	LogEvent_Warning // Use of deprecated APIs, poor use of API, 'almost' errors, other runtime situations that are undesirable or unexpected, but not necessarily "wrong". Expect these to be immediately visible on a status console.
	LogEvent_Error   // Other runtime errors or unexpected conditions
	LogEvent_Alert
	LogEvent_Panic // in panic() it will add debug stack to trace more easily panic errors
	LogEvent_Critical
	LogEvent_Emergency
	LogEvent_Fatal // Severe errors that cause premature termination
	LogEvent_Security
	LogEvent_Confidential
)

// If any below mode disabled, logger must not save that log type.
const (
	LogMode_Debug        = true
	LogMode_DeepDebug    = true
	LogMode_Confidential = false
)
