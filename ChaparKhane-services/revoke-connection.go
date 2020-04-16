/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane  "../ChaparKhane"

var revokeConnectionService = chaparkhane.Service{
	Name:       "RevokeConnection",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    RevokeConnection,
	Description: []string{
		`use to signal all server to revoke connection due leak data or user requested.`,
	},
	TAGS: []string{},
}

// RevokeConnection use to signal all server to revoke connection due leak data or user requested.
func RevokeConnection(s *chaparkhane.Server, st *chaparkhane.Stream) {
	// It can just send by related Domain.Service
}

type revokeConnectionReq struct {
}
