/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

var closeConnectionService = achaemenid.Service{
	Name:       "CloseConnection",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     achaemenid.ServiceStatePreAlpha,
	Handler:    CloseConnection,
	Description: []string{
		`Use to close connection and drop all incomplete data & stream`,
	},
	TAGS: []string{},
}

// CloseConnection use to close connection and drop all incomplete data & stream.
// Connection can be open again by use TransferConnection() if related service exist in platform!
func CloseConnection(s *achaemenid.Server, st *achaemenid.Stream) {
}

type closeConnectionReq struct {
}
