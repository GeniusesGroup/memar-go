/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/event"
	"memar/protocol"
	"memar/time/unix"
)

// Event implement protocol.LogEvent
type Event struct {
	event.Event

	level   protocol.LogLevel
	message protocol.String
	stack   protocol.String
}

//memar:impl memar/protocol.ObjectLifeCycle
func (e *Event) Init(dt protocol.DataType, level protocol.LogLevel, message, stack protocol.String) {
	e.level = level
	e.message = message
	e.stack = stack
	e.Event.Init(dt, unix.Now())
}

//memar:impl memar/protocol.LogEvent
func (e *Event) LogLevel() protocol.LogLevel { return e.level }

//memar:impl memar/protocol.LogEvent_Message
func (e *Event) LogMessage() protocol.String { return e.message }

//memar:impl memar/protocol.Runtime_Stack
func (e *Event) RuntimeStack() protocol.String { return e.stack }
