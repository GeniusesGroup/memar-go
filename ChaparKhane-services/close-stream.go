/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane "../ChaparKhane"

var closeStreamService = chaparkhane.Service{
	Name:       "CloseStream",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    CloseStream,
	Description: []string{
		`use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.`,
	},
	TAGS: []string{},
}

// CloseStream use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.
func CloseStream(s *chaparkhane.Server, st *chaparkhane.Stream) {
}

type closeStreamReq struct {
}
