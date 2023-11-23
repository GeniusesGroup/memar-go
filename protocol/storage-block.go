/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// StorageBlockVolatile or StorageMemory is the interface that show how an app access to volatile storage devices.
// Usually dev must not use this interface due to it can damage any data written by objects, files, records, k/v interfaces
type StorageBlockVolatile interface {
	// return volume capacity
	Cap

	// Extended length may vary of requested cap, Due to Extend() is base on storage device block size not bytes,
	// e.g. on SSDs block sizes are 256*page-size like 256*4(page-size)=1024(B)=1(MB)
	Extend(cap int) (extended int, err Error)

	// Change the return data not flush to any non-volatile storages. Use Write() to change data.
	Read(offset, data []byte) (err Error)
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

	// Flush force the storage device to write any changes to data (store in cache) before call Flush.
	Flush() (err Error)
}
