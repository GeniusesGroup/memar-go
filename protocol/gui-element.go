/* For license and copyright information please see LEGAL file in repository */

package protocol

// GUIElement
type GUIElement interface {
	Name() string // It must be unique e.g. product
	Type() GUIElementType
	State() GUIElementState

	HTML(outer bool) string
	Text() string

	// Create new element as copy of existing element, new element is a full (deep) copy of the element and
	// is disconnected initially from the DOM.
	Clone() GUIElement
	CloneTo(dom GUIElement)

	GetElementById(id string) GUIElement
	GetElementsByName(name string) GUIElement
	GetElementsByClassName(class string) []GUIElement
	GetElementsByTagName(tag string) []GUIElement
	GetElementsByTagNameNS(tag string) []GUIElement

	HasFocus() bool

	Append(GUIElement)
	Prepend(GUIElement)

	Parent() GUIElement
	NthChild(n int) GUIElement
	// NextSibling()
	// PreviousSibling()
	// RemoveChild()
	// ReplaceChild()

	// CaptureEvents()
	// CreateEvent()
	// ReleaseEvents()

	EventTarget
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientHeight
	GUIScroll
}

type GUIElementType uint8

// https://github.com/sciter-sdk/go-sciter/blob/master/types.go#L1363
// https://github.com/sciter-sdk/go-sciter/blob/master/sciter.go#L570
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
