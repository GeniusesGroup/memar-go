/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

var pingPeerService = achaemenid.Service{
	Name:       "PingPeer",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     achaemenid.ServiceStatePreAlpha,
	Handler:    PingPeer,
	Description: []string{
		`Endpoints can use PING to verify that their peers are still alive or to check reachability to the peer.`,
	},
	TAGS: []string{},
}

// PingPeer : Endpoints can use PING to verify that their peers are still alive or to check reachability to the peer.
func PingPeer(s *achaemenid.Server, st *achaemenid.Stream) {
	// If the payload is not empty, the recipient MUST generate a PONG frame containing the same Data.
}

type pingPeerReq struct {
}
