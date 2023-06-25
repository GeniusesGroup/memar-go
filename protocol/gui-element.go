/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// GUI_Element
type GUI_Element interface {
	Name() string // It must be unique e.g. product
	Type() GUI_Element_Type
	Status() GUI_Element_Status

	HTML(outer bool) string
	Text() string

	// Create new element as copy of existing element, new element is a full (deep) copy of the element and
	// is disconnected initially from the DOM.
	Clone() GUI_Element
	CloneTo(dom GUI_Element)

	GetElementById(id string) GUI_Element
	GetElementsByName(name string) GUI_Element
	GetElementsByClassName(class string) []GUI_Element
	GetElementsByTagName(tag string) []GUI_Element
	GetElementsByTagNameNS(tag string) []GUI_Element

	HasFocus() bool

	Append(GUI_Element)
	Prepend(GUI_Element)

	Parent() GUI_Element
	NthChild(n int) GUI_Element
	// NextSibling()
	// PreviousSibling()
	// RemoveChild()
	// ReplaceChild()

	// CaptureEvents()
	// CreateEvent()
	// ReleaseEvents()

	EventTarget
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientHeight
	GUI_Scroll
}

type GUI_Element_Type uint8

// https://github.com/sciter-sdk/go-sciter/blob/master/types.go#L1363
// https://github.com/sciter-sdk/go-sciter/blob/master/sciter.go#L570
type GUI_Element_Status uint8

// https://en.wikipedia.org/wiki/HTML_attribute
// https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes
// https://developer.mozilla.org/en-US/docs/Web/CSS/Pseudo-classes
const (
	GUI_Element_Status_Unset GUI_Element_Status = iota
	GUI_Element_Status_Normal
	GUI_Element_Status_Error

	GUI_Element_Status_Opening
	GUI_Element_Status_Opened
	GUI_Element_Status_Closing
	GUI_Element_Status_Closed

	GUI_Element_Status_Toggling
	GUI_Element_Status_Pushed

	// User actions states
	GUI_Element_Status_Hover
	GUI_Element_Status_Active
	GUI_Element_Status_Focused
	GUI_Element_Status_FocusedVisible
	GUI_Element_Status_FocusedWithin

	// Resource states
	GUI_Element_Status_Playing
	GUI_Element_Status_Paused
	GUI_Element_Status_Waiting

	// Input states
	GUI_Element_Status_Autofilled
	GUI_Element_Status_Enabled
	GUI_Element_Status_Disabled
	GUI_Element_Status_ReadOnly
	GUI_Element_Status_WriteOnly
	GUI_Element_Status_Selected
	GUI_Element_Status_Checked
	GUI_Element_Status_Visited
	GUI_Element_Status_Hidden

	// Drag states
	GUI_Element_Status_Draggable
	GUI_Element_Status_DragEnter
	GUI_Element_Status_DragStart
	GUI_Element_Status_Dragging
	GUI_Element_Status_DragEnd
	GUI_Element_Status_DragExit
	GUI_Element_Status_Droped
	GUI_Element_Status_Leaved
	GUI_Element_Status_DragOver
)
