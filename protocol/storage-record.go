/* For license and copyright information please see the LEGAL file in the code repository */

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
	ListMediatypeIDs(offset ElementIndex, limit NumberOfElement) (ids []uint64, err Error)

	RecordNumbers(mt MediaTypeID) (num uint64, err Error)
	ListRecords(mt MediaTypeID, offset ElementIndex, limit NumberOfElement) (ids []UUID, err Error)

	// Lock works only in versioned manner. use to reach strict consistency
	Lock(mt MediaTypeID, id UUID) (lastVersion Buffer, vo VersionOffset, err Error)
	Unlock(mt MediaTypeID, id UUID, newVersion Buffer) (err Error)

	// Count has eventual consistency behavior
	Count(mt MediaTypeID, id UUID, offset ElementIndex, limit NumberOfElement) (numbers NumberOfVersion, err Error)
	Length(mt MediaTypeID, id UUID, vo VersionOffset) (ln NumberOfElement, err Error)

	Get(mt MediaTypeID, id UUID, vo VersionOffset) (record Buffer, numbers NumberOfVersion, err Error)
	// GetLast has eventual consistency behavior
	// GetLast(mt MediaTypeID, id UUID) (record Buffer, vo VersionOffset, err Error)

	Save(mt MediaTypeID, id UUID, record Buffer) (err Error)
	Update(mt MediaTypeID, id UUID, record Buffer, vo VersionOffset) (err Error)
	// make invisible just by remove from primary index for all version of record.
	Delete(mt MediaTypeID, id UUID) (err Error)
	// make invisible just by remove from primary index. next Get() can know that a version exist, but data gone and no access to data anymore.
	DeleteVersion(mt MediaTypeID, id UUID, vo VersionOffset) (err Error)

	EventTarget
}

type StorageRecord_SaveOptions struct {
	// By none, hour, day, week, ...
	PrimaryIndexSplitting uint8
}
