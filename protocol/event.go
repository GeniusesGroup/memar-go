/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Event usually can be any other domain records that store in storage layer.
// https://www.w3.org/TR/DOM-Level-3-Events/#event-flow
// https://developer.mozilla.org/en-US/docs/Web/API/Event
// https://developer.mozilla.org/en-US/docs/Web/Events
type Event interface {
	Domain() DataType
	Time() Time

	// Returns true or false depending on how event was initialized. Its return value does not always carry meaning,
	// but true can indicate that part of the operation during which event was dispatched, can be canceled by invoking the preventDefault() method.
	// It also means subscribers receive events in asynchronous or synchronous manner. true means synchronous manner.
	Cancelable() bool
	// Returns true if preventDefault() was invoked successfully to indicate cancellation, and false otherwise.
	DefaultPrevented() bool

	Event_Methods
}

type Event_Methods interface {
	// If invoked when the cancelable attribute value is true, and while executing a listener for the event with passive set to false,
	// signals to the operation that caused event to be dispatched that it needs to be canceled.
	PreventDefault()
}

// EventTarget is a interface implemented to receive events and may have listeners for them.
// It MUST implement in each domain that need to dispatch, It means method only accept that domain event not other one.
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget[E Event, OPTs any] interface {
	// Appends an event listener for events whose type attribute value is type.
	// The callback argument sets the callback that will be invoked when the event is dispatched.
	// The event listener is appended to target's event listener list and is not appended if it has the same type, callback, and capture.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
	AddEventListener(callback EventListener[E], options OPTs) (err Error)

	// Removes the event listener in target's event listener list with the same type, callback, and options.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
	RemoveEventListener(callback EventListener[E], options OPTs) (err Error)

	// DispatchEvent() or Raise() invokes event handlers synchronously(immediately).
	// All applicable event handlers are called and return before DispatchEvent() returns.
	// The terms "notify clients", "send notifications", "trigger notifications", and "fire notifications" are used interchangeably with DispatchEvent.
	// Unlike web APIs, developers can check event.DefaultPrevented() after return, we don't return any data.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent
	//
	// Raise[T Event](event T) (err Error)
	DispatchEvent(event E) (err Error)

	// EventListeners() []EventListener[E] // Due to security problem, can't expose listeners to others
}

// EventListener Usually implement on some kind of services that:
// - Carry log event to desire node and show on screen e.g. in control room of the organization
// - Notify to related person about critical log that must check as soon as possible by pager, sms, email, web notification, user GUI app, ...
// - Local GUI application to notify the developers in AppMode_Dev
type EventListener[E Event] interface {
	// Non-Blocking, means It must not block the caller in any ways.
	EventHandler(event E)
}
