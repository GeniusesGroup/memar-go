/* For license and copyright information please see the LEGAL file in the code repository */

package l4

import (
	net_p "memar/net/protocol"
)

// var reSendBrokenPacketService = service.Service{
// 	URN:                "domain/srpc.scm.geniuses.group; type=service; name=resend-broken-packet",
// 	Domain:             DomainName,
// 	ID:                 16373901659613768690,
// 	IssueDate:          1595478242,
// 	ExpiryDate:         0,
// 	ExpireInFavorOfURN: "",
// 	ExpireInFavorOfID:  0,
// 	Status:             datatype_p.LifeCycle_PreAlpha,

// 	Authorization: authorization.Service{
// 		ActionType: operation_p.ActionType_Read,
// 		UserType:   user_p.TypeOwner,
// 	},

// 	Detail: map[lang_p.LanguageID]service.ServiceDetail{
// 		lang_p.LanguageEnglish: {
// 			Name:        "ReSend Broken Packet",
// 			Description: `use to send request to other party with broken packetID to send it again.`,
// 			TAGS:        []string{},
// 		},
// 	},

// 	SRPCHandler: ReSendBrokenPacket,
// }

// ReSendBrokenPacket use to send request to other party with broken packetID to send it again.
func ReSendBrokenPacket(sk net_p.Socket) {
}

type reSendBrokenPacketReq struct {
}
