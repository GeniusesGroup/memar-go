/* For license and copyright information please see the LEGAL file in the code repository */

package block_p

import (
	container_p "memar/computer/adt/container/protocol"
	marker_p "memar/computer/adt/marker/protocol"
	adt_p "memar/computer/adt/protocol"
	storage_page_p "memar/storage/page/protocol"
	storage_p "memar/storage/protocol"
)

// Several pages are summarized to a block.
// A block is the smallest unit of access on a memory.
// Currently, 128 pages are mostly combined into one block; therefore with 4KB page, a block contains 512 KB.
type Block interface {
	storage_page_p.Field_PageNumber

	container_p.Clear
	// TODO::: Capacity and Length is `byte` or `page`??
	container_p.Field_Capacity
	container_p.Field_OccupiedLength
	container_p.Field_AvailableLength

	storage_p.NonVolatile

	adt_p.ADT

	marker_p.Clone[Block]
	marker_p.Copy[Block]

	// EventTarget[]
}

type BlockNumber container_p.NumberOfElement
type Field_BlockNumber interface {
	BlockNumber() BlockNumber
}

type BlockIndex container_p.ElementIndex
