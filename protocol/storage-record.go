/* For license and copyright information please see LEGAL file in repository */

package protocol

// StorageRecords is the interface that show how an record storage work.
// Records are anything (such as a document or a phonograph record or a photograph) providing permanent evidence of or information about past events
// and equivalent to rows in RDBMS in some ways.
// - Record owner is one app so it must handle concurrent protection internally if not use version.
// Records or records or rows (RDBMS) all are same just difference in where schema must handle.
// Strongly think fundamentally due to all data in computer must store as k/v in storages finally even files or rows.
// - mediaTypeID equivalent to table name, namespace or BucketID in S3 standard, but mediaTypeID has more meaning.
// - version use to make StorageRecords as a "Time Series DBMS" or
// use to store big records with same key but in multiple part. Strongly suggest to save large record in multiple small size parts.
// - If need encryption, Implement requirements at block storage level.
type StorageRecords interface {
	MediatypeNumbers() (num uint64, err Error)
	ListMediatypeIDs(offset, limit uint64) (ids []uint64, err Error)
	RecordNumbers(mediaTypeID uint64) (num uint64, err Error)
	ListRecords(mediaTypeID uint64, offset, limit uint64) (ids [][16]byte, err Error)

	// Lock works only in versioned manner. use to reach strict consistency
	Lock(mediaTypeID uint64, id [16]byte) (lastVersion []byte, versionOffset uint64, err Error)
	Unlock(mediaTypeID uint64, id [16]byte, newVersion []byte) (err Error)

	// Count has eventual consistency behavior
	Count(mediaTypeID uint64, id [16]byte, offset, limit uint64) (numbers uint64, err Error)
	Length(mediaTypeID uint64, id [16]byte, versionOffset uint64) (ln int, err Error)

	Get(mediaTypeID uint64, id [16]byte, versionOffset uint64) (record []byte, numbers uint64, err Error)
	// GetLast has eventual consistency behavior
	// GetLast(mediaTypeID uint64, id [16]byte) (record []byte, versionOffset uint64, err Error)

	Save(mediaTypeID uint64, id [16]byte, record []byte, options StorageRecord_SaveOptions) (err Error)
	Update(mediaTypeID uint64, id [16]byte, record []byte, versionOffset uint64) (err Error)
	// make invisible just by remove from primary index for all version of record.
	Delete(mediaTypeID uint64, id [16]byte) (err Error)
	// make invisible just by remove from primary index. next Get() can know that a version exist, but data gone and no access to data anymore.
	DeleteVersion(mediaTypeID uint64, id [16]byte, versionOffset uint64) (err Error)

	EventTarget
}

type StorageRecord_SaveOptions struct {
	// MaxVersion == StorageRecord_NoVersion means this record don't need versioning.
	// MaxVersion > 0 indicate max version. e.g. 6 means just 6 version must store for the record.
	// MaxVersion == StorageRecord_LastSourceVersion indicate no version limit but logically it has limit up to uint64.
	MaxVersion uint64

	// By none, hour, day, week, ...
	PrimaryIndexSplitting uint8
}

const (
	StorageRecord_NoVersion         uint64 = 0
	StorageRecord_LastLocalVersion  uint64 = 18446744073709551613
	StorageRecord_LastEdgeVersion   uint64 = 18446744073709551614
	StorageRecord_LastSourceVersion uint64 = 18446744073709551615
)
