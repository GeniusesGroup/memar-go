/* For license and copyright information please see the LEGAL file in the code repository */

package storage_p

import (
	adt_p "memar/adt/protocol"
	buffer_p "memar/buffer/protocol"
	error_p "memar/error/protocol"
)

// BlockVolatile or StorageMemory is the interface that show how an app access to volatile storage devices.
// Usually dev must not use this interface due to it can damage any data written by objects, files, records, k/v interfaces
type BlockVolatile interface {
	// return volume capacity
	adt_p.Capacity

	// Extended length may vary of requested cap, Due to Extend() is base on storage device block size not bytes,
	// e.g. on SSDs block sizes are 256*page-size like 256*4(page-size)=1024(B)=1(MB)
	Extend(cap adt_p.NumberOfElement) (extended adt_p.NumberOfElement, err error_p.Error)

	// Change the return data not flush to any non-volatile storages. Use Write() to change data.
	Read(offset adt_p.ElementIndex, limit adt_p.NumberOfElement, buf buffer_p.Buffer) (err error_p.Error)
	// Write at more than block capacity cause block extend. extend capacity isn't equal to write length.
	Write(offset adt_p.ElementIndex, buf buffer_p.Buffer) (err error_p.Error)
	Search(data []byte, offset adt_p.ElementIndex) (loc adt_p.ElementIndex, err error_p.Error)

	Erase(offset adt_p.ElementIndex, limit adt_p.NumberOfElement) (err error_p.Error)
	Copy(desOffset, srcOffset adt_p.ElementIndex, limit adt_p.NumberOfElement) (err error_p.Error)
	Move(desOffset, srcOffset adt_p.ElementIndex, limit adt_p.NumberOfElement) (err error_p.Error)
}

// Block is the interface that show how an app access to storage devices.
// Usually dev must not use this interface due to it can damage any data written by objects, files, records or k/v interfaces
type Block interface {
	BlockVolatile

	// Flush force the storage device to write any changes to data (store in cache) before call Flush.
	Flush() (err error_p.Error)
}
