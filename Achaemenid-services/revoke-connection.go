/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

var revokeConnectionService = achaemenid.Service{
	Name:       "RevokeConnection",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     achaemenid.ServiceStatePreAlpha,
	Handler:    RevokeConnection,
	Description: []string{
		`use to signal all server to revoke connection due leak data or user requested.`,
	},
	TAGS: []string{},
}

// RevokeConnection use to signal all server to revoke connection due leak data or user requested.
func RevokeConnection(s *achaemenid.Server, st *achaemenid.Stream) {
	// It can just send by related Domain.Service
}

type revokeConnectionReq struct {
}
