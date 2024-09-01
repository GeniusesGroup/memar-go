/* For license and copyright information please see the LEGAL file in the code repository */

package log_p

import (
	event_p "memar/event/protocol"
	runtime_p "memar/runtime/protocol"
	string_p "memar/string/protocol"
)

// Event just suggest base structure, Additional data can structure in `Message` field as describe in many RFCs e.g.
// https://datatracker.ietf.org/doc/html/rfc5424
type Event interface {
	event_p.Event
	Event_Level
	Event_Message
	runtime_p.Stack
}

// Event_Message will implement by any capsule that CAN be a log event.
// Usually implement by `Error` capsules.
type Event_Message interface {
	// Log don't provide or suggest methods like Printf(format string, v ...interface{}) to writes a formatted message,
	// That must use some runtime logic e.g. fmt.Sprintf("Panic Exception: %s\nDebug Stack: %s", r, debug.Stack()).
	// Strongly suggest prepare formatted data in compile time by implement below method that provide log message.
	LogMessage() string_p.String
}
