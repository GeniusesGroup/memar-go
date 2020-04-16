/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane  "../ChaparKhane"

var reSendBrokenPacketService = chaparkhane.Service{
	Name:       "ReSendBrokenPacket",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    ReSendBrokenPacket,
	Description: []string{
		`use to send request to other party with broken packetID to send it again.`,
	},
	TAGS: []string{},
}

// ReSendBrokenPacket use to send request to other party with broken packetID to send it again.
func ReSendBrokenPacket(s *chaparkhane.Server, st *chaparkhane.Stream) {
}

type reSendBrokenPacketReq struct {
}
