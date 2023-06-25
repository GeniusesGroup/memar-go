/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// GUI_PageState is like window in browsers and each page state has its own window
type GUI_PageState interface {
	GUI_PageState_Fields
	GUI_PageState_Methods
}

type GUI_PageState_Fields interface {
	Page() GUI_Page
	Title() string
	Description() string

	URL() string // Path + Parameters + Fragments
	Conditions() map[string]any
	Fragment() string

	// DOM as Document-Object-Model will render to user screen
	// Same as Document interface of the Web declare in browsers in many ways but split not related concern to other interfaces
	// like Document.URI that not belong to this interface and relate to GUI_PageState because we develope multi page app not web page
	DOM() DOM // GUI_Element
	SOM() SOM
	Thumbnail() Image // The picture of the page in current state, use in social linking and SwitcherPage()

	ActiveDates() []Time // first is CreateDate and last is EndDate(When Close() called)
	State() ScreenMode
	Type() GUI_WindowType
}

type GUI_PageState_Methods interface {
	// Activate() or Show() Called call to render page in this state (brings to front).
	// Also can call to move the page state to other screen
	Activate(options GUI_PageState_ActivateOptions) (err Error)

	// Deactivate() or Minimize() Called before this state wants to remove from the render tree (brings to back)
	// Errors:
	// - NotApproveToLeave: let the caller know user of the GUI app let page in this state bring to back.
	// - HadActiveDialog: or hadActiveOverlay help navigator works in better way.
	// 		e.g. for some keyboard event like back button in android OS to close dialog not pop previous page state to front
	Deactivate() (err Error)

	// force to refresh
	Refresh() (err Error)

	SafeToSilentClose() bool
	// Remove render tree from screen and close DOM and SOM and archive it.
	// But do clean logic after some time e.g. 10sec maybe user close by mistake click action
	Close() Error

	// DynamicElements struct {}
}

type GUI_PageState_ActivateOptions struct {
	Screen GUIScreen
	// It will effect just on this app and this page state only, not OS level.
	Orientation ScreenOrientation
	WindowType  GUI_WindowType
	PopUp       bool // force to bring to front immediately
}

type GUI_WindowType uint8

const (
	GUI_WindowType_Unset      GUI_WindowType = 0
	GUI_WindowType_Resizeable GUI_WindowType = (1 << iota) // has resizeable frame
	GUI_WindowType_GLASSY                                  // glassy window
	GUI_WindowType_ALPHA                                   // transparent window
)
