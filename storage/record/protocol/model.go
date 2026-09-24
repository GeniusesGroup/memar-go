/* For license and copyright information please see the LEGAL file in the code repository */

package record_p

import (
	container_p "memar/computer/adt/container/protocol"
	error_p "memar/process/error/protocol"
)

// Record model is the interface that show how an record storage work.
// - Records are anything (such as a document or a phonograph record or a photograph) providing permanent evidence of or information about past events
// - Information about something that is written down or stored on computer, film etc so that it can be looked at in the future
//
// Record equivalent to rows in RDBMS in some ways, or object storage (S3, ...).
// - Record owner is one app so it must handle concurrent protection internally if not use version.
// Records or records or rows (RDBMS) all are same just difference in where schema must handle.
// Strongly think fundamentally due to all data in computer must store as k/v in storages finally even files or rows.
// - mediaTypeID equivalent to table-name(RDB), namespace(Key-Value) or BucketID(AWS S3), but mediaTypeID has more meaning and has fixed size.
// - version use to make StorageRecords as a "Time Series DBMS" or
// - use to store big records with same key but in multiple part. Strongly suggest to save large record in multiple small size parts.
// - 16 byte id choose to defeat https://en.wikipedia.org/wiki/Birthday_problem due to we suggest use hash(sha3-256 to 128bit) of any data as object key.
// - If need encryption, Implement requirements at block storage level.
type Model interface {
	RecordsNumbers() (num container_p.NumberOfElement, err error_p.Error)
	ListRecords(offset container_p.ElementIndex, limit container_p.NumberOfElement) (ids []UUID, err error_p.Error)

	// Lock use to reach strict consistency
	Lock(id UUID) (lastVersion Data, err error_p.Error)
	Unlock(id UUID, newVersion Data) (err error_p.Error)

	// Count has eventual consistency behavior
	// Count(id UUID, offset container_p.ElementIndex, limit container_p.NumberOfElement) (numbers NumberOfVersion, err error_p.Error)
	Length(id UUID) (ln container_p.NumberOfElement, err error_p.Error)

	Get(id UUID) (record Data, err error_p.Error)
	// GetLast has eventual consistency behavior
	// GetLast(id UUID) (record Buffer, vo VersionOffset, err error_p.Error)
	Read(id UUID, offset container_p.ElementIndex, limit container_p.NumberOfElement) (data Data, err error_p.Error)

	Save(id UUID, record Data) (err error_p.Error)
	Write(id UUID, offset container_p.ElementIndex, data Data) (err error_p.Error)

	Append(id UUID, data Data) (err error_p.Error)
	Prepend(id UUID, data Data) (err error_p.Error)
	Extend(id UUID, length container_p.NumberOfElement) (err error_p.Error)

	// make invisible just by remove from primary index.
	Delete(id UUID) (err error_p.Error)
	// make invisible by remove from primary index & write zero data to record location in NVM.
	Erase(id UUID) (err error_p.Error)

	// EventTarget[]

	// Multiple changes can be made in one atomic batch
	// Batch()
}

type StorageRecord_SaveOptions struct {
	// By none, hour, day, week, ...
	PrimaryIndexSplitting uint8
}
