/* For license and copyright information please see the LEGAL file in the code repository */

package block_p

import (
	container_p "memar/computer/adt/container/protocol"
	error_p "memar/process/error/protocol"
	storage_page_p "memar/storage/page/protocol"
)

// Model is storage memory interface that show how an app access to volatile storage devices.
// It can use as volatile and nonvolatile storage.
// Usually dev must not use this interface due to it can damage any data written by upper abstractions such as objects, files, records, k/v, ...
type Model interface {
	// return volume capacity
	// these methods have eventual consistency behavior due to extend and other methods.
	Field_BlockNumber
	storage_page_p.Field_PageNumber

	// Extended length may vary of requested cap, Due to Extend() is base on storage device block size not bytes,
	Extend(cap BlockNumber) (extended BlockNumber, err error_p.Error)

	GetPage(pi storage_page_p.PageIndex) (p storage_page_p.Page, err error_p.Error)
	GetBlock(bi BlockIndex) (b Block, err error_p.Error)

	// container_p.Iteration[Page]
	// TODO::: `Search` and `Iteration` is same but `Search` is a process in storage hardware and don't need to transfer data to `RAM,CPU`
	Search(query []byte, offset container_p.Offset, limit container_p.Limit) (loc container_p.Offset, err error_p.Error)
}
