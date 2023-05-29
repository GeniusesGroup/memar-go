/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// OperatingSystem_Storage is the interface that must implement by any OS object
type OperatingSystem_Storage interface {
	Storage() StorageBlock
	// Any data transfer between host and drive will be in multiples of logical block size (PhysicalSectorSize)
	StorageDevices() []StorageBlockDetails
}
