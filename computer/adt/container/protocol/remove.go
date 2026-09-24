/* For license and copyright information please see the LEGAL file in the code repository */

package container_p

import (
	error_p "memar/process/error/protocol"
)

// Delete is an operation that just mark as empty
type Delete interface {
	// Delete will just mark desire container location as empty to let other write operations CAN store data
	// if offset is 0 means it is from start of container.
	// if limit is 0 means to end of container.
	Delete(offset Offset, limit Limit) (err error_p.Error)
}

// Clear is an operation that wipe all data by write 0|1 to memory location.
type Clear interface {
	// Clear will remove (erase) all elements.
	// if offset is 0 means it is from start of container.
	// if limit is 0 means to end of container.
	Clear(offset Offset, limit Limit) (err error_p.Error)
}
