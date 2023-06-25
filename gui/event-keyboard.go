/* For license and copyright information please see the LEGAL file in the code repository */

package gui

// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent
type KeyboardEvent interface {
	// Event

	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent#events
	Type() KeyboardEventType
	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values
	Key() KeyboardEventKey
	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/code
	Code() KeyboardEventCode
	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/repeat
	Repeat() bool
	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/altKey
	AltKey() bool
	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/ctrlKey
	CtrlKey() bool
	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/shiftKey
	ShiftKey() bool
	// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/metaKey
	MetaKey() bool
}

type KeyboardEventType uint8

const (
	KeyboardEventType_Unset KeyboardEventType = iota
	// https://developer.mozilla.org/en-US/docs/Web/API/Document/keydown_event
	KeyboardEventType_Key_Down
	// https://developer.mozilla.org/en-US/docs/Web/API/Document/keyup_event
	KeyboardEventType_Key_Up
	// https://developer.mozilla.org/en-US/docs/Web/API/Document/keypress_event
	// KeyboardEventType_Key_Press **Deprecated**
)

type KeyboardEventKey uint16

const ()

type KeyboardEventCode uint16

const ()
