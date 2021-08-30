/* For license and copyright information please see LEGAL file in repository */

package protocol

type GUIHistory interface {
	ActiveState() GUIPageState
	PreviousState() GUIPageState
	FollowingState() GUIPageState

	// Find states by PageID, Title, ...
}

// GUIPageState :
type GUIPageState interface {
	Page() GUIPage
	Title() string
	Description() string
	Conditions() map[string]string
	Fragment() string
	ActiveDate() Time
	EndDate() Time
}
