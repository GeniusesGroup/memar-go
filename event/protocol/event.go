/* For license and copyright information please see the LEGAL file in the code repository */

package event_p

import (
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
	time_p "memar/time/protocol"
)

// Event usually can be any other domain records that store in storage layer.
// https://www.w3.org/TR/DOM-Level-3-Events/#event-flow
// https://developer.mozilla.org/en-US/docs/Web/API/Event
// https://developer.mozilla.org/en-US/docs/Web/Events
type Event interface {
	Domain() datatype_p.DataType
	Time() time_p.Time

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
	PreventDefault() (err error_p.Error)
}
