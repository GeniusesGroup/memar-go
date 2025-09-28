/* For license and copyright information please see the LEGAL file in the code repository */

package event_p

import (
	error_p "memar/error/protocol"
)

// Target is a interface implemented to receive events and may have listeners for them.
//
// It MUST implement for each data-type in each domain that need to dispatch,
// It means method only accept that data-type event not any other one.
//
// Other frameworks:
// https://developer.mozilla.org/en-US/docs/Web/API/Target
type Target[E Event] interface {
	// Appends an event listener for events whose type attribute value is type.
	// The callback argument sets the callback that will be invoked when the event is dispatched.
	// The event listener is appended to target's event listener list and is not appended if it has the same type, callback, and capture.
	//
	// Other frameworks:
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
	AddEventListener(callback EventListener[E]) (err error_p.Error)

	// Removes the event listener in target's event listener list with the same type, callback, and options.
	//
	// Other frameworks:
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
	RemoveEventListener(callback EventListener[E]) (err error_p.Error)

	// DispatchEvent() or Raise() invokes event handlers synchronously(immediately).
	// All applicable event handlers are called and return before DispatchEvent() returns.
	// The terms "notify clients", "send notifications", "trigger notifications", and "fire notifications" are used interchangeably with DispatchEvent.
	// Unlike web APIs, developers can check event.DefaultPrevented() after return, we don't return any data.
	//
	// Other frameworks:
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent
	//
	// Raise(event E) (err error_p.Error)
	DispatchEvent(event E) (err error_p.Error)

	// EventListeners() []EventListener[E] // Due to security problem, can't expose listeners to others
}

// EventListener Usually implement on some kind of services that:
// - Carry log event to desire node and show on screen e.g. in control room of the organization
// - Notify to related person about critical log that must check as soon as possible by pager, sms, email, web notification, user GUI app, ...
// - Local GUI application to notify the developers in `log.CNF_DevelopingMode`
type EventListener[E Event] interface {
	// It CAN be Blocking(Synchronous) & Non-Blocking, means It CAN block the caller in any ways.
	// **Strongly suggest just use blocking if you force by strong logic requirements like authorization**
	EventHandler(event E)

	//
	// EventListener can implement some functionality like:::
	//

	// It will part of process and can change the flow of the desire event.
	// Some consideration:
	// 	- Target can prevent in `AddEventListener` to register sync listener.
	// 	- May be `Event` dispatcher don't implement `Event.DefaultPrevented()` and not let you to change the flow.
	// 	- Just if `Event.Cancelable()`` is `true`, listener can change the flow.
	// Synchronous() bool

	// - AddEventListener: Capture indicating that events of this type will be dispatched to the registered listener
	// before being dispatched to any Target beneath it in the target tree.
	// - RemoveEventListener: Capture that specifies whether the EventListener to be removed is registered as a capturing listener or not.
	// Capture() bool

	// Once indicating that the listener should be invoked at most once after being added.
	// If true, the listener would be automatically removed when invoked.
	// Once() bool

	// Passive() is the same logic as Synchronous()
	// Passive indicates that the function specified by listener will never call preventDefault().
	// If a passive listener does call preventDefault(), the user agent will do nothing other than generate a warning log.
	// See Improving scrolling performance with passive listeners to learn more.
	// Passive() bool

	// The listener will be removed when receive true on AbortSignal channel.
	// It is not free lunch, so we decide to not support it. Developers can use RemoveEventListener() to remove any listener explicitly.
	// AbortSignal() chan bool
}
