/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane  "../ChaparKhane"

var pingPeerService = chaparkhane.Service{
	Name:       "PingPeer",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    PingPeer,
	Description: []string{
		`Endpoints can use PING to verify that their peers are still alive or to check reachability to the peer.`,
	},
	TAGS: []string{},
}

// PingPeer : Endpoints can use PING to verify that their peers are still alive or to check reachability to the peer.
func PingPeer(s *chaparkhane.Server, st *chaparkhane.Stream) {
	// If the payload is not empty, the recipient MUST generate a PONG frame containing the same Data.
}

type pingPeerReq struct {
}
