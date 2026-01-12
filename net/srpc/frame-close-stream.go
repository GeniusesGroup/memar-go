/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	net_p "memar/net/protocol"
)

// var closeStreamService = service.Service{
// 	URN:                "domain/srpc.scm.geniuses.group; type=service; name=close-stream",
// 	Domain:             DomainName,
// 	ID:                 6917897595815184909,
// 	IssueDate:          1595478242,
// 	ExpiryDate:         0,
// 	ExpireInFavorOfURN: "",
// 	ExpireInFavorOfID:  0,
// 	Status:             datatype_p.LifeCycle_PreAlpha,

// 	Authorization: authorization.Service{
// 		ActionType:     operation_p.ActionType_Create,
// 		UserType: user_p.Type_All,
// 	},

// 	Detail: map[lang_p.LanguageID]service.ServiceDetail{
// 		lang_p.LanguageEnglish: {
// 			Name:        "Close Stream",
// 			Description: `use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.`,
// 			TAGS:        []string{},
// 		},
// 	},

// 	SRPCHandler: CloseStream,
// }

// CloseStream use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.
func CloseStream(sk net_p.Socket) {
}

type closeStreamReq struct {
}
