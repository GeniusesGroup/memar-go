/* For license and copyright information please see the LEGAL file in the code repository */

package event_p

import (
	"memar/protocol"
)

// EventTarget is a interface implemented to receive events and may have listeners for them.
// It MUST implement in each domain that need to dispatch, It means method only accept that domain event not other one.
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget[E Event, OPTs any] interface {
	// Appends an event listener for events whose type attribute value is type.
	// The callback argument sets the callback that will be invoked when the event is dispatched.
	// The event listener is appended to target's event listener list and is not appended if it has the same type, callback, and capture.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
	AddEventListener(callback EventListener[E], options OPTs) (err protocol.Error)

	// Removes the event listener in target's event listener list with the same type, callback, and options.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
	RemoveEventListener(callback EventListener[E], options OPTs) (err protocol.Error)

	// DispatchEvent() or Raise() invokes event handlers synchronously(immediately).
	// All applicable event handlers are called and return before DispatchEvent() returns.
	// The terms "notify clients", "send notifications", "trigger notifications", and "fire notifications" are used interchangeably with DispatchEvent.
	// Unlike web APIs, developers can check event.DefaultPrevented() after return, we don't return any data.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent
	//
	// Raise[T Event](event T) (err Error)
	DispatchEvent(event E) (err protocol.Error)

	// EventListeners() []EventListener[E] // Due to security problem, can't expose listeners to others
}

// EventListener Usually implement on some kind of services that:
// - Carry log event to desire node and show on screen e.g. in control room of the organization
// - Notify to related person about critical log that must check as soon as possible by pager, sms, email, web notification, user GUI app, ...
// - Local GUI application to notify the developers in `log.CNF_DevelopingMode`
type EventListener[E Event] interface {
	// Non-Blocking, means It must not block the caller in any ways.
	EventHandler(event E)
}
