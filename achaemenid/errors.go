/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
	er "libgo/error"
)

// Errors
var (
	ErrBadRequest                er.Error
	ErrBadResponse               er.Error
	ErrInternalError             er.Error
	ErrProtocolHandler           er.Error
	ErrGuestConnectionNotAllow   er.Error
	ErrGuestConnectionMaxReached er.Error
	ErrNotStandardID             er.Error
)

func init() {
	ErrBadRequest.Init("domain/libgo.scm.geniuses.group; package=achaemenid; type=error; name=bad-request")
	ErrBadResponse.Init("domain/libgo.scm.geniuses.group; package=achaemenid; type=error; name=bad-response")
	ErrInternalError.Init("domain/libgo.scm.geniuses.group; package=achaemenid; type=error; name=internal-error")
	ErrProtocolHandler.Init("domain/libgo.scm.geniuses.group; package=achaemenid; type=error; name=protocol-handler")
	ErrGuestConnectionNotAllow.Init("domain/libgo.scm.geniuses.group; package=achaemenid; type=error; name=guest-connection-not-allow")
	ErrGuestConnectionMaxReached.Init("domain/libgo.scm.geniuses.group; package=achaemenid; type=error; name=guest-connection-max-reached")
	ErrNotStandardID.Init("domain/libgo.scm.geniuses.group; package=achaemenid; type=error; name=not-standard-id")
}
