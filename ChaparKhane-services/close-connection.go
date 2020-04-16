/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane  "../ChaparKhane"

var closeConnectionService = chaparkhane.Service{
	Name:       "CloseConnection",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    CloseConnection,
	Description: []string{
		`Use to close connection and drop all incomplete data & stream`,
	},
	TAGS: []string{},
}

// CloseConnection use to close connection and drop all incomplete data & stream.
// Connection can be open again by use TransferConnection() if related service exist in platform!
func CloseConnection(s *chaparkhane.Server, st *chaparkhane.Stream) {
}

type closeConnectionReq struct {
}
