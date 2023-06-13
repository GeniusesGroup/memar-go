/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Logger provide logging mechanism to prevent application from runtime crashes and
// save details about runtime events to check by developers to fix bugs or develope better features.
// We can't accept all data in below post and proposal, just add to more details.
// https://docs.google.com/document/d/1nFRxQ5SJVPpIBWTFHV-q5lBYiwGrfCMkESFGNzsrvBU/
// https://dave.cheney.net/2015/11/05/lets-talk-about-logging
type Logger interface {
	// TODO::: Log or Logging:
	// Log suggest to:
	// - First Dispatch(event).
	// - Cache log events in the node that create it.
	// - Save all logs per day for a node in the record with LogMediatypeID as record type and NodeID as primary key.
	Log(event LogEvent) (err Error)

	// Due to expect Fatal terminate app and it brake the app, Dev must design it in the app architecture with panic and log the event with LogLevel_Fatal
	// LogFatal(event LogEvent)

	// Log listener mechanism usually implement on some kind of services that:
	// - Carry log event to desire node and show on screen e.g. in control room of the organization
	// - Notify to related person about critical log that must check as soon as possible by pager, sms, email, web notification, user GUI app, ...
	// - Local GUI application to notify the developers in AppMode_Dev
	EventTarget
}

// LogEvent just suggest base structure, Additional data can structure in `Message` field as describe in many RFCs e.g.
// https://datatracker.ietf.org/doc/html/rfc5424
type LogEvent interface {
	Event

	Level() LogLevel //
	Message() string // save formatted data e.g. fmt.Sprintf("Panic Exception: %s\nDebug Stack: %s", r, debug.Stack())
	Stack() []byte   // if log need to trace, specially in panic situation. Default fill by `debug.Stack()`

	Codec
}

// LogLevel indicate log level.
type LogLevel uint32

const (
	LogLevel_Unset LogLevel = 0

	//
	LogLevel_Information LogLevel = (1 << iota)

	_

	// for normal, but significant messages
	LogLevel_Notice

	_

	// Any runtime errors or unexpected conditions that haven't any level.
	// Usually in old library that implement errors in text only structure live in go std modules.
	LogLevel_Error

	_

	// Detailed information on the flow through the system.
	// Expect these to be written to logs only. Generally speaking, most lines logged by your application should be written as DEBUG.
	LogLevel_Debug

	_

	LogLevel_DeepDebug

	_

	// Use of deprecated APIs, poor use of API, 'almost' errors, other runtime situations that are undesirable or unexpected,
	// but not necessarily "wrong". Expect these to be immediately visible on a status console.
	LogLevel_Warning

	_

	// for alerts, actions that must be taken immediately, ex. corrupted database
	LogLevel_Alert

	_

	// for critical conditions e.g. component unavailable, unexpected exception, ...
	LogLevel_Critical

	_

	// when system is unusable, panics
	LogLevel_Emergency

	_

	// Severe errors that cause premature termination
	LogLevel_Fatal

	_

	LogLevel_Security

	_

	// It can be enabled along with any above level indicate log carry sensitive data like full http data.
	LogLevel_Confidential
)

// If any below mode disabled, logger must not save that log type.
// But logger must Dispatch() it to any client requested any types even it is not enabled.
const (
	LogMode LogLevel = LogLevel_Debug | LogLevel_DeepDebug | LogLevel_Warning | LogLevel_Error | LogLevel_Alert |
		LogLevel_Critical | LogLevel_Emergency | LogLevel_Fatal | LogLevel_Security | LogLevel_Confidential
)
