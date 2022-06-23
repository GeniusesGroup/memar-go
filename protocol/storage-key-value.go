/* For license and copyright information please see LEGAL file in repository */

package protocol

// StorageKeyValue is the interface that show how an key-value directory work.
// If need encryption, Implement requirements at block storage level.
// https://github.com/cockroachdb/pebble/blob/master/open.go
type StorageKeyValue interface {
	KeyNumbers() (num uint64, err Error)
	ListKeys(offset, limit uint64) (keys [][]byte, err Error)

	Lock(key []byte) (value []byte, err Error)
	Unlock(key []byte, value []byte) (err Error)

	Length(key []byte) (ln int, err Error)
	Get(key []byte) (value []byte, err Error)
	Set(key []byte, value []byte, options StorageKeyValue_SaveOptions) (err Error)

	// make invisible just by remove from primary index
	Delete(key []byte) (err Error)
	// make invisible by remove from primary index & write zero data to value location
	Erase(key []byte) (err Error)

	// Multiple changes can be made in one atomic batch
	// Batch()
}

type StorageKeyValue_SaveOptions struct {
	// TTL(Time-To-Live) or Expiration Number of nanoseconds until record expires.
	TTL Duration
}
