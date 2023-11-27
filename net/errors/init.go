/* For license and copyright information please see the LEGAL file in the code repository */

package errs

func init() {
	ErrNoConnection.Init()
	ErrSendRequest.Init()
	ErrProtocolHandler.Init()
	ErrGuestNotAllow.Init()
	ErrGuestMaxReached.Init()
}
