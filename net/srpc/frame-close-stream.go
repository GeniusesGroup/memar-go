/* For license and copyright information please see LEGAL file in repository */

package srpc

// import (
// 	"../authorization"
// 	"../protocol"
// )

// var closeStreamService = service.Service{
// 	URN:                "urn:giti:achaemenid.protocol:service:close-stream",
// 	Domain:             DomainName,
// 	ID:                 6917897595815184909,
// 	IssueDate:          1595478242,
// 	ExpiryDate:         0,
// 	ExpireInFavorOfURN: "",
// 	ExpireInFavorOfID:  0,
// 	Status:             protocol.SoftwareStatePreAlpha,

// 	Authorization: authorization.Service{
// 		CRUD:     authorization.CRUDCreate,
// 		UserType: authorization.UserTypeAll,
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

// // CloseStream use by peer to close unwanted active StreamID on other party due to MaxConcurrentStreams restriction.
// func CloseStream(st protocol.Stream) {
// }

// type closeStreamReq struct {
// }
