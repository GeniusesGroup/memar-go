/* For license and copyright information please see the LEGAL file in the code repository */

package errs

const domainBaseMediatype = "domain/memar.scm.geniuses.group; package=achaemenid; type=error; "

func init() {
	ErrBadRequest.Init()
	ErrBadResponse.Init()
	ErrInternalError.Init()
	ErrProtocolHandler.Init()
	ErrGuestConnectionNotAllow.Init()
	ErrGuestConnectionMaxReached.Init()
	ErrNotStandardID.Init()
}
