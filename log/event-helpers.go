/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/protocol"
)

func NewEvent(dt protocol.DataType, level protocol.LogLevel, message, stack protocol.String) (event *Event) {
	//go:gc stack
	var e Event
	e.Init(dt, level, message, stack)
	return &e
}

func NoticeEvent(dt protocol.DataType, message, stack protocol.String) (event *Event) {
	var e Event
	e.Init(dt, protocol.LogLevel_Notice, message, stack)
	return &e
}

func ErrorEvent(dt protocol.DataType, message, stack protocol.String) (event *Event) {
	var e Event
	e.Init(dt, protocol.LogLevel_Error, message, stack)
	return &e
}

func DebugEvent(dt protocol.DataType, message, stack protocol.String) (event *Event) {
	var e Event
	e.Init(dt, protocol.LogLevel_Debug, message, stack)
	return &e
}

func DeepDebugEvent(dt protocol.DataType, message, stack protocol.String) (event *Event) {
	var e Event
	e.Init(dt, protocol.LogLevel_DeepDebug, message, stack)
	return &e
}

func InfoEvent(dt protocol.DataType, message, stack protocol.String) (event *Event) {
	var e Event
	e.Init(dt, protocol.LogLevel_Information, message, stack)
	return &e
}

func WarnEvent(dt protocol.DataType, message, stack protocol.String) (event *Event) {
	var e Event
	e.Init(dt, protocol.LogLevel_Warning, message, stack)
	return &e
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(dt protocol.DataType, message, stack protocol.String) (event *Event) {
	var e Event
	e.Init(dt, protocol.LogLevel_Fatal, message, stack)
	return &e
}

// ConfEvent return new event with "Confidential" level
func ConfEvent(dt protocol.DataType, message, stack protocol.String) (event *Event) {
	var e Event
	e.Init(dt, protocol.LogLevel_Confidential, message, stack)
	return &e
}
