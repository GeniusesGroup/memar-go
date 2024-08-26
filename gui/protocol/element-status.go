/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

// https://github.com/sciter-sdk/go-sciter/blob/master/types.go#L1363
// https://github.com/sciter-sdk/go-sciter/blob/master/sciter.go#L570
type Element_Status uint8

// https://en.wikipedia.org/wiki/HTML_attribute
// https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes
// https://developer.mozilla.org/en-US/docs/Web/CSS/Pseudo-classes
const (
	Element_Status_Unset Element_Status = iota
	Element_Status_Normal
	Element_Status_Error

	Element_Status_Opening
	Element_Status_Opened
	Element_Status_Closing
	Element_Status_Closed

	Element_Status_Toggling
	Element_Status_Pushed

	// User actions states
	Element_Status_Hover
	Element_Status_Active
	Element_Status_Focused
	Element_Status_FocusedVisible
	Element_Status_FocusedWithin

	// Resource states
	Element_Status_Playing
	Element_Status_Paused
	Element_Status_Waiting

	// Input states
	Element_Status_AutoFilled
	Element_Status_Enabled
	Element_Status_Disabled
	Element_Status_ReadOnly
	Element_Status_WriteOnly
	Element_Status_Selected
	Element_Status_Checked
	Element_Status_Visited
	Element_Status_Hidden

	// Drag states
	Element_Status_Draggable
	Element_Status_DragEnter
	Element_Status_DragStart
	Element_Status_Dragging
	Element_Status_DragEnd
	Element_Status_DragExit
	Element_Status_Droped
	Element_Status_Leaved
	Element_Status_DragOver
)
