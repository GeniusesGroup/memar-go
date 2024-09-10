/* For license and copyright information please see the LEGAL file in the code repository */

package primitive_p

import (
	error_p "memar/error/protocol"
)

type Locker interface {
	Lock() (err error_p.Error)
	Unlock() (err error_p.Error)
}
