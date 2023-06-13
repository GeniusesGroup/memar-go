/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"libgo/protocol"
)

func NewEvent(mediaType protocol.MediaType, level protocol.LogLevel, message string) (event *Event) {
	//go:gc stack
	var e Event
	e.Init(mediaType, level, message, false)
	return &e
}

func TraceEvent(mediaType protocol.MediaType, level protocol.LogLevel, message string) (event *Event) {
	var e Event
	e.Init(mediaType, level, message, true)
	return &e
}

func NoticeEvent(mediaType protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Notice, message, false)
	return &e
}

func ErrorEvent(mediaType protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Error, message, true)
	return &e
}

func DebugEvent(mediaType protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Debug, message, false)
	return &e
}

func DeepDebugEvent(mediaType protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_DeepDebug, message, true)
	return &e
}

func InfoEvent(mediaType protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Information, message, false)
	return &e
}

func WarnEvent(mediaType protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Warning, message, false)
	return &e
}

// FatalEvent return new event with fatal level and added stack trace.
func FatalEvent(mediaType protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Fatal, message, true)
	return &e
}

// ConfEvent return new event with "Confidential" level
func ConfEvent(mediaType protocol.MediaType, message string) (event *Event) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Confidential, message, false)
	return &e
}
