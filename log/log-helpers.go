/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/protocol"
)

// TODO::: Can't force compiler to inline below functions, Delete file to force developers use this way:
// Logger.DispatchEvent(log.ConfEvent(&domain, "???"))

// Trace make new event with given level and add stack trace and log it to Logger
func Trace(domain protocol.MediaType, level protocol.LogLevel, message protocol.String) (err protocol.Error) {
	// var e Event
	// e.Init(level, message, nil)
	return Logger.DispatchEvent(TraceEvent(domain, level, message))
}

// Info make new event with "Information" level and log it to Logger
func Info(domain protocol.MediaType, message protocol.String) (err protocol.Error) {
	var e Event
	e.Init(domain, protocol.LogLevel_Information, message, nil)
	return Logger.DispatchEvent(&e)
}

// Notice make new event with "Notice" level and log it to Logger
func Notice(domain protocol.MediaType, message protocol.String) (err protocol.Error) {
	var e Event
	e.Init(domain, protocol.LogLevel_Notice, message, nil)
	return Logger.DispatchEvent(&e)
}

// Debug make new event with "Debug" level and log it to Logger
func Debug(domain protocol.MediaType, message protocol.String) (err protocol.Error) {
	var e Event
	e.Init(domain, protocol.LogLevel_Debug, message, nil)
	return Logger.DispatchEvent(&e)
}

// DeepDebug make new event with "DeepDebug" level and log it to Logger
func DeepDebug(domain protocol.MediaType, message protocol.String) (err protocol.Error) {
	var e Event
	e.Init(domain, protocol.LogLevel_DeepDebug, message, nil)
	e.DefaultStack()
	return Logger.DispatchEvent(&e)
}

// Warn make new event with "Warning" level and log it to Logger
func Warn(domain protocol.MediaType, message protocol.String) (err protocol.Error) {
	var e Event
	e.Init(domain, protocol.LogLevel_Warning, message, nil)
	return Logger.DispatchEvent(&e)
}

// Fatal make new event with "Fatal" level and log it to Logger
func Fatal(domain protocol.MediaType, message protocol.String) (err protocol.Error) {
	var e Event
	e.Init(domain, protocol.LogLevel_Fatal, message, nil)
	e.DefaultStack()
	return Logger.DispatchEvent(&e)
}

// Conf make new event with "Confidential" level and log it to Logger
func Conf(domain protocol.MediaType, message protocol.String) (err protocol.Error) {
	var e Event
	e.Init(domain, protocol.LogLevel_Confidential, message, nil)
	return Logger.DispatchEvent(&e)
}
