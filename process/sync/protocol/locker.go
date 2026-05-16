/* For license and copyright information please see the LEGAL file in the code repository */

package sync_p

import (
	error_p "memar/process/error/protocol"
)

type Locker interface {
	// Be aware that the requested `thread` which call `Lock()` CAN `yield`!
	Lock() (err error_p.Error)
	Unlock() (err error_p.Error)
}

type ReadLocker interface {
	// Be aware that the requested `thread` which call `Lock()` CAN `yield`!
	RLock() (err error_p.Error)
	RUnlock() (err error_p.Error)
}
