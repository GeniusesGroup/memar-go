/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

var closeStreamService = achaemenid.Service{
	Name:       "CloseStream",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     achaemenid.ServiceStatePreAlpha,
	Handler:    CloseStream,
	Description: []string{
		`use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.`,
	},
	TAGS: []string{},
}

// CloseStream use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.
func CloseStream(s *achaemenid.Server, st *achaemenid.Stream) {
}

type closeStreamReq struct {
}
