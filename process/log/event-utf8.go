/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"runtime/debug"

	"memar/codec/string/utf8"
	datatype_p "memar/computer/datatype/protocol"
	log_p "memar/process/log/protocol"
)

// Event is a log event
type Event_UTF8 struct {
	Event

	msgSTR   utf8.STR
	stackSTR utf8.STR
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (e *Event_UTF8) Init(dt datatype_p.DataType, level log_p.Level, message string, stack bool) {
	e.msgSTR.FromString_Go(message)
	if stack {
		e.stackSTR.Init(debug.Stack())
	}

	e.Event.Init(dt, level, &e.msgSTR, &e.stackSTR)
}
