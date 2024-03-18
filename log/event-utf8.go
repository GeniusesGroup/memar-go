/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"runtime/debug"

	"memar/ce/utf8"
	"memar/protocol"
)

// Event implement protocol.LogEvent
type Event_UTF8 struct {
	Event

	msgSTR   utf8.String
	stackSTR utf8.ByteSlice
}

//memar:impl memar/protocol.ObjectLifeCycle
func (e *Event_UTF8) Init(dt protocol.DataType, level protocol.LogLevel, message string, stack bool) {
	e.msgSTR = utf8.String(message)
	if stack {
		e.stackSTR.Init(debug.Stack())
	}

	e.Event.Init(dt, level, &e.msgSTR, &e.stackSTR)
}
