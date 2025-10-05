/* For license and copyright information please see the LEGAL file in the code repository */

package thread_p

import (
	error_p "memar/process/error/protocol"
	"memar/time/duration"
)

type Field_ThreadID interface {
	ThreadID() ThreadID
}

type ThreadID uint64
