/* For license and copyright information please see LEGAL file in repository */

package protocol

type GUIHistory interface {
	PreviousPageState() GUIPageState
	FollowingPageState() GUIPageState

	// Find states by PageID, Title, ...
}
