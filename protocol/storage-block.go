/* For license and copyright information please see LEGAL file in repository */

package protocol

// StorageMemory is the interface that show how an app access to volatile storage devices.
// Usually dev must not use this interface due to it can damage any data written by objects, files or records interfaces
type StorageBlockVolatile interface {
	// return volume capacity
	Cap() int
	Extend(length int) (err Error)

	// Change the return data not flush to any non-volatile storages. Use Write() to change data.
	Read(offset, limit int) (data []byte, err Error)
	// Write at more than block capacity cause block extend. extend capacity isn't equal to write length.
	Write(offset, data []byte) (err Error)
	Erase(offset, limit int) (err Error)

	Copy(desOffset, srcOffset int, limit int) (err Error)
	Move(desOffset, srcOffset int, limit int) (err Error)

	Search(data []byte, offset int) (loc int, err Error)
}

// StorageBlock is the interface that show how an app access to storage devices.
// Usually dev must not use this interface due to it can damage any data written by objects, files, records or k/v interfaces
type StorageBlock interface {
	StorageBlockVolatile

	// Any changes to data before call Flush will write to storage.
	Flush() (err Error)
}
