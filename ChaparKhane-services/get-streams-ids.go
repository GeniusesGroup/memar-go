/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane  "../ChaparKhane"

var getStreamsIDsService = chaparkhane.Service{
	Name:       "GetStreamsIDs",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    GetStreamsIDs,
	Description: []string{
		`Use by peer can get knowledge about all active StreamID on other party that use to close streams due to Due MaxConcurrentStreams.`,
	},
	TAGS: []string{},
}

// GetStreamsIDs use by peer can get knowledge about all active StreamID on other party that use to close streams due to Due MaxConcurrentStreams.
func GetStreamsIDs(s *chaparkhane.Server, st *chaparkhane.Stream) {
}

type getStreamsIDsReq struct {
}
