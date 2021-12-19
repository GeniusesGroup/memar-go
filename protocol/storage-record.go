/* For license and copyright information please see LEGAL file in repository */

package protocol

// StorageRecords is the interface that show how an record storage work.
// Records are anything (such as a document or a phonograph record or a photograph) providing permanent evidence of or information about past events
// and equivalent to rows in RDBMS.
// - Record owner is one app so it must handle concurrent protection internally if not use version.
// Records or records or rows (RDBMS) all are same just difference in where schema must handle.
// Strongly think fundementally due to all data in computer must store as k/v in storages finally even files or rows.
// - 32 byte uuid choose to defeat https://en.wikipedia.org/wiki/Birthday_problem due to we suggest use hash(sha3-256) of any data as record key.
// - mediaTypeID equivalent to table name, namespace or BucketID in S3 standard, but mediaTypeID has more meaning.
// - version is always time.Now().UnixMilli() represent as int64 to make StorageRecords as a "Time Series DBMS" or
// use to store big records with same key but in multiple part. Strongly suggest to save large record in multiple small size parts.
// - If need encryption, Implement requirements at block storage level.
type StorageRecords interface {
	MediatypeNumbers() (num uint64, err Error)
	ListMediatypeIDs(offset, limit uint64) (ids []uint64, err Error)
	RecordNumbers(mediaTypeID uint64) (num uint64, err Error)
	ListRecords(mediaTypeID uint64, offset, limit uint64) (uuids [][32]byte, err Error)

	ListenToMediatype(mediaTypeID uint64, crud CRUD, uuid chan [32]byte) (err Error)
	ListenToKey(mediaTypeID uint64, uuid [32]byte, crud CRUD, versionOffset chan uint64) (err Error)

	// Lock works only in versioned manner.
	Lock(mediaTypeID uint64, uuid [32]byte) (record []byte, version int64, err Error)
	Unlock(mediaTypeID uint64, uuid [32]byte, record []byte) (version int64, err Error)

	// versionOffset+1 also indicate totall version number save for the record.
	LastVersion(mediaTypeID uint64, uuid [32]byte) (versionOffset uint64, version int64, err Error)
	Versions(mediaTypeID uint64, uuid [32]byte, offset, limit uint64) (ver []int64, err Error)
	Length(mediaTypeID uint64, uuid [32]byte, versionOffset uint64) (ln int, err Error)
	// Expiration is TTL (Time-To-Live). Number of seconds until record expires.
	// Expiration() int64

	// Get last version of record
	Get(mediaTypeID uint64, uuid [32]byte) (record []byte, versionOffset uint64, version int64, err Error)
	GetByVersion(mediaTypeID uint64, uuid [32]byte, version int64) (record []byte, versionOffset uint64, err Error)
	GetByVersionOffset(mediaTypeID uint64, uuid [32]byte, versionOffset uint64) (record []byte, version int64, err Error)
	// save or update record.
	// versionNumber == RecordNoVersion means this record don't need versioning.
	// versionNumber > 0 indicate max version. e.g. 6 means just 6 version must store for the record.
	// versionNumber == RecordLastVersionOffset indicate no version limit but logically it has limit up to uint64.
	Save(mediaTypeID uint64, uuid [32]byte, record []byte, versionNumber uint64) (version int64, err Error)
	Update(mediaTypeID uint64, uuid [32]byte, record []byte, versionOffset uint64) (version int64, err Error)
	// make invisible just by remove from primary index for all version of record.
	Delete(mediaTypeID uint64, uuid [32]byte) (err Error)
	// make invisible just by remove from primary index. next Get() can know that a version exist, but data gone and no access to data anymore.
	DeleteVersion(mediaTypeID uint64, uuid [32]byte, versionOffset uint64) (err Error)
}

const (
	RecordNoVersion         uint64 = 0x0000000000000000
	RecordLastVersionOffset uint64 = 0xFFFFFFFFFFFFFFFF
)
