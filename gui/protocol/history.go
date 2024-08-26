/* For license and copyright information please see the LEGAL file in the code repository */

package gui_p

type History interface {
	PreviousPageState() PageState
	FollowingPageState() PageState

	// Find states by PageID, Title, ...
}
