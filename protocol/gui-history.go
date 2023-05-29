/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type GUIHistory interface {
	PreviousPageState() GUI_PageState
	FollowingPageState() GUI_PageState

	// Find states by PageID, Title, ...
}
