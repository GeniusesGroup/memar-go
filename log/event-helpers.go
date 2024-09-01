/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	log_p "memar/log/protocol"
	"memar/protocol"
	string_p "memar/string/protocol"
)

func NewEvent(dt protocol.DataType, level log_p.Level, message, stack string_p.String) (event *Event) {
	//go:gc stack
	var e Event
	e.Init(dt, level, message, stack)
	return &e
}

func NoticeEvent(dt protocol.DataType, message, stack string_p.String) (event *Event) {
	var e Event
	e.Init(dt, log_p.Level_Notice, message, stack)
	return &e
}

func ErrorEvent(dt protocol.DataType, message, stack string_p.String) (event *Event) {
	var e Event
	e.Init(dt, log_p.Level_Error, message, stack)
	return &e
}

func DebugEvent(dt protocol.DataType, message, stack string_p.String) (event *Event) {
	var e Event
	e.Init(dt, log_p.Level_Debug, message, stack)
	return &e
}

func DeepDebugEvent(dt protocol.DataType, message, stack string_p.String) (event *Event) {
	var e Event
	e.Init(dt, log_p.Level_DeepDebug, message, stack)
	return &e
}

func InfoEvent(dt protocol.DataType, message, stack string_p.String) (event *Event) {
	var e Event
	e.Init(dt, log_p.Level_Information, message, stack)
	return &e
}

func WarnEvent(dt protocol.DataType, message, stack string_p.String) (event *Event) {
	var e Event
	e.Init(dt, log_p.Level_Warning, message, stack)
	return &e
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(dt protocol.DataType, message, stack string_p.String) (event *Event) {
	var e Event
	e.Init(dt, log_p.Level_Fatal, message, stack)
	return &e
}

// ConfEvent return new event with "Confidential" level
func ConfEvent(dt protocol.DataType, message, stack string_p.String) (event *Event) {
	var e Event
	e.Init(dt, log_p.Level_Confidential, message, stack)
	return &e
}
