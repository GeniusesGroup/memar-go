/* For license and copyright information please see the LEGAL file in the code repository */

package block_p

import (
	container_p "memar/adt/container/protocol"
	buffer_p "memar/buffer/protocol"
	error_p "memar/error/protocol"
)

// Volatile or StorageMemory is the interface that show how an app access to volatile storage devices.
// Usually dev must not use this interface due to it can damage any data written by objects, files, records, k/v interfaces
type Volatile interface {
	// return volume capacity
	container_p.Capacity

	// Extended length may vary of requested cap, Due to Extend() is base on storage device block size not bytes,
	// e.g. on SSDs block sizes are 256*page-size like 256*4(page-size)=1024(B)=1(MB)
	Extend(cap container_p.NumberOfElement) (extended container_p.NumberOfElement, err error_p.Error)

	// Change the return data not flush to any non-volatile storages. Use Write() to change data.
	Read(offset container_p.ElementIndex, limit container_p.NumberOfElement, buf buffer_p.Buffer) (err error_p.Error)
	// Write at more than block capacity cause block extend. extend capacity isn't equal to write length.
	Write(offset container_p.ElementIndex, buf buffer_p.Buffer) (err error_p.Error)
	Search(data []byte, offset container_p.ElementIndex) (loc container_p.ElementIndex, err error_p.Error)

	Erase(offset container_p.ElementIndex, limit container_p.NumberOfElement) (err error_p.Error)
	Copy(desOffset, srcOffset container_p.ElementIndex, limit container_p.NumberOfElement) (err error_p.Error)
	Move(desOffset, srcOffset container_p.ElementIndex, limit container_p.NumberOfElement) (err error_p.Error)
}

// NonVolatile is the interface that show how an app access to storage devices.
// Usually dev must not use this interface due to it can damage any data written by objects, files, records or k/v interfaces
type NonVolatile interface {
	Volatile

	// Flush force the storage device to write any changes to data (store in cache) before call Flush.
	Flush() (err error_p.Error)
}
