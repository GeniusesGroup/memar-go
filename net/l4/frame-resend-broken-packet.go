/* For license and copyright information please see the LEGAL file in the code repository */

package l4

import (
	"memar/protocol"
)

// var reSendBrokenPacketService = service.Service{
// 	URN:                "domain/srpc.scm.geniuses.group; type=service; name=resend-broken-packet",
// 	Domain:             DomainName,
// 	ID:                 16373901659613768690,
// 	IssueDate:          1595478242,
// 	ExpiryDate:         0,
// 	ExpireInFavorOfURN: "",
// 	ExpireInFavorOfID:  0,
// 	Status:             protocol.Software_PreAlpha,

// 	Authorization: authorization.Service{
// 		CRUD:     authorization.CRUDRead,
// 		UserType: protocol.UserTypeOwner,
// 	},

// 	Detail: map[protocol.LanguageID]service.ServiceDetail{
// 		protocol.LanguageEnglish: {
// 			Name:        "ReSend Broken Packet",
// 			Description: `use to send request to other party with broken packetID to send it again.`,
// 			TAGS:        []string{},
// 		},
// 	},

// 	SRPCHandler: ReSendBrokenPacket,
// }

// ReSendBrokenPacket use to send request to other party with broken packetID to send it again.
func ReSendBrokenPacket(sk protocol.Socket) {
}

type reSendBrokenPacketReq struct {
}
