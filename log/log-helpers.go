/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"libgo/protocol"
)

// TODO::: Can't force compiler to inline below functions, Delete file to force developers use this way:
// Logger.Log(log.ConfEvent(domainEnglish, "???"))

// Trace make new event with given level and add stack trace and log it to Logger
func Trace(mediaType protocol.MediaType, level protocol.LogLevel, message string) (err protocol.Error) {
	// var e Event
	// e.Init(level, message, true)
	return Logger.Log(TraceEvent(mediaType, level, message))
}

// Info make new event with "Information" level and log it to Logger
func Info(mediaType protocol.MediaType, message string) (err protocol.Error) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Information, message, false)
	return Logger.Log(&e)
}

// Notice make new event with "Notice" level and log it to Logger
func Notice(mediaType protocol.MediaType, message string) (err protocol.Error) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Notice, message, false)
	return Logger.Log(&e)
}

// Debug make new event with "Debug" level and log it to Logger
func Debug(mediaType protocol.MediaType, message string) (err protocol.Error) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Debug, message, false)
	return Logger.Log(&e)
}

// DeepDebug make new event with "DeepDebug" level and log it to Logger
func DeepDebug(mediaType protocol.MediaType, message string) (err protocol.Error) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_DeepDebug, message, true)
	return Logger.Log(&e)
}

// Warn make new event with "Warning" level and log it to Logger
func Warn(mediaType protocol.MediaType, message string) (err protocol.Error) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Warning, message, false)
	return Logger.Log(&e)
}

// Fatal make new event with "Fatal" level and log it to Logger
func Fatal(mediaType protocol.MediaType, message string) (err protocol.Error) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Fatal, message, true)
	return Logger.Log(&e)
}

// Conf make new event with "Confidential" level and log it to Logger
func Conf(mediaType protocol.MediaType, message string) (err protocol.Error) {
	var e Event
	e.Init(mediaType, protocol.LogLevel_Confidential, message, false)
	return Logger.Log(&e)
}
