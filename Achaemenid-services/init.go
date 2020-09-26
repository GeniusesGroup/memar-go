/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

// Server store address location to server use by other part of app!
var server *achaemenid.Server

// Init use to register all available server services to given achaemenid.
func Init(s *achaemenid.Server) {
	server = s

	s.Services.RegisterService(&revokeConnectionService)
	s.Services.RegisterService(&closeConnectionService)
	s.Services.RegisterService(&pingPeerService)
	s.Services.RegisterService(&setConnectionSettingsService)
	s.Services.RegisterService(&getStreamsIDsService)
	s.Services.RegisterService(&closeStreamService)
	s.Services.RegisterService(&getPlatformServiceDetailService)
	s.Services.RegisterService(&getPlatformErrorDetailService)
	s.Services.RegisterService(&reSendBrokenPacketService)
	s.Services.RegisterService(&transferConnectionService)
	s.Services.RegisterService(&syncTimeService)
	// s.Services.RegisterService()
}
