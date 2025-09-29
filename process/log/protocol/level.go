/* For license and copyright information please see the LEGAL file in the code repository */

package log_p

// Event_Level will implement by any capsule that CAN be a log event.
// Usually implement by `log event` capsules.
type Event_Level interface {
	LogLevel() Level
}

// Level indicate log level.
type Level uint32

const (
	Level_Unset Level = 0

	// Level_AllSet = Level_Debug | Level_DeepDebug | Level_Warning | Level_Error | Level_Alert |
	// Level_Critical | Level_Emergency | Level_Fatal | Level_Security | Level_Confidential
	Level_AllSet Level = ^(Level_Unset)

	//
	Level_Information Level = (1 << iota)

	_

	// for normal, but significant messages
	Level_Notice

	_

	// Any runtime errors or unexpected conditions that haven't any level.
	// Usually in old library that implement errors in text only structure live in go std modules.
	Level_Error

	_

	// Detailed information on the flow through the system.
	// Expect these to be written to logs only. Generally speaking, most lines logged by your application should be written as DEBUG.
	Level_Debug

	_

	Level_DeepDebug

	_

	// Use of deprecated APIs, poor use of API, 'almost' errors, other runtime situations that are undesirable or unexpected,
	// but not necessarily "wrong". Expect these to be immediately visible on a status console.
	Level_Warning

	_

	// for alerts, actions that must be taken immediately, ex. corrupted database
	Level_Alert

	_

	// for critical conditions e.g. component unavailable, unexpected exception, ...
	Level_Critical

	_

	// when system is unusable, panics
	Level_Emergency

	_

	// Severe errors that cause premature termination
	Level_Fatal

	_

	Level_Security

	_

	// It can be enabled along with any above level indicate log carry sensitive data like full http data.
	Level_Confidential
)
