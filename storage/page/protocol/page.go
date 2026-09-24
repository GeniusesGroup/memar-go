/* For license and copyright information please see the LEGAL file in the code repository */

package storage_page_p

import (
	container_p "memar/computer/adt/container/protocol"
	marker_p "memar/computer/adt/marker/protocol"
	adt_p "memar/computer/adt/protocol"
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
	sync_p "memar/process/sync/protocol"
	storage_p "memar/storage/protocol"
)

// Page or sector is the smallest unit, which is composed of several memory cells,
// and is usually 4 KB in size as advanced format of physical 4KB sectors.
// https://en.wikipedia.org/wiki/Advanced_Format
// In the modern SSD with 25 or 20 nm elements, the page size is 8KB (8192B), 16kB (16384B), 32KB(32768B) or even 64KB(65536B).
// e.g. on SSDs block sizes are 256*page-size like 256*4(page-size)=1024(B)=1(MB)
//
// The minimum physical storage unit of a hard disk drive (HDD) is a sector. The solid state drive (SSD) equivalent is a page.
//
// A page can't be rewritten without read first due to change in checksum know as ECC(Error Correction Code) and need to recalculation.
type Page interface {
	adt_p.ADT

	marker_p.Clone[Page]
	marker_p.Move[Page]
	Copy(desOffset, srcOffset container_p.Offset, limit container_p.Limit) (err error_p.Error)
	Move(desOffset, srcOffset container_p.Offset, limit container_p.Limit) (err error_p.Error)

	sync_p.Locker

	// Flash memory written in pages and erased in blocks, so once a page is written, a rewrite require the whole block to be erased first.
	storage_p.NonVolatile

	// container_p.Accessor[Byte]
	// container_p.Assignment[Byte]
	container_p.Field_Capacity
	container_p.Field_OccupiedLength
	container_p.Field_AvailableLength
	container_p.Delete
	container_p.Clear

	// EventTarget[]

	// Change the return data not flush to any non-volatile storages. Use Write() to change data.
	Read(offset container_p.Offset, limit container_p.Limit, buf buffer_p.Buffer) (err error_p.Error)
	// Write at more than block capacity cause block extend. extend capacity isn't equal to write length.
	Write(offset container_p.Offset, buf buffer_p.Buffer) (err error_p.Error)

	Search(data []byte, offset container_p.Offset, limit container_p.Limit) (loc container_p.Offset, err error_p.Error)
}

type PageNumber container_p.NumberOfElement
type Field_PageNumber interface {
	// TODO::: on bad page e.g. bad sector return 0 or have error?
	PageNumber() PageNumber
}

type PageIndex container_p.ElementIndex
