/* For license and copyright information please see LEGAL file in repository */

package protocol

// Screen indicate some information about screen that a page-state(window) can render to it.
// https://developer.mozilla.org/en-US/docs/Web/API/Screen
type GUIScreen interface {
	ID() int
	Type() ScreenType
	Mode() ScreenMode

	Height() int
	Width() int
	AvailHeight() int
	AvailWidth() int
	// The rate where screen updates are performed (per seconds).
	UpdateRate() int
	PixelDepth() int
	ColorDepth() int
	Orientation() ScreenOrientation

	EventTarget
}

type ScreenType uint8

const (
	ScreenType_Unset ScreenType = iota
	ScreenType_Primary
	ScreenType_Secondary
	ScreenType_Extend
	ScreenType_Duplicate
)

// ScreenMode is the window mode (ScreenMode.Option sets it).
// Note that mode can be changed programmatically as well as by the user
// clicking on the minimize/maximize buttons on the window's title bar.
type ScreenMode uint8

const (
	ScreenMode_Unset ScreenMode = iota
	// Screened is the normal window mode with any OS specific window decorations.
	ScreenMode_Screened
	// FullScreen is the full screen window mode.
	ScreenMode_FullScreen
	// Minimized is for systems where the window can be minimized to an icon.
	ScreenMode_Minimized
	// Maximized is for systems where the window can be made to fill the available monitor area.
	ScreenMode_Maximized
)

type ScreenOrientation uint8

const (
	// ScreenOrientation_Any or Unset or Natural allows the window to be freely orientated.
	ScreenOrientation_Any ScreenOrientation = iota
	// ScreenOrientation_Landscape constrains the window to landscape orientations.
	ScreenOrientation_Landscape
	// ScreenOrientation_Portrait constrains the window to portrait orientations.
	ScreenOrientation_Portrait
)

type ScreenEvent uint8

// absolute?: boolean;
// alpha?: number | null;
// beta?: number | null;
// gamma?: number | null;
