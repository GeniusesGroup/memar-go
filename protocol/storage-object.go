/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Storage_Object = Buffer

// StorageObjects is the interface that show how an object directory work.
// Object owner is one app so it must handle concurrent protection internally.
// Strongly think fundamentally due to all data in computer must store as k/v in storages finally even files or rows.
// - mediaTypeID equivalent to table-name(RDB), namespace(Key-Value) or BucketID(AWS S3), but mediaTypeID has more meaning and has fixed size.
// - 16 byte id choose to defeat https://en.wikipedia.org/wiki/Birthday_problem due to we suggest use hash(sha3-256 to 128bit) of any data as object key.
// If need encryption, Implement requirements at block storage level.
type StorageObjects interface {
	MediatypeNumbers() (num uint64, err Error)
	ListMediatypeIDs(offset ElementIndex, limit NumberOfElement) (ids []uint64, err Error)

	ObjectNumbers(mt MediaTypeID) (num uint64, err Error)
	ListObjects(mt MediaTypeID, offset ElementIndex, limit NumberOfElement) (ids []UUID, err Error)

	Lock(mt MediaTypeID, id UUID) (object Storage_Object, err Error)
	Unlock(mt MediaTypeID, id UUID, object Storage_Object) (err Error)

	Length(mt MediaTypeID, id UUID) (ln NumberOfElement, err Error)
	Get(mt MediaTypeID, id UUID) (object Storage_Object, err Error)
	Read(mt MediaTypeID, id UUID, offset ElementIndex, limit NumberOfElement) (data Storage_Object, err Error)
	Save(mt MediaTypeID, id UUID, object Storage_Object) (err Error)
	Write(mt MediaTypeID, id UUID, offset ElementIndex, data Storage_Object) (err Error)

	Append(mt MediaTypeID, id UUID, data Storage_Object) (err Error)
	Prepend(mt MediaTypeID, id UUID, data Storage_Object) (err Error)
	Extend(mt MediaTypeID, id UUID, length uint64) (err Error)

	// make invisible just by remove from primary index
	Delete(mt MediaTypeID, id UUID) (err Error)
	// make invisible by remove from primary index & write zero data to object location
	Erase(mt MediaTypeID, id UUID) (err Error)

	// Multiple changes can be made in one atomic batch
	// Batch()
}
