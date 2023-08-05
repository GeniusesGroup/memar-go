/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/protocol"
)

// var closeStreamService = service.Service{
// 	URN:                "domain/srpc.scm.geniuses.group; type=service; name=close-stream",
// 	Domain:             DomainName,
// 	ID:                 6917897595815184909,
// 	IssueDate:          1595478242,
// 	ExpiryDate:         0,
// 	ExpireInFavorOfURN: "",
// 	ExpireInFavorOfID:  0,
// 	Status:             protocol.Software_PreAlpha,

// 	Authorization: authorization.Service{
// 		CRUD:     authorization.CRUDCreate,
// 		UserType: protocol.UserType_All,
// 	},

// 	Detail: map[protocol.LanguageID]service.ServiceDetail{
// 		protocol.LanguageEnglish: {
// 			Name:        "Close Stream",
// 			Description: `use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.`,
// 			TAGS:        []string{},
// 		},
// 	},

// 	SRPCHandler: CloseStream,
// }

// CloseStream use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.
func CloseStream(sk protocol.Socket) {
}

type closeStreamReq struct {
}
