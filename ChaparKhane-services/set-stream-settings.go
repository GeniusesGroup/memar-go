/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane  "../ChaparKhane"

var setStreamSettingsService = chaparkhane.Service{
	Name:       "SetStreamSettings",
	IssueDate:  0,
	ExpiryDate: 0,
	Status:     chaparkhane.ServiceStatePreAlpha,
	Handler:    SetStreamSettings,
	Description: []string{
		`use to set stream settings like time sensitive use in VoIP, IPTV, ...`,
	},
	TAGS: []string{},
}

// SetStreamSettings use to set stream settings like time sensitive use in VoIP, IPTV, ...
func SetStreamSettings(s *chaparkhane.Server, st *chaparkhane.Stream) {
	// Dropping packets is preferable to waiting for packets delayed due to retransmissions.
	// Developer can ask to complete data for offline usage after first data usage.
}

type setStreamTimeSensitiveReq struct {
}
