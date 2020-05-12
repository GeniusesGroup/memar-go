/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

var getStreamsIDsService = achaemenid.Service{
	Name:       "GetStreamsIDs",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     achaemenid.ServiceStatePreAlpha,
	Handler:    GetStreamsIDs,
	Description: []string{
		`Use by peer can get knowledge about all active StreamID on other party that use to close streams due to Due MaxConcurrentStreams.`,
	},
	TAGS: []string{},
}

// GetStreamsIDs use by peer can get knowledge about all active StreamID on other party that use to close streams due to Due MaxConcurrentStreams.
func GetStreamsIDs(s *achaemenid.Server, st *achaemenid.Stream) {
}

type getStreamsIDsReq struct {
}
