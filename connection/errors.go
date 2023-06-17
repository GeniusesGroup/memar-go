/* For license and copyright information please see the LEGAL file in the code repository */

package connection

import (
	er "libgo/error"
)

var (
	ErrNoConnection    er.Error
	ErrSendRequest     er.Error
	ErrProtocolHandler er.Error
	ErrGuestNotAllow   er.Error
	ErrGuestMaxReached er.Error
)

func init() {
	ErrNoConnection.Init("domain/connection.protocol; type=error; name=no-connection")
	ErrSendRequest.Init("domain/connection.protocol; type=error; name=send-request")
	ErrProtocolHandler.Init("domain/connection.protocol; type=error; name=protocol-handler")
	ErrGuestNotAllow.Init("domain/connection.protocol; type=error; name=guest-not-allow")
	ErrGuestMaxReached.Init("domain/connection.protocol; type=error; name=guest-max-reached")
}
