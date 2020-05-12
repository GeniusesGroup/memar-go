/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

var setConnectionSettingsService = achaemenid.Service{
	Name:       "SetConnectionSettings",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     achaemenid.ServiceStatePreAlpha,
	Handler:    SetConnectionSettings,
	Description: []string{
		`use to change settings set in RegisterConnection
		 Don't send this request between active stream due may streams data will lost!`,
	},
	TAGS: []string{},
}

// SetConnectionSettings use to change settings set in RegisterConnection
// Don't send this request between active stream due may streams data will lost!
func SetConnectionSettings(s *achaemenid.Server, st *achaemenid.Stream) {}

type setConnectionSettingsReq struct {
	PacketPayloadSize uint16 // Defaults: 1200 byte. It can't be under 1200 byte. Exclude network or transport header.
	// Packet data compression type e.g. gzip, ...
}
