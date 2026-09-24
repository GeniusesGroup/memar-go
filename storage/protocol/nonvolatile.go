/* For license and copyright information please see the LEGAL file in the code repository */

package storage_p

import (
	error_p "memar/process/error/protocol"
)

// NonVolatile is the interface that provide how a page or block forced to be nonvolatile
// If flush don't return any error, Caller MUST be sure that data had been keep secure,
// BUT it doesn't mean application have strong consistency! This storage CAN easily broken due to many reason.
type NonVolatile interface {
	// Flush force the storage device to write any changes to data (store in any cache layer like RAM, ...) before call Flush.
	Flush() (err error_p.Error)
}
