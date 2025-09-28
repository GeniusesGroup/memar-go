/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	error_p "memar/error/protocol"
)

//memar:impl memar/error/protocol.Internal
func IsInternal(err error_p.Error) bool {
	var interErr, ok = err.(error_p.Internal)
	return ok && interErr.Internal()
}

//memar:impl memar/error/protocol.Temporary
func IsTemporary(err error_p.Error) bool {
	var tempErr, ok = err.(error_p.Temporary)
	return ok && tempErr.Temporary()
}

//memar:impl memar/error/protocol.Timeout
func IsTimeout(err error_p.Error) bool {
	var timeErr, ok = err.(error_p.Timeout)
	return ok && timeErr.Timeout()
}
