/* For license and copyright information please see LEGAL file in repository */

package ss

import "../achaemenid"

// Init use to register all available server services to given achaemenid.
func Init(s *achaemenid.Server) {
	s.Services.RegisterService(&registerGuestConnectionService)
	// s.Services.RegisterService()
	// s.Services.RegisterService()
}
