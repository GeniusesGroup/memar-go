/* For license and copyright information please see LEGAL file in repository */

package protocol

// StorageObjects is the interface that show how an object directory work.
// Object owner is one app so it must handle concurrent protection internally.
// Strongly think fundementally due to all data in computer must store as k/v in storages finally even files or rows.
// - mediaTypeID equivalent to table-name(RDB), namespace(Key-Value) or BucketID(AWS S3), but mediaTypeID has more meaning and has fixed size.
// - 16 byte id choose to defeat https://en.wikipedia.org/wiki/Birthday_problem due to we suggest use hash(sha3-256 to 128bit) of any data as object key.
// If need encryption, Implement requirements at block storage level.
type StorageObjects interface {
	MediatypeNumbers() (num uint64, err Error)
	ListMediatypeIDs(offset, limit uint64) (ids []uint64, err Error)
	ObjectNumbers(mediaTypeID uint64) (num uint64, err Error)
	ListObjects(mediaTypeID uint64, offset, limit uint64) (ids [][16]byte, err Error)

	Lock(mediaTypeID uint64, id [16]byte) (err Error)
	Unlock(mediaTypeID uint64, id [16]byte) (err Error)

	Length(mediaTypeID uint64, id [16]byte) (ln int, err Error)
	Get(mediaTypeID uint64, id [16]byte) (object []byte, err Error)
	Read(mediaTypeID uint64, id [16]byte, offset, limit uint64) (data []byte, err Error)
	Save(mediaTypeID uint64, id [16]byte, object []byte) (err Error)
	Write(mediaTypeID uint64, id [16]byte, offset uint64, data []byte) (err Error)

	Append(mediaTypeID uint64, id [16]byte, data []byte) (err Error)
	Prepend(mediaTypeID uint64, id [16]byte, data []byte) (err Error)
	Extend(mediaTypeID uint64, id [16]byte, length uint64) (err Error)

	// make invisible just by remove from primary index
	Delete(mediaTypeID uint64, id [16]byte) (err Error)
	// make invisible by remove from primary index & write zero data to object location
	Erase(mediaTypeID uint64, id [16]byte) (err Error)

	// Multiple changes can be made in one atomic batch
	// Batch()
}
