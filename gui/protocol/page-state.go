/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

import (
	error_p "memar/error/protocol"
	picture_p "memar/picture/protocol"
	time_p "memar/time/protocol"
)

// PageState is like window in browsers and each page state has its own window
type PageState interface {
	PageState_Fields
	PageState_Methods
}

type PageState_Fields interface {
	Page() Page
	Title() string
	Description() string

	URL() string // Path + Parameters + Fragments
	Conditions() map[string]any
	Fragment() string

	// DOM as Document-Object-Model will render to user screen
	// Same as Document interface of the Web declare in browsers in many ways but split not related concern to other interfaces
	// like Document.URI that not belong to this interface and relate to PageState because we develope multi page app not web page
	DOM() DOM // Element
	SOM() SOM
	Thumbnail() picture_p.Image // The picture of the page in current state, use in social linking and SwitcherPage()

	ActiveDates() []time_p.Time // first is CreateDate and last is EndDate(When Close() called)
	State() ScreenMode
	Type() WindowType
}

type PageState_Methods interface {
	// Activate() or Show() Called call to render page in this state (brings to front).
	// Also can call to move the page state to other screen
	Activate(options PageState_ActivateOptions) (err error_p.Error)

	// Deactivate() or Minimize() Called before this state wants to remove from the render tree (brings to back)
	// Errors:
	// - NotApproveToLeave: let the caller know user of the GUI app let page in this state bring to back.
	// - HadActiveDialog: or hadActiveOverlay help navigator works in better way.
	// 		e.g. for some keyboard event like back button in android OS to close dialog not pop previous page state to front
	Deactivate() (err error_p.Error)

	// force to refresh
	Refresh() (err error_p.Error)

	SafeToSilentClose() bool
	// Remove render tree from screen and close DOM and SOM and archive it.
	// But do clean logic after some time e.g. 10sec maybe user close by mistake click action
	Close() (err error_p.Error)

	// DynamicElements struct {}
}

type PageState_ActivateOptions struct {
	Screen Screen
	// It will effect just on this app and this page state only, not OS level.
	Orientation ScreenOrientation
	WindowType  WindowType
	PopUp       bool // force to bring to front immediately
}

type WindowType uint8

const (
	WindowType_Unset      WindowType = 0
	WindowType_Resizeable WindowType = (1 << iota) // has resizeable frame
	WindowType_GLASSY                              // glassy window
	WindowType_ALPHA                               // transparent window
)
