//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package authorization

import (
	"libgo/protocol"
)

const domainEnglish = "Authorization"

func init() {
	ErrUserNotAllow.SetDetail(protocol.LanguageEnglish, domainEnglish, "User Not Allow",
		"Request service is not allow by user of the connection",
		"",
		"",
		nil)
	ErrUserNotOwnRecord.SetDetail(protocol.LanguageEnglish, domainEnglish, "User Not Own Record",
		"Request record is not own by user of the connection and by service rule can't access to it by other users",
		"",
		"",
		nil)

		ErrNotAllowSociety.SetDetail(protocol.LanguageEnglish, domainEnglish, "Not Allow Society",
			"Request send by society that is not in allow list of the connection",
			"",
			"",
			nil)
	ErrDeniedSociety.SetDetail(protocol.LanguageEnglish, domainEnglish, "Denied Society",
		"Request send by society that is in deny list of the connection",
		"",
		"",
		nil)
		
	ErrNotAllowRouter.SetDetail(protocol.LanguageEnglish, domainEnglish, "Not Allow Router",
	"Request send by router that is not in allow list of the connection",
	"",
	"",
	nil)
	ErrDeniedRouter.SetDetail(protocol.LanguageEnglish, domainEnglish, "Denied Router",
		"Request send by router that is in deny list of the connection",
		"",
		"",
		nil)
		
	ErrDayNotAllow.SetDetail(protocol.LanguageEnglish, domainEnglish, "Day Not Allow",
	"Request send in the day of week that is not in allow list of the connection",
	"",
	"",
	nil)
	ErrDayDenied.SetDetail(protocol.LanguageEnglish, domainEnglish, "Day Denied",
		"Request send in the day of week that is in deny list of the connection",
		"",
		"",
		nil)

	ErrHourNotAllow.SetDetail(protocol.LanguageEnglish, domainEnglish, "Hour Not Allow",
	"Request send in the hour of day that is not in allow list of the connection",
	"",
	"",
	nil)
ErrHourDenied.SetDetail(protocol.LanguageEnglish, domainEnglish, "Hour Denied",
	"Request send in the hour of day that is in deny list of the connection",
	"",
	"",
	nil)
	
	ErrServiceNotAllow.SetDetail(protocol.LanguageEnglish, domainEnglish, "Service Not Allow",
		"Request service is not in allow list of the connection",
		"",
		"",
		nil)
	ErrServiceDenied.SetDetail(protocol.LanguageEnglish, domainEnglish, "Service Denied",
		"Request service is in deny list of the connection",
		"",
		"",
		nil)

		ErrCrudNotAllow.SetDetail(protocol.LanguageEnglish, domainEnglish, "CRUD Not Allow",
			"Request type is not in allow list of the connection",
			"",
			"",
			nil)
		ErrCRUDDenied.SetDetail(protocol.LanguageEnglish, domainEnglish, "CRUD Denied",
			"Request type is deny list of the connection",
			"",
			"",
			nil)
			
	ErrNotAllowToDelegate.SetDetail(protocol.LanguageEnglish, domainEnglish, "Not Allow To Delegate",
	"Platforms rules not permit active type of user to register the delegate connection with given details",
	"",
	"",
	nil)
ErrNotAllowToNotDelegate.SetDetail(protocol.LanguageEnglish, domainEnglish, "Not Allow To Not Delegate",
	"Platforms rules not permit active type of user (usually Org type) register not delegate connection.",
	"",
	"",
	nil)
}
