/* For license and copyright information please see the LEGAL file in the code repository */

package os_p

import (
	storage_p "memar/storage/protocol"
)

// OperatingSystem_Storage is the interface that must implement by any OS object
type OperatingSystem_Storage interface {
	Storage() storage_p.Block
	// Any data transfer between host and drive will be in multiples of logical block size (PhysicalSectorSize)
	StorageDevices() []StorageBlockDetails
}
