/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane "../ChaparKhane"

// Init use to register all available server services to given chaparkhane.
func Init(s *chaparkhane.Server) {
	s.Services.RegisterService(&revokeConnectionService)
	s.Services.RegisterService(&closeConnectionService)
	s.Services.RegisterService(&pingPeerService)
	s.Services.RegisterService(&setConnectionSettingsService)
	s.Services.RegisterService(&getStreamsIDsService)
	s.Services.RegisterService(&closeStreamService)
	// s.Services.RegisterService()
	s.Services.RegisterService(&setStreamSettingsService)
	s.Services.RegisterService(&reSendBrokenPacketService)
	s.Services.RegisterService(&transferConnectionService)
	s.Services.RegisterService(&syncTimeService)
	// s.Services.RegisterService()
}
