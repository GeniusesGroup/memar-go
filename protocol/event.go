/* For license and copyright information please see LEGAL file in repository */

package protocol

// https://www.w3.org/TR/DOM-Level-3-Events/#event-flow
// https://developer.mozilla.org/en-US/docs/Web/API/Event
// https://developer.mozilla.org/en-US/docs/Web/Events
type Event interface {
	MainType() EventMainType
	SubType() EventSubType
	Domain() string // (codes) package name
	NodeID() NodeID
	Time() Time
	// Returns true or false depending on how event was initialized. Its return value does not always carry meaning,
	// but true can indicate that part of the operation during which event was dispatched, can be canceled by invoking the preventDefault() method.
	// It also means subscribers receive events in asynchronous or synchronous manner. true means synchronous manner.
	Cancelable() bool
	// Returns true if preventDefault() was invoked successfully to indicate cancelation, and false otherwise.
	DefaultPrevented() bool
	// Returns true or false depending on how event was initialized.
	// True if event goes through its target's ancestors in reverse tree order, and false otherwise.
	// When set to true, options's capture prevents callback from being invoked when the event's eventPhase attribute value is BUBBLING_PHASE.
	// When false (or not present), callback will not be invoked when event's eventPhase attribute value is CAPTURING_PHASE.
	// Either way, callback will be invoked if event's eventPhase attribute value is AT_TARGET.
	// When set to true, options's passive indicates that the callback will not cancel the event by invoking preventDefault().
	// This is used to enable performance optimizations described in § 2.8 Observing event listeners.
	// When set to true, options's once indicates that the callback will only be invoked once after which the event listener will be removed.
	Bubbles() bool

	// If invoked when the cancelable attribute value is true, and while executing a listener for the event with passive set to false,
	// signals to the operation that caused event to be dispatched that it needs to be canceled.
	PreventDefault()
}

// EventTarget is a interface implemented to receive events and may have listeners for them.
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget interface {
	// Appends an event listener for events whose type attribute value is type.
	// The callback argument sets the callback that will be invoked when the event is dispatched.
	// The event listener is appended to target's event listener list and is not appended if it has the same type, callback, and capture.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
	AddEventListener(mainType EventMainType, subType EventSubType, callback EventListener, options AddEventListenerOptions)

	// Removes the event listener in target's event listener list with the same type, callback, and options.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
	RemoveEventListener(mainType EventMainType, subType EventSubType, callback EventListener, options EventListenerOptions)

	// DispatchEvent() invokes event handlers synchronously. All applicable event handlers are called and return before DispatchEvent() returns.
	// The terms "notify clients", "send notifications", "trigger notifications", and "fire notifications" are used interchangeably with DispatchEvent.
	// Unlike web APIs, developers can check event.DefaultPrevented() after return, we don't return any data.
	// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent
	// Raise domain events immediately when it is called
	// type DomainEventTarget interface {
	// Raise[T Event](event T)
	// }
	DispatchEvent(event Event)

	// EventListeners() []EventListener // Due to security problem, can't expose listeners to others
}

type EventListenerOptions struct {
	// - AddEventListener: A boolean value indicating that events of this type will be dispatched to the registered listener
	// before being dispatched to any EventTarget beneath it in the DOM tree.
	// - RemoveEventListener: A boolean value that specifies whether the EventListener to be removed is registered as a capturing listener or not.
	// If this parameter is absent, a default value of false is assumed.
	Capture bool
}

type AddEventListenerOptions struct {
	EventListenerOptions

	// A boolean value indicating that the listener should be invoked at most once after being added.
	// If true, the listener would be automatically removed when invoked. If not specified, defaults to false.
	Once bool

	// A boolean value that, if true, indicates that the function specified by listener will never call preventDefault().
	// If a passive listener does call preventDefault(), the user agent will do nothing other than generate a warning log.
	// If this parameter is absent, a default value of false is assumed.
	// See Improving scrolling performance with passive listeners to learn more.
	Passive bool

	// The listener will be removed when receive true on AbortSignal channel.
	// It is not free lunch, so we decide to not support it. Developers can use RemoveEventListener() to remove any listener explicitly.
	// AbortSignal chan bool
}

// EventListener Usually implement on some kind of services that:
// - Carry log event to desire node and show on screen e.g. in control room of the organization
// - Notify to related person about critical log that must check as soon as possible by pager, sms, email, web notification, user GUI app, ...
// - Local GUI application to notify the developers in AppMode_Dev
type EventListener interface {
	// Non-Blocking, means It must not block the caller in any ways.
	EventHandler(event Event)
}

// EventMainType indicate main type of events
type EventMainType = MediaTypeID

const (
	EventMainType_Unset  EventMainType = iota
	EventMainType_Custom               // specific to an application not app engine. like pages or widgets events.

	EventMainType_Log
	EventMainType_System // Kernel, OS GUI app, Other apps,
	EventMainType_Event  // notify when an event listener register. It can cause some security problems.
	EventMainType_Navigation

	EventMainType_Storage_Record
	EventMainType_Text

	// https://developer.mozilla.org/en-US/docs/Web/Events
	EventMainType_Animation
	EventMainType_AudioProcessing
	EventMainType_BeforeUnload
	EventMainType_Blob
	EventMainType_Clipboard
	EventMainType_Close
	EventMainType_Composition
	// EventMainType_Custom
	EventMainType_DeviceMotion
	EventMainType_DeviceOrientation
	EventMainType_DeviceProximity
	EventMainType_Drag
	EventMainType_Error
	EventMainType_Fetch
	EventMainType_Focus
	EventMainType_FormData
	EventMainType_Gamepad
	EventMainType_HashChange
	EventMainType_HIDInputReport
	EventMainType_IDBVersionChange
	EventMainType_Input
	EventMainType_Keyboard
	EventMainType_MediaStream
	EventMainType_Message
	EventMainType_Mouse
	EventMainType_Mutation
	EventMainType_OfflineAudioCompletion
	EventMainType_PageTransition
	EventMainType_PaymentRequestUpdate
	EventMainType_Pointer
	EventMainType_PopState
	EventMainType_Progress
	EventMainType_RTCDataChannel
	EventMainType_RTCPeerConnectionIce
	EventMainType_Storage
	EventMainType_Submit
	EventMainType_SVG
	EventMainType_Time
	EventMainType_Touch
	EventMainType_Track
	EventMainType_Transition
	EventMainType_UI
	EventMainType_UserProximity
	EventMainType_WebGLContext
	EventMainType_Wheel
)

// EventSubType indicate sub type of events
type EventSubType uint64

const (
	EventSubType_Unset EventSubType = 0
	// other sub types indicate in each main type file e.g. ./log.go
)
