/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane  "../ChaparKhane"

var syncTimeService = chaparkhane.Service{
	Name:       "SyncTime",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    SyncTime,
	Description: []string{
		`use to sync time between peer!`,
	},
	TAGS: []string{},
}

// SyncTime use to sync time between peer!
// https://en.wikipedia.org/wiki/Network_Time_Protocol
func SyncTime(s *chaparkhane.Server, st *chaparkhane.Stream) {
	// If the payload is not empty, the recipient MUST generate a PONG frame containing the same Data.
}

type syncTimeReq struct {
}
