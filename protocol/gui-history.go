/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type GUI_History interface {
	PreviousPageState() GUI_PageState
	FollowingPageState() GUI_PageState

	// Find states by PageID, Title, ...
}
