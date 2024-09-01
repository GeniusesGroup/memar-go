/* For license and copyright information please see the LEGAL file in the code repository */

package log_p

import (
	event_p "memar/event/protocol"
)

// Logger provide logging mechanism to dispatch events about runtime events
// to check by developers to fix bugs or develope better features.
//
// `Logger` and `Event` are local data structures.
// Distributed log listener mechanism usually implement on some kind of services that:
// - Provide many filter for listens on events.
// - Carry log event to desire node and show on screen e.g. in control room of the organization
// - Notify to related person about critical log that must check as soon as possible by pager, sms, email, web notification, user GUI app, ...
// - Local GUI application to notify the developers in `log.CNF_DevelopingMode`
// For distributed usage the related domain module MUST provide other one that include e.g. `AppNodeID`, ...
// Distributed log module can do these logic (not limit to these):
// - Dispatch events to their listeners.
// - Cache log events in the node that create it.
// - Save all logs per day for a node in the record with LogMediatypeID as record type and `AppNodeID` as primary key.
//
// Log or Logging package can provide some helper function to let developers log more easily.
// Log functions make related event and call DispatchEvent(event) to carry events to local listeners.
//
// Due to expect Fatal terminate app and it brake the app, Dev must design it in the app architecture with panic and log the event with LogLevel_Fatal
// LogFatal(le Event)
//
// We can't accept all data in below post and proposal, just add to more details.
// https://docs.google.com/document/d/1nFRxQ5SJVPpIBWTFHV-q5lBYiwGrfCMkESFGNzsrvBU/
// https://dave.cheney.net/2015/11/05/lets-talk-about-logging
type Logger[LE Event, OPTs any] event_p.EventTarget[LE, OPTs]
