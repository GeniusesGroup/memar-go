/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

var reSendBrokenPacketService = achaemenid.Service{
	Name:       "ReSendBrokenPacket",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     achaemenid.ServiceStatePreAlpha,
	Handler:    ReSendBrokenPacket,
	Description: []string{
		`use to send request to other party with broken packetID to send it again.`,
	},
	TAGS: []string{},
}

// ReSendBrokenPacket use to send request to other party with broken packetID to send it again.
func ReSendBrokenPacket(s *achaemenid.Server, st *achaemenid.Stream) {
}

type reSendBrokenPacketReq struct {
}
