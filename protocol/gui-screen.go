/* For license and copyright information please see LEGAL file in repository */

package protocol

/** A screen, usually the one on which the current window is being rendered, and is obtained using window.screen. */
type GUIScreen interface {
	ID() int
	Type() ScreenType
	State() ScreenMode

	Height() int
	Width() int
	AvailHeight() int
	AvailWidth() int
	// The rate where screen updates are performed (per seconds).
	UpdateRate() int
	PixelDepth() int
	ColorDepth() int
	Orientation() ScreenOrientation

	Show()     // bring to front
	Minimize() // bring to back
	SetTitle(title string)

	LockOrientation(orientation ScreenOrientation)
	UnlockOrientation()

	EventTarget
}

type ScreenType uint8

// Unset "primary" "secondary" "any"

type ScreenOrientation uint8

// Unset or Any
// Natural Landscape" Portrait
const (
	// AnyOrientation allows the window to be freely orientated.
	AnyOrientation ScreenOrientation = iota
	// LandscapeOrientation constrains the window to landscape orientations.
	LandscapeOrientation
	// PortraitOrientation constrains the window to portrait orientations.
	PortraitOrientation
)

type ScreenEvent uint8

// absolute?: boolean;
// alpha?: number | null;
// beta?: number | null;
// gamma?: number | null;

// ScreenMode is the window mode (ScreenMode.Option sets it).
// Note that mode can be changed programatically as well as by the user
// clicking on the minimize/maximize buttons on the window's title bar.
type ScreenMode uint8

const (
	// Screened is the normal window mode with OS specific window decorations.
	Screened ScreenMode = iota
	// Fullscreen is the full screen window mode.
	Fullscreen
	// Minimized is for systems where the window can be minimized to an icon.
	Minimized
	// Maximized is for systems where the window can be made to fill the available monitor area.
	Maximized
)
