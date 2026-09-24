/* For license and copyright information please see the LEGAL file in the code repository */

package pehrest_p

import (
	container_p "memar/computer/adt/container/protocol"
	error_p "memar/process/error/protocol"
	record_p "memar/storage/record/protocol"
)

// Model is the interface that show how a pehrest index system work.
// https://github.com/geniusesgroup/memar/
type Model interface {
	KeysNumbers() (num container_p.NumberOfElement, err error_p.Error)
	ListKeys(offset container_p.ElementIndex, limit container_p.NumberOfElement) (keys []Key, err error_p.Error)

	Lock(key Key) (lastUUID record_p.UUID, err error_p.Error)
	Unlock(key Key, latestUUID record_p.UUID) (err error_p.Error)

	Get(key Key) (uuid record_p.UUID, err error_p.Error)
	Set(key Key, uuid record_p.UUID) (err error_p.Error)
	// TODO::: Do we need update method?? or just return error in set() and force dev to delete old key??
	// Update()

	// make invisible just by remove from primary index
	Delete(key Key) (err error_p.Error)
}
