/* For license and copyright information please see LEGAL file in repository */

package ss

import chaparkhane "../ChaparKhane"

// Init use to register all available server services to given chaparkhane.
func Init(s *chaparkhane.Server) {
	s.Services.RegisterService(&registerGuestConnectionService)
	// s.Services.RegisterService()
	// s.Services.RegisterService()
}
