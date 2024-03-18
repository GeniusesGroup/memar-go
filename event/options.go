/* For license and copyright information please see the LEGAL file in the code repository */

package event

// If not specified, defaults to false.
type Options struct {
	// - AddEventListener: A boolean value indicating that events of this type will be dispatched to the registered listener
	// before being dispatched to any EventTarget beneath it in the DOM tree.
	// - RemoveEventListener: A boolean value that specifies whether the EventListener to be removed is registered as a capturing listener or not.
	Capture bool

	// A boolean value indicating that the listener should be invoked at most once after being added.
	// If true, the listener would be automatically removed when invoked.
	Once bool

	// A boolean value that, if true, indicates that the function specified by listener will never call preventDefault().
	// If a passive listener does call preventDefault(), the user agent will do nothing other than generate a warning log.
	// See Improving scrolling performance with passive listeners to learn more.
	Passive bool

	// The listener will be removed when receive true on AbortSignal channel.
	// It is not free lunch, so we decide to not support it. Developers can use RemoveEventListener() to remove any listener explicitly.
	// AbortSignal chan bool
}
