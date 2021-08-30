/* For license and copyright information please see LEGAL file in repository */

package protocol

// GUIElement
type GUIElement interface {
	Name() string // It must be unique e.g. product
	Type() int
	State() GUIElementState
}

type GUIElementState uint8

// https://en.wikipedia.org/wiki/HTML_attribute
// https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes
// https://developer.mozilla.org/en-US/docs/Web/CSS/Pseudo-classes
const (
	ElementStateUnset GUIElementState = iota
	ElementStateNormal
	ElementStateError

	ElementStateOpening
	ElementStateOpened
	ElementStateClosing
	ElementStateClosed

	ElementStateToggling
	ElementStatePushed

	// User actions states
	ElementStateHover
	ElementStateActive
	ElementStateFocused
	ElementStateFocusedVisible
	ElementStateFocusedWithin

	// Resource states
	ElementStatePlaying
	ElementStatePaused
	ElementStateWaiting

	// Input states
	ElementStateAutofilled
	ElementStateEnabled
	ElementStateDisabled
	ElementStateReadOnly
	ElementStateWriteOnly
	ElementStateSelected
	ElementStateChecked
	ElementStateVisited
	ElementStateHidden

	// Drag states
	ElementStateDraggable
	ElementStateDragEnter
	ElementStateDragStart
	ElementStateDragging
	ElementStateDragEnd
	ElementStateDragExit
	ElementStateDroped
	ElementStateLeaved
	ElementStateDragOver
)
