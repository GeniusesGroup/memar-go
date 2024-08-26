/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

// Event usually can be any other domain records that store in storage layer.
// https://www.w3.org/TR/DOM-Level-3-Events/#event-flow
// https://developer.mozilla.org/en-US/docs/Web/API/Event
// https://developer.mozilla.org/en-US/docs/Web/Events
type Event interface {
	// Returns true or false depending on how event was initialized.
	// True if event goes through its target's ancestors in reverse tree order, and false otherwise.
	// When set to true, options's capture prevents callback from being invoked when the event's eventPhase attribute value is BUBBLING_PHASE.
	// When false (or not present), callback will not be invoked when event's eventPhase attribute value is CAPTURING_PHASE.
	// Either way, callback will be invoked if event's eventPhase attribute value is AT_TARGET.
	// When set to true, options's passive indicates that the callback will not cancel the event by invoking preventDefault().
	// This is used to enable performance optimizations described in § 2.8 Observing event listeners.
	// When set to true, options's once indicates that the callback will only be invoked once after which the event listener will be removed.
	Bubbles() bool
	// Indicates whether or not the event was initiated by the browser (after a user click, for instance) or by a script (using an event creation method, for example).
	// IsTrusted() bool
}
