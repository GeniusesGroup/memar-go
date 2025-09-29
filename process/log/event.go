/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	log_p "memar/audit/log/protocol"
	datatype_p "memar/datatype/protocol"
	"memar/event"
	string_p "memar/string/protocol"
	"memar/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	event.Event

	level   log_p.Level
	message string_p.String
	stack   string_p.String
}

//memar:impl memar/computer/language/object/protocol.LifeCycle
func (e *Event) Init(dt datatype_p.DataType, level log_p.Level, message, stack string_p.String) {
	e.level = level
	e.message = message
	e.stack = stack
	e.Event.Init(dt, unix.Now())
}

//memar:impl memar/protocol.LogEvent
func (e *Event) LogLevel() log_p.Level { return e.level }

//memar:impl memar/protocol.LogEvent_Message
func (e *Event) LogMessage() string_p.String { return e.message }

//memar:impl memar/protocol.Runtime_Stack
func (e *Event) RuntimeStack() string_p.String { return e.stack }
