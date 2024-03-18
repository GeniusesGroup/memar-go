/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Logger provide logging mechanism to dispatch events about runtime events
// to check by developers to fix bugs or develope better features.
//
// `Logger` and `LogEvent` are local data structures.
// Distributed log listener mechanism usually implement on some kind of services that:
// - Provide many filter for listens on events.
// - Carry log event to desire node and show on screen e.g. in control room of the organization
// - Notify to related person about critical log that must check as soon as possible by pager, sms, email, web notification, user GUI app, ...
// - Local GUI application to notify the developers in AppMode_Dev
// For distributed usage the related domain module MUST provide other one that include e.g. `AppNodeID`, ...
// Distributed log module can do these logic (not limit to these):
// - Dispatch events to their listeners.
// - Cache log events in the node that create it.
// - Save all logs per day for a node in the record with LogMediatypeID as record type and `AppNodeID` as primary key.
//
// Log or Logging package can provide some helper function to let developers log more easily.
// Log functions make related event and call DispatchEvent(event) to carry events to local listeners.
//
// Due to expect Fatal terminate app and it brake the app, Dev must design it in the app architecture with panic and log the event with LogLevel_Fatal
// LogFatal(event LogEvent)
//
// We can't accept all data in below post and proposal, just add to more details.
// https://docs.google.com/document/d/1nFRxQ5SJVPpIBWTFHV-q5lBYiwGrfCMkESFGNzsrvBU/
// https://dave.cheney.net/2015/11/05/lets-talk-about-logging
type Logger[LE LogEvent, OPTs any] EventTarget[LE, OPTs]

// LogEvent just suggest base structure, Additional data can structure in `Message` field as describe in many RFCs e.g.
// https://datatracker.ietf.org/doc/html/rfc5424
type LogEvent interface {
	Event

	LogLevel() LogLevel

	LogEvent_Message
	Runtime_Stack
}

// LogEvent_Message will implement by any capsule that CAN be a log event.
// Usually implement by `Error` capsules.
type LogEvent_Message interface {
	// Log don't provide or suggest methods like Printf(format string, v ...interface{}) to writes a formatted message,
	// That must use some runtime logic e.g. fmt.Sprintf("Panic Exception: %s\nDebug Stack: %s", r, debug.Stack()).
	// Strongly suggest prepare formatted data in compile time by implement below method that provide log message.
	LogMessage() String
}

// LogLevel indicate log level.
type LogLevel uint32

const (
	LogLevel_Unset LogLevel = 0

	// LogLevel_AllSet = LogLevel_Debug | LogLevel_DeepDebug | LogLevel_Warning | LogLevel_Error | LogLevel_Alert |
	// LogLevel_Critical | LogLevel_Emergency | LogLevel_Fatal | LogLevel_Security | LogLevel_Confidential
	LogLevel_AllSet LogLevel = ^(LogLevel_Unset)

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
