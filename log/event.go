/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"runtime/debug"

	"memar/ce/utf8"
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

	msgSTR   utf8.String
	stackSTR utf8.ByteSlice
}

//memar:impl memar/protocol.ObjectLifeCycle
func (e *Event) Init(domain protocol.MediaType, level protocol.LogLevel, message, stack protocol.String) {
	e.level = level
	e.message = message
	e.stack = stack
	e.Event.Init(domain, unix.Now())
}

//memar:impl memar/protocol.LogEvent
func (e *Event) Level() protocol.LogLevel { return e.level }
func (e *Event) Message() protocol.String { return e.message }
func (e *Event) Stack() protocol.String   { return e.stack }

func (e *Event) DefaultStack() {
	e.stackSTR.Init(debug.Stack())
	e.stack = &e.stackSTR
}

func (e *Event) MSG_string(msg string) {
	e.msgSTR = utf8.String(msg)
	e.message = &e.msgSTR
}
