/* For license and copyright information please see LEGAL file in repository */

package protocol

// https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent
type TouchEvent interface {
	Event

	Type() string
}

// https://developer.mozilla.org/en-US/docs/Web/API/Touch
