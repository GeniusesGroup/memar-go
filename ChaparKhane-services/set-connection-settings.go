/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane  "../ChaparKhane"

var setConnectionSettingsService = chaparkhane.Service{
	Name:       "SetConnectionSettings",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    SetConnectionSettings,
	Description: []string{
		`use to change settings set in RegisterConnection
		 Don't send this request between active stream due may streams data will lost!`,
	},
	TAGS: []string{},
}

// SetConnectionSettings use to change settings set in RegisterConnection
// Don't send this request between active stream due may streams data will lost!
func SetConnectionSettings(s *chaparkhane.Server, st *chaparkhane.Stream) {}

type setConnectionSettingsReq struct {
	PacketPayloadSize uint16 // Defaults: 1200 byte. It can't be under 1200 byte. Exclude network or transport header.
	// Packet data compression type e.g. gzip, ...
}
