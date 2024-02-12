/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Storage_Key []byte
type Storage_Value = Buffer

// StorageKeyValue is the interface that show how an key-value directory work.
// If need encryption, Implement requirements at block storage level.
// https://github.com/cockroachdb/pebble/blob/master/open.go
type StorageKeyValue interface {
	KeyNumbers() (num uint64, err Error)
	ListKeys(offset ElementIndex, limit NumberOfElement) (keys []Storage_Key, err Error)

	Lock(key Storage_Key) (value Storage_Value, err Error)
	Unlock(key Storage_Key, value Storage_Value) (err Error)

	Length(key Storage_Key) (ln NumberOfElement, err Error)
	Get(key Storage_Key) (value Storage_Value, err Error)
	Set(key Storage_Key, value Storage_Value) (err Error)

	// make invisible just by remove from primary index
	Delete(key Storage_Key) (err Error)
	// make invisible by remove from primary index & write zero data to value location
	Erase(key Storage_Key) (err Error)

	// Multiple changes can be made in one atomic batch
	// Batch()
}
